package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// 定义窗口
type MyMainWindow struct {
	*walk.MainWindow
}

// 下拉框
type Species struct {
	Id   int
	Name string
}

func main() {

	var te *walk.LineEdit
	var gc *walk.Action

	mw := new(MyMainWindow)

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Bitcoin",
		Icon:     "./rc/bitcoin.ico",
		Layout:   VBox{},
		Name:     "bitcoin",
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						Text:        "Exit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&Option",
				Items: []MenuItem{
					Action{
						AssignTo: &gc,
						Text:     "Generate Coins",
						OnTriggered: func() {
							gc.SetChecked(!gc.Checked())
						},
					},
					Action{
						Text: "Options",
						OnTriggered: func() {
							OptionDialog(mw)
						},
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text: "About",
						OnTriggered: func() {
							AboutDialog(mw)
						},
					},
				},
			},
		},
		ToolBar: ToolBar{

			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				Action{
					Text:  "Send Coins",
					Image: "./rc/send16.bmp",

					OnTriggered: func() {
						SendDialog(mw)
					},
				},
				Separator{},
				Action{
					Text:  "Address Book",
					Image: "./rc/addressbook16.bmp",
					OnTriggered: func() {
						BookDialog(mw)
					},
				},
			},
		},

		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 6},
				Children: []Widget{
					Label{
						Text: "Your Bitcoin Address:",
					},
					LineEdit{
						Text:     "1PrMRMbBhGhDKw1C2Zm92nY4uwtQViaeyd",
						Enabled:  false,
						AssignTo: &te,
					},
					HSpacer{},
					HSpacer{},
					PushButton{
						Text: "Copy to Clipboard",
						OnClicked: func() {
							if err := walk.Clipboard().SetText(te.Text()); err != nil {
								log.Print("Copy: ", err)
							}
						},
					},
					PushButton{
						Text: "Change...",
						OnClicked: func() {
							AddressDialog(mw)
						},
					},
				},
			},
			Composite{
				Layout: Grid{Columns: 6},
				Children: []Widget{
					Label{
						Text: "Balance:",
					},
					HSpacer{},
					Label{
						Text: "0.00",
					},
					HSpacer{},
					HSpacer{},
					HSpacer{},
				},
			},
			TableView{
				Name: "tableView",
				AlternatingRowBGColor: walk.RGB(255, 255, 200),
				ColumnsOrderable:      true,
				Columns: []TableViewColumn{

					{Name: "Status"}, // Use DataMember, if names differ
					{Name: "Data", Format: "2006-01-02 15:04:05", Width: 130},
					{Name: "Description", Format: "2006-01-02 15:04:05", Width: 155},
					{Name: "Debit", Format: "2006-01-02 15:04:05", Width: 140},
					{Name: "Credit", Format: "2006-01-02 15:04:05", Width: 140},
				},
				OnItemActivated: func() {},
				Model:           NewFooModel(),
			},
			Composite{
				Layout: Grid{Columns: 5},
				Children: []Widget{
					Label{
						Text: "0 connections",
					},
					HSpacer{},
					Label{
						Text: "1 blocks",
					},
					HSpacer{},
					Label{
						Text: "0 transactions",
					},
				},
			},
		},

		MinSize: Size{705, 484},
	}.Create()); err != nil {
		log.Fatal(err)
	}
	mw.Run()
}

func (mw *MyMainWindow) showAboutBoxAction_Triggered() {

	var LableHello = Label{
		Text: "Bitcoin world!",
	}

	var widget = []Widget{
		LableHello,
	}
	var About_window = MainWindow{
		Title:    "About",
		MinSize:  Size{400, 200},
		Layout:   VBox{},
		Children: widget,
	}
	About_window.Run()
}
func NewFooModel() *FooModel {
	now := time.Now()

	rand.Seed(now.UnixNano())

	m := &FooModel{items: make([]*Foo, 10)}

	for i := range m.items {
		m.items[i] = &Foo{
			Status:      i,
			Data:        time.Unix(rand.Int63n(now.Unix()), 0),
			Description: strings.Repeat("*", rand.Intn(5)+1),
			Debit:       strings.Repeat("*", rand.Intn(5)+1),
			Credit:      strings.Repeat("*", rand.Intn(5)+1),
		}
	}

	return m
}

type FooModel struct {
	walk.SortedReflectTableModelBase
	items []*Foo
}

func (m *FooModel) Items() interface{} {
	return m.items
}

type Foo struct {
	Status      int
	Data        time.Time
	Description string
	Debit       string
	Credit      string
}

