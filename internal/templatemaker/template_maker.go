package templatemaker

import (
	"boilerplate/internal/modelparser"
	"boilerplate/internal/utils"
	"errors"
	"os"
	"strings"
	"text/template"
)

type ModelTemplateData struct {
	Version     string
	PackageName string
	*modelparser.Model
	Import string
}

type EmbedModelTemplateData struct {
	PackageName string
	Models      []*modelparser.Model
}

func CreateTemplateFile(container *modelparser.ModelContainer, path string, _version ...string) error {
	return nil
}

func CreateModelsFromTemplate(model *modelparser.ModelContainer, path string, currentDir string, modulePath string) error {
	path = pathSuffixCheck(path)
	modulePath = pathSuffixCheck(modulePath)

	done := make(chan error)

	go func(ch chan<- error) {
		defer close(ch)
		for _, m := range model.Models {
			ch <- createModelFromTemplate(m, path, currentDir, modulePath)
			return
		}
	}(done)
	for res := range done {
		if res != nil {
			return res
		}
	}

	err := createEmbedModelFromTemplate(model.EmbedModels, path, currentDir, modulePath)
	if err != nil {
		return err
	}
	return nil
}

func createEmbedModelFromTemplate(models []*modelparser.Model, path string, currentDir string, modulePath string) error {
	path = pathSuffixCheck(path)
	modulePath = pathSuffixCheck(modulePath)

	modelTmpl, err := template.ParseFiles("./templates/embed_model_template.go.tmpl")
	if err != nil {
		return errors.New("model 템플릿 파일 파싱 실패")
	}

	makePath := pathSuffixCheck(path + currentDir)
	modelPath := pathSuffixCheck(makePath + "mg_embed")
	err = utils.MkDir(modelPath)
	if err != nil {
		return err
	}

	f, err := os.Create(modelPath + "mg_" + "embed" + ".go")
	if err != nil {
		return err
	}

	data := &EmbedModelTemplateData{
		PackageName: "mgembed" + currentDir,
		Models:      models,
	}

	err = modelTmpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}

func createModelFromTemplate(model *modelparser.Model, path string, currentDir string, modulePath string) error {
	path = pathSuffixCheck(path)
	modulePath = pathSuffixCheck(modulePath)

	modelTmpl, err := template.ParseFiles("./templates/model_template.go.tmpl")
	if err != nil {
		return errors.New("model 템플릿 파일 파싱 실패")
	}

	makePath := pathSuffixCheck(path + currentDir)
	modelPath := pathSuffixCheck(makePath + "mg_" + strings.ToLower(model.Name))
	err = utils.MkDir(modelPath)
	if err != nil {
		return err
	}

	f, err := os.Create(modelPath + "mg_" + strings.ToLower(model.Name) + ".go")
	if err != nil {
		return err
	}

	data := &ModelTemplateData{
		PackageName: strings.ToLower(model.Name) + currentDir,
		Model:       model,
	}

	for _, v := range model.Fields {
		if !utils.IsValidType(v.Type) {
			data.Import = modulePath + makePath + "mg_embed"
			v.Type = "mgembed" + currentDir + "." + v.Type
		}
	}

	err = modelTmpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}

func pathSuffixCheck(s string) string {
	if strings.HasSuffix(s, "/") {
		return s
	}
	return s + "/"
}
