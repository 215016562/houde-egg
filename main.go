package main

import (
	"log"

	//"github.com/tarm/serial"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// var isSpecialMode = walk.NewMutableCondition()

type LayoutMainWindow struct{
	*walk.MainWindow
}

func main(){
	// MustRegisterCondition("isSpecialMode", isSpecialMode)
	mw := new(LayoutMainWindow)

	if err := (MainWindow{
		AssignTo : &mw.MainWindow,
		Title : "金翼蛋品蛋品收购系统- 网络数据管理器",
		Background: GradientBrush{
			Vertexes: []walk.GradientVertex{
				{X: 0, Y: 0, Color: walk.RGB(255, 255, 127)},
				{X: 1, Y: 0, Color: walk.RGB(127, 191, 255)},
				{X: 0.5, Y: 0.5, Color: walk.RGB(255, 255, 255)},
				{X: 1, Y: 1, Color: walk.RGB(127, 255, 127)},
				{X: 0, Y: 1, Color: walk.RGB(255, 127, 127)},
			},
			Triangles: []walk.GradientTriangle{
				{0, 1, 2},
				{1, 3, 2},
				{3, 4, 2},
				{4, 0, 2},
			},
		},
		MenuItems: []MenuItem{
			Menu{
				Text: "端口设置",
			},
			Menu{
				Text: "五号称设置",
			},
		},
		MinSize: Size{500, 300},
		Layout: HBox{},
		Children: []Widget{
			GroupBox{
				Title: "称头选择区",
				Layout: HBox{},
				Children: []Widget{
					CheckBox{
						Name: "",
						Text: "一号称",
						Checked: false,
					},
					CheckBox{
						Name: "",
						Text: "二号称",
						Checked: false,
					},
					CheckBox{
						Name: "",
						Text: "三号称",
						Checked: false,
					},
					CheckBox{
						Name: "",
						Text: "四号称",
						Checked: false,
					},
					CheckBox{
						Name: "",
						Text: "五号称",
						Checked: false,
					},
				},
			},
			// CheckBox{
			// 	Name: "one",
			// 	Text: "一号称",
			// 	Checked: false,
			// },
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	mw.Run()
}