func AboutDialog(owner walk.Form) (int, error) {
	var dlg *walk.Dialog
	return Dialog{
		AssignTo: &dlg,
		Title:    "About Bitcoin",
		MinSize:  Size{507, 300},
		Layout:   VBox{},
		Children: []Widget{
			TextEdit{
				Text: `BitCoin v0.1.3 ALPHA
Copyright (c) 2009 Satoshi Nakamoto
Distributed under the MIT/X11 software license, see the accompanying
file license.txt or http://www.opensource.org/licenses/mit-license.php.
This product includes software developed by the OpenSSL Project for use in
the OpenSSL Toolkit (http://www.openssl.org/).  This product includes
cryptographic software written by Eric Young (eay@cryptsoft.com).`,
				Enabled: false,
			},
			PushButton{
				Text:      "OK",
				OnClicked: func() { dlg.Cancel() },
			},
		},
	}.Run(owner)
}
func OptionDialog(owner walk.Form) (int, error) {
	var dlg *walk.Dialog
	return Dialog{
		AssignTo: &dlg,
		Title:    "Options",
		MinSize:  Size{500, 260},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text: "Transaction fee.",
			},
			Label{
				Text: "Transaction fee:",
			},
			NumberEdit{
				Value:    Bind("Weight", Range{0.01, 9999.99}),
				Suffix:   " kg",
				Decimals: 2,
			},
			PushButton{
				Text:      "OK",
				OnClicked: func() {},
			},
			PushButton{
				Text:      "Cancel",
				OnClicked: func() { dlg.Cancel() },
			},
		},
	}.Run(owner)
}
func SendDialog(owner walk.Form) (int, error) {
	var dlg *walk.Dialog
	return Dialog{
		AssignTo: &dlg,
		Title:    "Options",
		MinSize:  Size{500, 260},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{

					Label{
						Text: "e.g. 1NS17iag9",
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					HSplitter{
						Children: []Widget{
							Label{
								Text: "Pay To:",
							},
							TextEdit{
								ColumnSpan: 1,
								MinSize:    Size{200, 20},
								Text:       "",
							},
							PushButton{
								Text:      "Paste",
								OnClicked: func() {},
							},
							PushButton{
								Text:      "Address Book...",
								OnClicked: func() {},
							},
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					HSplitter{
						Children: []Widget{
							Label{
								Text: "Amount:",
							},
							TextEdit{
								ColumnSpan: 1,
								MinSize:    Size{200, 20},
								Text:       "",
							},
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					HSplitter{
						Children: []Widget{
							Label{
								Text: "Transfer:",
							},
							ComboBox{
								Value:         Bind("SpeciesId", SelRequired{}),
								BindingMember: "Id",
								DisplayMember: "Name",
								Model: []*Species{
									{1, "Standard"},
								},
							},
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					Label{
						Text: "From:",
					},
					TextEdit{
						ColumnSpan: 1,
						MinSize:    Size{200, 20},
						Text:       "",
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					Label{
						Text: "Message:",
					},
					TextEdit{
						ColumnSpan: 3,
						MinSize:    Size{200, 20},
						Text:       "",
					},
				},
			},
		},
	}.Run(owner)
}

func BookDialog(owner walk.Form) (int, error) {
	var dlg *walk.Dialog

	return Dialog{
		AssignTo: &dlg,
		Title:    "Options",
		MinSize:  Size{610, 390},
		Layout:   VBox{},
		Children: []Widget{
			TableView{
				Name: "tableView",
				AlternatingRowBGColor: walk.RGB(255, 255, 200),
				ColumnsOrderable:      true,
				Columns: []TableViewColumn{

					{Name: "Status"}, // Use DataMember, if names differ
					{Name: "Data", Format: "2006-01-02 15:04:05", Width: 130},
					{Name: "Description", Format: "2006-01-02 15:04:05", Width: 155},
					{Name: "Debit", Format: "2006-01-02 15:04:05", Width: 140},
					{Name: "Credit", Format: "2006-01-02 15:04:05", Width: 140},
				},
				OnItemActivated: func() {},
				Model:           NewFooModel(),
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text:      "Edit...",
						OnClicked: func() {},
					},
					PushButton{
						Text:      "NewA ddress...",
						OnClicked: func() {},
					},
					PushButton{
						Text:      "Delete",
						OnClicked: func() {},
					},
					PushButton{
						Text:      "OK",
						OnClicked: func() {},
					},
				},
			},
		},
	}.Run(owner)
}
func AddressDialog(owner walk.Form) (int, error) {
	var dlg *walk.Dialog
	return Dialog{
		AssignTo: &dlg,
		Title:    "Options",
		MinSize:  Size{610, 390},
		Layout:   VBox{},
		Children: []Widget{
			TableView{
				Name: "tableView",
				AlternatingRowBGColor: walk.RGB(255, 255, 200),
				ColumnsOrderable:      true,
				Columns: []TableViewColumn{

					{Name: "Status"}, // Use DataMember, if names differ
					{Name: "Data", Format: "2006-01-02 15:04:05", Width: 130},
					{Name: "Description", Format: "2006-01-02 15:04:05", Width: 155},
					{Name: "Debit", Format: "2006-01-02 15:04:05", Width: 140},
					{Name: "Credit", Format: "2006-01-02 15:04:05", Width: 140},
				},
				OnItemActivated: func() {},
				Model:           NewFooModel(),
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text:      "Rename...",
						OnClicked: func() {},
					},
					PushButton{
						Text:      "NewA ddress...",
						OnClicked: func() {},
					},
					PushButton{
						Text:      "Copy to Clipboard",
						OnClicked: func() {},
					},
					PushButton{
						Text:      "OK",
						OnClicked: func() {},
					},
					PushButton{
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}
