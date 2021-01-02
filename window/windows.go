package window

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	myJson "github.com/araoko/cspusage/model/json"
	"github.com/araoko/cspusagefe/panel"
	"github.com/araoko/cspusagefe/rest"
	"github.com/icza/gowut/gwu"
)

type WindowMaker struct {
	r          rest.RestClient
	pubSession gwu.Session
}

func NewWindowMaker(r rest.RestClient, s gwu.Session) WindowMaker {

	return WindowMaker{
		r:          r,
		pubSession: s,
	}
}

func (wm WindowMaker) CustomerListWindow(name string, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	msgLabel.Style().SetWhiteSpace("pre")
	dowloadRequestPanel := panel.NewDownloadAsPanel(wm.r.Server(), "cl")
	dowloadRequestPanel.SetLink("cl")
	win.Add(msgLabel)
	win.Add(dowloadRequestPanel)
	customers, _ := wm.r.GetCustomers()
	p := panel.PanelOfCustomers(customers)
	win.Add(p)
	return win
}

func (wm WindowMaker) CustomerMonthlyBillWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	msgLabel.Style().SetWhiteSpace("pre")
	dowloadRequestPanel := panel.NewDownloadAsPanel(wm.r.Server(), "cmb")
	win.Add(msgLabel)
	customers, err := wm.r.GetCustomers()
	if err != nil {
		msgLabel.SetText(err.Error())
	}
	resPanel := gwu.NewPanel()
	inputPanel := panel.NewCustomerYearMonthPanel(customers, 2017)
	inputPanel.SetHandler(func(e gwu.Event, c myJson.Customer, yr, mo int) {
		fr, err := wm.r.GetCustomerMontlyBill(c, yr, mo)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Fecthing result:: %v", err))
			e.MarkDirty(msgLabel)
			return
		}
		p2 := panel.PanelOfCustomerMonthyUsage(fr)
		resPanel.Clear()
		resPanel.Add(p2)
		dowloadRequestPanel.Style().Set("display", "")
		inj := myJson.CustomerIDAndDate{
			CustomerId: c.CustomerId,
			Date: myJson.YearMonth{
				Year:  yr,
				Month: mo,
			},
		}

		inb, err := json.Marshal(inj)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Marshaling::%v\nError::%v", inj, err))
			e.MarkDirty(msgLabel)
			return
		}
		b64 := base64.URLEncoding.EncodeToString(inb)
		dowloadRequestPanel.SetLink(b64)
		e.MarkDirty(resPanel, dowloadRequestPanel)

	})
	win.Add(inputPanel)

	dowloadRequestPanel.Style().Set("display", "none")
	win.Add(dowloadRequestPanel)
	win.Add(resPanel)

	return win

}

func (wm WindowMaker) CustomerMonthlyCostPerSubWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	msgLabel.Style().SetWhiteSpace("pre")
	dowloadRequestPanel := panel.NewDownloadAsPanel(wm.r.Server(), "cmbps")
	win.Add(msgLabel)
	customers, err := wm.r.GetCustomers()
	if err != nil {
		msgLabel.SetText(err.Error())
	}
	resPanel := gwu.NewPanel()
	inputPanel := panel.NewCustomerYearMonthPanel(customers, 2017)
	inputPanel.SetHandler(func(e gwu.Event, c myJson.Customer, yr, mo int) {
		fr, err := wm.r.GetCustomerMonthlyCostPerSub(c, yr, mo)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Fecthing result:: %v", err))
			e.MarkDirty(msgLabel)
			return
		}
		p2 := panel.PanelOfCustomerMonthyUsagePerSub(fr)
		resPanel.Clear()
		resPanel.Add(p2)
		dowloadRequestPanel.Style().Set("display", "")
		inj := myJson.CustomerIDAndDate{
			CustomerId: c.CustomerId,
			Date: myJson.YearMonth{
				Year:  yr,
				Month: mo,
			},
		}

		inb, err := json.Marshal(inj)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Marshaling::%v\nError::%v", inj, err))
			e.MarkDirty(msgLabel)
			return
		}
		b64 := base64.URLEncoding.EncodeToString(inb)
		dowloadRequestPanel.SetLink(b64)
		e.MarkDirty(resPanel, dowloadRequestPanel)

	})
	win.Add(inputPanel)

	dowloadRequestPanel.Style().Set("display", "none")
	win.Add(dowloadRequestPanel)
	win.Add(resPanel)

	return win

}

