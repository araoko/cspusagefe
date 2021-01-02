package panel

import (
	"fmt"
	"strconv"
	"time"

	myJson "github.com/araoko/cspusage/model/json"
	"github.com/icza/gowut/gwu"
)

type CustomerListBox struct {
	gwu.ListBox

	customerList []myJson.Customer
}

func (c CustomerListBox) GetSelectedCustomer() myJson.Customer {
	s := c.SelectedIdx()
	if s == -1 {
		return myJson.Customer{}
	}
	return c.customerList[s]
}

func NewCustomerListBox(customers []myJson.Customer) CustomerListBox {
	dl := make([]string, len(customers))
	for i, c := range customers {
		dl[i] = c.CustomerCompanyName
	}

	lb := CustomerListBox{
		ListBox:      gwu.NewListBox(dl),
		customerList: customers,
	}
	lb.SetMulti(false)
	lb.SetRows(1)
	lb.SetSelected(0, true)
	return lb
}

type YearMonthPanel struct {
	gwu.Panel
	handler func(e gwu.Event, yr, mo int)
}

func (c *YearMonthPanel) SetHandler(f func(e gwu.Event, yr, mo int)) {
	c.handler = f
}

func NewYearMonthPanel(startYr int) *YearMonthPanel {

	embPanel := gwu.NewHorizontalPanel()
	embPanel.SetCellPadding(3)
	embPanel.SetCellSpacing(3)

	p := YearMonthPanel{Panel: embPanel}

	yrLb := getYearListBox(startYr)
	p.Add(yrLb)

	monLb := getMonthListBox()
	p.Add(monLb)

	b := gwu.NewButton("Submit")

	b.AddEHandlerFunc(func(e gwu.Event) {
		if p.handler == nil {
			return
		}
		y := p.CompAt(0).(gwu.ListBox)
		m := p.CompAt(1).(gwu.ListBox)

		//log.Println("in handle, yr index =", y, " Month index=", m)
		p.handler(e, y.SelectedIdx()+startYr, m.SelectedIdx()+1)

		//log.Println("item selected: ", cl.GetSelectedCustomer())
	}, gwu.ETypeClick)
	p.Add(b)

	return &p

}

type CustomerYearMonthPanel struct {
	gwu.Panel
	handler func(e gwu.Event, c myJson.Customer, yr, mo int)
}

func (c *CustomerYearMonthPanel) SetHandler(f func(e gwu.Event, c myJson.Customer, yr, mo int)) {
	c.handler = f
}

func NewCustomerYearMonthPanel(cust []myJson.Customer, startYr int) *CustomerYearMonthPanel {

	embPanel := gwu.NewHorizontalPanel()
	embPanel.SetCellPadding(3)
	embPanel.SetCellSpacing(3)

	p := CustomerYearMonthPanel{Panel: embPanel}
	p.Add(NewCustomerListBox(cust))

	yrLb := getYearListBox(startYr)
	p.Add(yrLb)

	monLb := getMonthListBox()
	p.Add(monLb)

	b := gwu.NewButton("Submit")

	b.AddEHandlerFunc(func(e gwu.Event) {
		if p.handler == nil {
			return
		}
		c := p.CompAt(0).(CustomerListBox)
		y := p.CompAt(1).(gwu.ListBox)
		m := p.CompAt(2).(gwu.ListBox)

		//log.Println("in handle, yr index =", y, " Month index=", m)
		p.handler(e, c.GetSelectedCustomer(), y.SelectedIdx()+startYr, m.SelectedIdx()+1)

		//log.Println("item selected: ", cl.GetSelectedCustomer())
	}, gwu.ETypeClick)
	p.Add(b)

	return &p

}

type DateRangePanel struct {
	gwu.Panel
	handler func(e gwu.Event, syr, smo, eyr, emo int)
}

func (c *DateRangePanel) SetHandler(f func(e gwu.Event, syr, smo, eyr, emo int)) {
	c.handler = f
}

func NewDateRangePanel(startYr int) *DateRangePanel {

	embPanel := gwu.NewHorizontalPanel()
	embPanel.SetCellPadding(3)
	embPanel.SetCellSpacing(3)

	p := DateRangePanel{Panel: embPanel}
	//p.Add(NewCustomerListBox(cust))
	dateRangeTble := gwu.NewTable()
	dateRangeTble.SetCellSpacing(3)
	dateRangeTble.SetCellPadding(3)

	syrLb := getYearListBox(startYr)
	dateRangeTble.Add(syrLb, 0, 0)

	smonLb := getMonthListBox()
	dateRangeTble.Add(smonLb, 0, 1)

	eyrLb := getYearListBox(startYr)
	dateRangeTble.Add(eyrLb, 1, 0)

	emonLb := getMonthListBox()
	dateRangeTble.Add(emonLb, 1, 1)

	p.Add(dateRangeTble)

	b := gwu.NewButton("Submit")

	b.AddEHandlerFunc(func(e gwu.Event) {
		if p.handler == nil {
			return
		}
		t := p.CompAt(0).(gwu.Table)
		sy := t.CompAt(0, 0).(gwu.ListBox)
		sm := t.CompAt(0, 1).(gwu.ListBox)

		ey := t.CompAt(1, 0).(gwu.ListBox)
		em := t.CompAt(1, 1).(gwu.ListBox)

		//log.Println("in handle, yr index =", y, " Month index=", m)
		p.handler(e, sy.SelectedIdx()+startYr, sm.SelectedIdx()+1, ey.SelectedIdx()+startYr, em.SelectedIdx()+1)

		//log.Println("item selected: ", cl.GetSelectedCustomer())
	}, gwu.ETypeClick)
	p.Add(b)

	return &p

}

