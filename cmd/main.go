package main

import (
	"boilerplate/internal/modelparser"
	"boilerplate/internal/structparser"
	"boilerplate/internal/templatemaker"
	"flag"
	"fmt"
	"os"
	"sync"
)

var (
	versionFlag = flag.String("v", "", "specific Version")
	pathFlag    = flag.String("p", "", "specific Path")
	moduleFlag  = flag.String("m", "", "specific Module")
)

func exitProgram(s string) {
	fmt.Println(s)
	os.Exit(0)
}

func main() {
	//flag.Parse()
	//
	//if *pathFlag == "" {
	//	exitProgram("path 입력 없음")
	//}
	//path := *pathFlag
	//if *moduleFlag == "" {
	//	exitProgram("module 입력 없음")
	//}
	//module := *moduleFlag
	//
	//var version string
	//if *versionFlag == "" {
	//	version = *versionFlag
	//}
	makePath := "gen"
	inputPath := "test"

	files, err := os.ReadDir(inputPath)
	if err != nil {
		exitProgram("디렉토리 읽기 오류")
	}

	fileNames := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, inputPath+"/"+file.Name())
		}
	}

	var wg sync.WaitGroup
	var structs []*structparser.StructInfo
	for _, fileName := range fileNames {
		wg.Add(1)
		go func() {
			structs_, err := structparser.ParseStructs(fileName)
			if err != nil {
				exitProgram(err.Error())
			}
			structs = append(structs, structs_...)
			wg.Done()
		}()
	}
	wg.Wait()

	modelContainer := modelparser.ParseModel(structs)
	//modelparser.PrintModelContainer(modelContainer)
	err = templatemaker.CreateModelsFromTemplate(modelContainer, makePath, inputPath, "boilerplate")
	if err != nil {
		fmt.Println(err)
	}
	//templatemaker.CreateTemplateFile(modelContainer, "aaa")
}