func (wm WindowMaker) MonthlyBillWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	dowloadRequestPanel := panel.NewDownloadAsPanel(wm.r.Server(), "mb")
	win.Add(msgLabel)
	// customers, err := getCustomers()
	// if err != nil {
	// 	msgLabel.SetText(err.Error())
	// }

	resPanel := gwu.NewPanel()
	inputPanel := panel.NewYearMonthPanel(2017)
	inputPanel.SetHandler(func(e gwu.Event, yr, mo int) {
		fr, err := wm.r.GetMontlyBill(yr, mo)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Fecthing result:: %v", err))
			e.MarkDirty(msgLabel)
			return
		}
		p2 := panel.PanelOfMonthyUsage(fr)
		resPanel.Clear()
		resPanel.Add(p2)
		dowloadRequestPanel.Style().Set("display", "")
		inj := myJson.YearMonth{
			Year:  yr,
			Month: mo,
		}
		inb, err := json.Marshal(inj)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Marshaling::%v\nError::%v", inj, err))
			e.MarkDirty(msgLabel)
			return
		}
		b64 := base64.URLEncoding.EncodeToString(inb)
		dowloadRequestPanel.SetLink(b64)
		e.MarkDirty(resPanel, dowloadRequestPanel)

	})
	win.Add(inputPanel)

	dowloadRequestPanel.Style().Set("display", "none")
	win.Add(dowloadRequestPanel)
	win.Add(resPanel)

	return win

}

func (wm WindowMaker) RangeBillWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	dowloadRequestPanel := panel.NewDownloadAsPanel(wm.r.Server(), "rb")
	win.Add(msgLabel)

	resPanel := gwu.NewPanel()
	inputPanel := panel.NewDateRangePanel(2017)
	inputPanel.SetHandler(func(e gwu.Event, syr, smo, eyr, emo int) {
		fr, err := wm.r.GetRangeBill(syr, smo, eyr, emo)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Fecthing result:: %v", err))
			e.MarkDirty(msgLabel)
			return
		}
		p2 := panel.PanelOfRangeUsage(fr)
		resPanel.Clear()
		resPanel.Add(p2)
		dowloadRequestPanel.Style().Set("display", "")
		inj := myJson.YearMonthRange{
			StartDate: myJson.YearMonth{
				Year:  syr,
				Month: smo,
			},
			EndDate: myJson.YearMonth{
				Year:  eyr,
				Month: emo,
			},
		}
		inb, err := json.Marshal(inj)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Marshaling::%v\nError::%v", inj, err))
			e.MarkDirty(msgLabel)
			return
		}
		b64 := base64.URLEncoding.EncodeToString(inb)
		dowloadRequestPanel.SetLink(b64)
		e.MarkDirty(resPanel, dowloadRequestPanel)

	})
	win.Add(inputPanel)
	dowloadRequestPanel.Style().Set("display", "none")
	win.Add(dowloadRequestPanel)
	win.Add(resPanel)
	return win
}

func (wm WindowMaker) CustomerMonthlyTrendPerSubWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	dowloadRequestPanel := panel.NewDownloadAsPanel(wm.r.Server(), "cmtps")
	win.Add(msgLabel)

	resPanel := gwu.NewPanel()
	custs, _ := wm.r.GetCustomers()

	inputPanel := panel.NewCustomerDateRangePanel(custs, 2017)
	inputPanel.SetHandler(func(e gwu.Event, c myJson.Customer, syr, smo, eyr, emo int) {

		fr, err := wm.r.GetCustomerPerSubTrend(c.CustomerId, syr, smo, eyr, emo)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Fectching result:: %v", err))
			e.MarkDirty(msgLabel)
			return
		}

		p2 := panel.PanelOfCustomerMonthlyPerSubTrend(fr)

		resPanel.Clear()
		resPanel.Add(p2)
		dowloadRequestPanel.Style().Set("display", "")
		inj := myJson.CustomerIDAndDateRange{
			CustomerId: c.CustomerId,
			DateRange: myJson.YearMonthRange{
				StartDate: myJson.YearMonth{
					Year:  syr,
					Month: smo,
				},
				EndDate: myJson.YearMonth{
					Year:  eyr,
					Month: emo,
				},
			},
		}

		inb, err := json.Marshal(inj)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Marshaling::%v\nError::%v", inj, err))
			e.MarkDirty(msgLabel)
			return
		}
		b64 := base64.URLEncoding.EncodeToString(inb)
		dowloadRequestPanel.SetLink(b64)
		e.MarkDirty(resPanel, dowloadRequestPanel)

	})
	win.Add(inputPanel)
	dowloadRequestPanel.Style().Set("display", "none")
	win.Add(dowloadRequestPanel)
	win.Add(resPanel)
	return win
}