type CustomerDateRangePanel struct {
	gwu.Panel
	handler func(e gwu.Event, c myJson.Customer, syr, smo, eyr, emo int)
}

func (c *CustomerDateRangePanel) SetHandler(f func(e gwu.Event, c myJson.Customer, syr, smo, eyr, emo int)) {
	c.handler = f
}

func NewCustomerDateRangePanel(cust []myJson.Customer, startYr int) *CustomerDateRangePanel {

	embPanel := gwu.NewHorizontalPanel()
	embPanel.SetCellPadding(3)
	embPanel.SetCellSpacing(3)

	p := CustomerDateRangePanel{Panel: embPanel}
	p.Add(NewCustomerListBox(cust))
	dateRangeTble := gwu.NewTable()
	dateRangeTble.SetCellSpacing(3)
	dateRangeTble.SetCellPadding(3)

	syrLb := getYearListBox(startYr)
	dateRangeTble.Add(syrLb, 0, 0)

	smonLb := getMonthListBox()
	dateRangeTble.Add(smonLb, 0, 1)

	eyrLb := getYearListBox(startYr)
	dateRangeTble.Add(eyrLb, 1, 0)

	emonLb := getMonthListBox()
	dateRangeTble.Add(emonLb, 1, 1)

	p.Add(dateRangeTble)

	b := gwu.NewButton("Submit")

	b.AddEHandlerFunc(func(e gwu.Event) {
		if p.handler == nil {
			return
		}
		c := p.CompAt(0).(CustomerListBox)
		t := p.CompAt(1).(gwu.Table)
		sy := t.CompAt(0, 0).(gwu.ListBox)
		sm := t.CompAt(0, 1).(gwu.ListBox)

		ey := t.CompAt(1, 0).(gwu.ListBox)
		em := t.CompAt(1, 1).(gwu.ListBox)

		//log.Println("in handle, yr index =", y, " Month index=", m)
		p.handler(e, c.GetSelectedCustomer(), sy.SelectedIdx()+startYr, sm.SelectedIdx()+1, ey.SelectedIdx()+startYr, em.SelectedIdx()+1)

		//log.Println("item selected: ", cl.GetSelectedCustomer())
	}, gwu.ETypeClick)
	p.Add(b)

	return &p

}

const (
	DownloadMediaTypeExcel = 1
	DownloadMediaTypePDF   = 2
)

type DownloadAsPanel struct {
	gwu.Panel
	urlTemplate string
	b64         string
	reqType     string
	server      string
	excelLink   gwu.Link
	pdfLink     gwu.Link
}

func NewDownloadAsPanel(server string, reqType string) *DownloadAsPanel {

	p := DownloadAsPanel{Panel: gwu.NewHorizontalPanel(), reqType: reqType, server: server}
	p.urlTemplate = "https://%s/export%s?t=%s&h=%s"

	excelButton := gwu.NewImage("Excel", "image/excel.png")

	p.excelLink = gwu.NewLink("", "")
	p.excelLink.SetComp(excelButton)

	pdfButton := gwu.NewImage("PDF", "image/pdf.png")
	p.pdfLink = gwu.NewLink("", "")
	p.pdfLink.SetComp(pdfButton)
	p.Add(gwu.NewLabel("Download Report: "))
	//p.AddHConsumer()
	p.Add(p.excelLink)
	p.Add(p.pdfLink)

	return &p
}

func (p *DownloadAsPanel) SetLink(b64 string) {
	p.excelLink.SetURL(fmt.Sprintf(p.urlTemplate, p.server, "/"+b64, "excel", p.reqType))
	p.pdfLink.SetURL(fmt.Sprintf(p.urlTemplate, p.server, "/"+b64, "pdf", p.reqType))
}

func getYearListBox(startYr int) gwu.ListBox {
	currYr := time.Now().Year()
	yrRange := make([]string, currYr-startYr+1)
	for i := range yrRange {
		yrRange[i] = strconv.Itoa(startYr + i)
	}
	lb := gwu.NewListBox(yrRange)
	set2Dropdown(lb)
	return lb
}

func getMonthListBox() gwu.ListBox {
	mons := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	lb := gwu.NewListBox(mons)
	set2Dropdown(lb)
	return lb
}

func set2Dropdown(l gwu.ListBox) {
	l.SetMulti(false)
	l.SetRows(1)
	l.SetSelected(0, true)
}
