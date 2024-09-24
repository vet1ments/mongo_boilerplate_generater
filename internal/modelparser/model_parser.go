package modelparser

import (
	"boilerplate/internal/structparser"
	"boilerplate/internal/utils"
	"fmt"
	"strings"
)

type ModelField struct {
	IsEmbed bool
	Upper   string
	Snake   string
	Type    string
}

type Model struct {
	Name   string
	Fields []*ModelField
}

type ModelContainer struct {
	Models      []*Model
	EmbedModels []*Model
}

func IsModel(s string) bool {
	return strings.HasSuffix(s, "Model")
}

func ParseModel(structs []*structparser.StructInfo) *ModelContainer {
	container := &ModelContainer{}
	for _, st := range structs {
		model := &Model{
			Name: st.Name,
		}
		for _, f := range st.Fields {
			model.Fields = append(model.Fields, &ModelField{
				Upper: f.Name,
				Snake: utils.ToSnakeCase(f.Name),
				Type:  f.Type,
			})
		}
		if IsModel(model.Name) {
			model.Name = model.Name[:len(model.Name)-len("Model")]
			container.Models = append(container.Models, model)
		} else {
			container.EmbedModels = append(container.EmbedModels, model)
		}
	}
	return container
}

func PrintModelContainer(container *ModelContainer) {
	fmt.Println("----- Model Start -----")
	for _, model := range container.Models {
		fmt.Println("Model ", model.Name)
		for _, field := range model.Fields {
			fmt.Println("  ", field.Upper, " ", field.Type)
		}
	}
	fmt.Println("----- Model End -----")
	fmt.Println("\n\n")
	fmt.Println("----- EmbedModel Start -----")
	for _, embedModel := range container.EmbedModels {
		fmt.Println("EmbedModel ", embedModel.Name)
		for _, field := range embedModel.Fields {
			fmt.Println("  ", field.Upper, " ", field.Type)
		}
	}
	fmt.Println("----- EmbedModel End -----")
}