func (wm WindowMaker) CustomerMonthlyTrendWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	dowloadRequestPanel := panel.NewDownloadAsPanel(wm.r.Server(), "cmt")
	win.Add(msgLabel)

	resPanel := gwu.NewPanel()
	custs, _ := wm.r.GetCustomers()

	inputPanel := panel.NewCustomerDateRangePanel(custs, 2017)
	inputPanel.SetHandler(func(e gwu.Event, c myJson.Customer, syr, smo, eyr, emo int) {
		fr, err := wm.r.GetCustomerTrend(c.CustomerId, syr, smo, eyr, emo)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Fecthing result:: %v", err))
			e.MarkDirty(msgLabel)
			return
		}
		p2 := panel.PanelOfCustomerMonthlyTrend(fr)
		resPanel.Clear()
		resPanel.Add(p2)
		dowloadRequestPanel.Style().Set("display", "")
		inj := myJson.CustomerIDAndDateRange{
			CustomerId: c.CustomerId,
			DateRange: myJson.YearMonthRange{
				StartDate: myJson.YearMonth{
					Year:  syr,
					Month: smo,
				},
				EndDate: myJson.YearMonth{
					Year:  eyr,
					Month: emo,
				},
			},
		}

		inb, err := json.Marshal(inj)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Marshaling::%v\nError::%v", inj, err))
			e.MarkDirty(msgLabel)
			return
		}
		b64 := base64.URLEncoding.EncodeToString(inb)
		dowloadRequestPanel.SetLink(b64)
		e.MarkDirty(resPanel, dowloadRequestPanel)

	})
	win.Add(inputPanel)
	dowloadRequestPanel.Style().Set("display", "none")
	win.Add(dowloadRequestPanel)
	win.Add(resPanel)
	return win
}

func (wm WindowMaker) MonthlyTrendWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	dowloadRequestPanel := panel.NewDownloadAsPanel(wm.r.Server(), "mt")
	win.Add(msgLabel)

	resPanel := gwu.NewPanel()
	inputPanel := panel.NewDateRangePanel(2017)
	inputPanel.SetHandler(func(e gwu.Event, syr, smo, eyr, emo int) {
		fr, err := wm.r.GetTrend(syr, smo, eyr, emo)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Fecthing result:: %v", err))
			e.MarkDirty(msgLabel)
			return
		}
		p2 := panel.PanelOfMonthlyTrend(fr)
		resPanel.Clear()
		resPanel.Add(p2)
		dowloadRequestPanel.Style().Set("display", "")
		inj := myJson.YearMonthRange{
			StartDate: myJson.YearMonth{
				Year:  syr,
				Month: smo,
			},
			EndDate: myJson.YearMonth{
				Year:  eyr,
				Month: emo,
			},
		}
		inb, err := json.Marshal(inj)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Marshaling::%v\nError::%v", inj, err))
			e.MarkDirty(msgLabel)
			return
		}
		b64 := base64.URLEncoding.EncodeToString(inb)
		dowloadRequestPanel.SetLink(b64)
		e.MarkDirty(resPanel, dowloadRequestPanel)

	})
	win.Add(inputPanel)
	dowloadRequestPanel.Style().Set("display", "none")
	win.Add(dowloadRequestPanel)
	win.Add(resPanel)
	return win
}

func (wm WindowMaker) CustomerRangeBillWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	dowloadRequestPanel := panel.NewDownloadAsPanel(wm.r.Server(), "crb")
	win.Add(msgLabel)
	customers, err := wm.r.GetCustomers()
	if err != nil {
		msgLabel.SetText(err.Error())
	}
	resPanel := gwu.NewPanel()
	inputPanel := panel.NewCustomerDateRangePanel(customers, 2017)
	inputPanel.SetHandler(func(e gwu.Event, c myJson.Customer, syr, smo, eyr, emo int) {
		fr, err := wm.r.GetCustomerRangeBill(c, syr, smo, eyr, emo)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Fecthing result:: %v", err))
			e.MarkDirty(msgLabel)
			return
		}
		p2 := panel.PanelOfCustomerRangeUsage(fr)
		resPanel.Clear()
		resPanel.Add(p2)
		dowloadRequestPanel.Style().Set("display", "")
		inj := myJson.CustomerIDAndDateRange{
			CustomerId: c.CustomerId,
			DateRange: myJson.YearMonthRange{
				StartDate: myJson.YearMonth{
					Year:  syr,
					Month: smo,
				},
				EndDate: myJson.YearMonth{
					Year:  eyr,
					Month: emo,
				},
			},
		}

		inb, err := json.Marshal(inj)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Marshaling::%v\nError::%v", inj, err))
			e.MarkDirty(msgLabel)
			return
		}
		b64 := base64.URLEncoding.EncodeToString(inb)
		dowloadRequestPanel.SetLink(b64)
		e.MarkDirty(resPanel, dowloadRequestPanel)

	})
	win.Add(inputPanel)
	dowloadRequestPanel.Style().Set("display", "none")
	win.Add(dowloadRequestPanel)
	win.Add(resPanel)
	return win

}

func (wm WindowMaker) CustomerRangeCostPerSubWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	dowloadRequestPanel := panel.NewDownloadAsPanel(wm.r.Server(), "crbps")
	win.Add(msgLabel)
	customers, err := wm.r.GetCustomers()
	if err != nil {
		msgLabel.SetText(err.Error())
	}
	resPanel := gwu.NewPanel()
	inputPanel := panel.NewCustomerDateRangePanel(customers, 2017)
	inputPanel.SetHandler(func(e gwu.Event, c myJson.Customer, syr, smo, eyr, emo int) {
		fr, err := wm.r.GetCustomerRangeCostPerSub(c, syr, smo, eyr, emo)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Fecthing result:: %v", err))
			e.MarkDirty(msgLabel)
			return
		}
		p2 := panel.PanelOfCustomerRangeUsagePerSub(fr)
		resPanel.Clear()
		resPanel.Add(p2)
		dowloadRequestPanel.Style().Set("display", "")
		inj := myJson.CustomerIDAndDateRange{
			CustomerId: c.CustomerId,
			DateRange: myJson.YearMonthRange{
				StartDate: myJson.YearMonth{
					Year:  syr,
					Month: smo,
				},
				EndDate: myJson.YearMonth{
					Year:  eyr,
					Month: emo,
				},
			},
		}

		inb, err := json.Marshal(inj)
		if err != nil {
			msgLabel.SetText(fmt.Sprintf("Error Marshaling::%v\nError::%v", inj, err))
			e.MarkDirty(msgLabel)
			return
		}
		b64 := base64.URLEncoding.EncodeToString(inb)
		dowloadRequestPanel.SetLink(b64)
		e.MarkDirty(resPanel, dowloadRequestPanel)

	})
	win.Add(inputPanel)
	dowloadRequestPanel.Style().Set("display", "none")
	win.Add(dowloadRequestPanel)
	win.Add(resPanel)
	return win

}

func (wm WindowMaker) LogInWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Watch This Space.....")
	win.Add(msgLabel)
	statusLabel := gwu.NewLabel("")
	statusLabel.Style().SetColor("red")
	win.AddVSpace(20)
	win.Add(statusLabel)
	win.AddVSpace(20)
	headerLabel := gwu.NewLabel("Log In with Sidmach AD Credentials")
	hdStyle := headerLabel.Style()
	hdStyle.SetFontSize("30")
	hdStyle.SetFontWeight("bold")
	win.Add(headerLabel)
	win.AddVSpace(25)

	resPanel := gwu.NewPanel()
	tbl := gwu.NewTable()
	tbl.SetCellSpacing(10)
	tbl.Add(gwu.NewLabel("Username"), 0, 0)
	tbl.Add(gwu.NewLabel("Password"), 1, 0)
	userNameInput := gwu.NewTextBox("")
	userNameInput.SetToolTip("UserName")
	tbl.Add(userNameInput, 0, 1)

	passwordInput := gwu.NewPasswBox("")
	passwordInput.SetToolTip("Password")
	tbl.Add(passwordInput, 1, 1)

	logBtn := gwu.NewButton("OK")

	logBtn.AddEHandlerFunc(func(e gwu.Event) {
		res, err := wm.r.Auth(userNameInput.Text(), passwordInput.Text())
		if err != nil {
			statusLabel.SetText("Login Error")
			msgLabel.SetText(err.Error())
			e.MarkDirty(statusLabel, msgLabel)
			return
		}
		fmt.Printf("Login response: %v\n", res)
		if res.LoggedIn == 0 {
			fmt.Println("Login succesful for ", res.UserName)
			//e.MarkDirty(statusLabel)
			return
		}
		e.Session().RemoveWin(win)
		e.Session().SetAttr("userName", res.UserName)
		fmt.Println("building session windows")
		buildsessionwindows(e.Session(), wm)
		fmt.Println("session windows built")
		e.ReloadWin("main")
	}, gwu.ETypeClick)

	cancleBtn := gwu.NewButton("Cancel")
	cancleBtn.AddEHandlerFunc(func(e gwu.Event) {
		e.RemoveSess()
		e.ReloadWin("home")
	}, gwu.ETypeClick)

	tbl.Add(cancleBtn, 2, 1)
	tbl.Add(logBtn, 2, 0)
	resPanel.Add(tbl)
	win.Add(resPanel)

	win.SetFocusedCompID(userNameInput.ID())
	return win

}

func (wm WindowMaker) HomeWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Click the button Bellow to Log in and start your Session")
	win.Add(msgLabel)
	win.AddVSpace(20)
	resPanel := gwu.NewPanel()
	btn := gwu.NewButton("Start LogIn")
	btn.AddEHandlerFunc(func(e gwu.Event) {
		e.ReloadWin("login")
	}, gwu.ETypeClick)
	resPanel.Add(btn)
	win.Add(resPanel)
	return win
}

func (wm WindowMaker) MainWindow(name, text string) gwu.Window {
	win := getWin(name, text)
	msgLabel := gwu.NewLabel("Welcome")
	win.Add(msgLabel)
	win.AddVSpace(20)
	pubTbl := gwu.NewTable()
	pubTbl.Add(gwu.NewLabel("Shared Windows"), 0, 0)
	pubTbl.SetCellSpacing(5)
	privTble := gwu.NewTable()
	privTble.Add(gwu.NewLabel("Private Windows"), 0, 0)
	resPanel := gwu.NewPanel()
	resPanel.Add(pubTbl)
	pubWins := wm.pubSession.SortedWins()

	for i, w := range pubWins {
		if w.Text() == "Home Window" {
			continue
		}
		winTxt := fmt.Sprintf("%s", w.Text())
		l := gwu.NewLink(winTxt, w.Name())
		pubTbl.Add(l, i+1, 0)
	}
	resPanel.AddVSpace(40)
	resPanel.Add(privTble)
	win.Add(resPanel)
	win.AddEHandlerFunc(func(e gwu.Event) {
		s := e.Session()
		if s.Private() {
			msgLabel.SetText(fmt.Sprintf("Welcome %v", s.Attr("userName")))
			privWins := s.SortedWins()
			for i, w := range privWins {
				winTxt := fmt.Sprintf("%s", w.Text())
				l := gwu.NewLink(winTxt, w.Name())
				privTble.Add(l, i+1, 0)
			}
			e.MarkDirty(msgLabel, privTble)

		}
	}, gwu.ETypeWinLoad)

	return win
}

func BuildLoginWindow(s gwu.Session, wm WindowMaker) {
	loginWin := wm.LogInWindow("login", "Login Window")
	err := s.AddWin(loginWin)

	if err != nil {
		fmt.Println("Error adding login window:", err)
	}
}

func buildsessionwindows(s gwu.Session, wm WindowMaker) {

	mainWindow := wm.MainWindow("main", "Main Window")
	fmt.Printf("Created %v window: %v \n", mainWindow.Name(), mainWindow)

	customerList := wm.CustomerListWindow("customerList", "List of CSP Accounts")
	fmt.Printf("Created %v window: %v \n", customerList.Name(), customerList)

	custMonthlyBillWin := wm.CustomerMonthlyBillWindow("monthlyBill", "Azure Usage Cost - Specified Month, Specified Account")
	fmt.Printf("Created %v window: %v \n", custMonthlyBillWin.Name(), custMonthlyBillWin)

	custMonthlyCostPerSubWin := wm.CustomerMonthlyCostPerSubWindow("monthlyCostPerSub", "Azure Usage Cost Per Subscription - Specified Month, Specified Account")
	fmt.Printf("Created %v window: %v \n", custMonthlyCostPerSubWin.Name(), custMonthlyCostPerSubWin)

	monthlyBillWin := wm.MonthlyBillWindow("monthlyBills", "Azure Usage Cost - Specified Month, All Accounts")
	fmt.Printf("Created %v window: %v \n", monthlyBillWin.Name(), monthlyBillWin)

	custRangeCostPerSubWin := wm.CustomerRangeCostPerSubWindow("rangeCostPerSub", "Azure Usage Cost Per Subscription - Specified Period, Specified Account")
	fmt.Printf("Created %v window: %v \n", custRangeCostPerSubWin.Name(), custRangeCostPerSubWin)

	custRangeBillWin := wm.CustomerRangeBillWindow("rangeBill", "Azure Usage Cost - Specified Period, Specified Account")
	fmt.Printf("Created %v window: %v \n", custRangeBillWin.Name(), custRangeBillWin)

	rangeBillWin := wm.RangeBillWindow("rangeBills", "Azure Usage Cost - Specified Period, All Accounts")
	fmt.Printf("Created %v window: %v \n", custRangeBillWin.Name(), rangeBillWin)

	monthlyTrendWin := wm.MonthlyTrendWindow("monthlyTrend", "Azure Monthly Trend - Specified Period, All Accounts")
	fmt.Printf("Created %v window: %v \n", monthlyTrendWin.Name(), monthlyTrendWin)

	custMonthlyTrendWin := wm.CustomerMonthlyTrendWindow("custMonthlyTrend", "Azure Monthly Trend - Specified Period, Specified Account")
	fmt.Printf("Created %v window: %v \n", custMonthlyTrendWin.Name(), custMonthlyTrendWin)

	custMonthlyTrendPerSubWin := wm.CustomerMonthlyTrendPerSubWindow("custMonthlySubTrend", "Azure Monthly Trend Per Subscription- Specified Period, Specified Account")
	fmt.Printf("Created %v window: %v \n", custMonthlyTrendPerSubWin.Name(), custMonthlyTrendPerSubWin)
	var err error
	err = s.AddWin(mainWindow)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

	err = s.AddWin(customerList)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

	err = s.AddWin(custMonthlyBillWin)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

	err = s.AddWin(rangeBillWin)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

	err = s.AddWin(monthlyTrendWin)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

	err = s.AddWin(custMonthlyTrendWin)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

	err = s.AddWin(custMonthlyTrendPerSubWin)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

	err = s.AddWin(monthlyBillWin)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

	err = s.AddWin(custMonthlyCostPerSubWin)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

	err = s.AddWin(custRangeCostPerSubWin)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

	err = s.AddWin(custRangeBillWin)
	if err != nil {
		fmt.Printf("Error Adding Window %s : %s \n", mainWindow.Name(), err.Error())
	}

}

func getWin(name, text string) gwu.Window {
	win := gwu.NewWindow(name, text)
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HACenter)
	win.SetCellPadding(2)
	return win
}
