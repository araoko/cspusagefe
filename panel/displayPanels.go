package panel

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	myJson "github.com/araoko/cspusage/model/json"
	"github.com/icza/gowut/gwu"
)

func PanelOfCustomers(customers []myJson.Customer) gwu.Panel {
	p := gwu.NewPanel()
	tbl := gwu.NewTable()
	tbl.Add(gwu.NewLabel("Customer ID"), 0, 0)
	tbl.Add(gwu.NewLabel("Customer Name"), 0, 1)
	tbl.Add(gwu.NewLabel("Other Name(s)"), 0, 2)
	cfmt := tbl.RowFmt(0)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	for i, c := range customers {
		tbl.Add(gwu.NewLabel(c.CustomerId), i+1, 0)
		tbl.Add(gwu.NewLabel(c.CustomerCompanyName), i+1, 1)
		tbl.Add(gwu.NewLabel(c.FormerNames), i+1, 2)
	}
	tbl.SetBorder(1)
	//tbl.SetCellSpacing(10)
	tbl.SetCellPadding(5)
	p.Add(tbl)
	return p

}

func PanelOfCustomerMonthyUsage(cust myJson.CustomerMonthlyBill) gwu.Panel {
	p := gwu.NewPanel()
	customerInfo := customerInfoTable(cust.Owner)
	p.Add(customerInfo)
	p.AddVSpace(5)
	periodInfo := dateInfoTable(cust.Date)
	p.Add(periodInfo)
	p.AddVSpace(10)
	resultTable := monthlyBillItemsTable(cust.LineItems)
	p.Add(resultTable)
	return p
}

func PanelOfCustomerMonthyUsagePerSub(cust myJson.CustomerMonthlyCostPerSub) gwu.Panel {
	p := gwu.NewPanel()
	customerInfo := customerInfoTable(cust.Owner)
	p.Add(customerInfo)
	p.AddVSpace(5)
	periodInfo := dateInfoTable(cust.Date)
	p.Add(periodInfo)
	p.AddVSpace(10)
	resultTable := subscriptionCostItemsTable(cust.LineItems)
	p.Add(resultTable)
	return p
}

func PanelOfMonthyUsage(cust myJson.MonthlyBill) gwu.Panel {
	p := gwu.NewPanel()
	periodInfo := dateInfoTable(cust.Date)
	p.Add(periodInfo)
	p.AddVSpace(10)
	tPanel := gwu.NewTabPanel()
	sTable := customerCostTable(cust.Summary.Customers)
	sTable.Style().SetFullWidth()
	tPanel.AddString("Summary", sTable)

	for _, c := range cust.CustomerMonthlyBills {
		cTable := monthlyBillItemsTable(c.LineItems)
		cTable.Style().SetFullWidth()
		tPanel.AddString(c.Owner.CustomerCompanyName, cTable)

	}
	p.Add(tPanel)
	return p
}

func PanelOfRangeUsage(cust myJson.RangeBill) gwu.Panel {
	p := gwu.NewPanel()
	periodInfo := dateRangeInfoTable(cust.DateRange)
	p.Add(periodInfo)
	p.AddVSpace(10)
	tPanel := gwu.NewTabPanel()
	sTable := customerCostTable(cust.Summary.Customers)
	sTable.Style().SetFullWidth()
	tPanel.AddString("Summary", sTable)

	for _, c := range cust.CustomerRangeBills {
		cTable := monthlyBillItemsTable(c.LineItems)
		cTable.Style().SetFullWidth()
		tPanel.AddString(c.Owner.CustomerCompanyName, cTable)

	}
	p.Add(tPanel)
	return p
}

func PanelOfMonthlyTrend(cust myJson.MonthlyTrend) gwu.Panel {
	p := gwu.NewPanel()
	periodInfo := dateRangeInfoTable(cust.DateRange)
	p.Add(periodInfo)
	p.AddVSpace(10)
	tPanel := gwu.NewTabPanel()
	sTable := dataCostItemsTable(cust.Summary)
	sTable.Style().SetFullWidth()
	tPanel.AddString("Summary", sTable)

	for _, c := range cust.Trend {
		cTable := dataCostItemsTable(c.Trend)
		cTable.Style().SetFullWidth()
		tPanel.AddString(c.Owner.CustomerCompanyName, cTable)

	}
	p.Add(tPanel)
	return p
}

func PanelOfCustomerMonthlyPerSubTrend(cust myJson.CustomerMonthlyPerSubTrend) gwu.Panel {
	p := gwu.NewPanel()
	custInfor := customerInfoTable(cust.Owner)
	p.Add(custInfor)
	p.AddVSpace(5)
	periodInfo := dateRangeInfoTable(cust.DateRange)
	p.Add(periodInfo)
	p.AddVSpace(10)
	resultTable := dateCostSubsItemsTable(cust.Trend)
	p.Add(resultTable)
	return p
}

func PanelOfCustomerMonthlyTrend(cust myJson.CustomerMonthlyTrend) gwu.Panel {
	p := gwu.NewPanel()
	custInfor := customerInfoTable(cust.Owner)
	p.Add(custInfor)
	p.AddVSpace(5)
	periodInfo := dateRangeInfoTable(cust.DateRange)
	p.Add(periodInfo)
	p.AddVSpace(10)
	resultTable := dataCostItemsTable(cust.Trend)
	p.Add(resultTable)
	return p
}

func PanelOfCustomerRangeUsage(cust myJson.CustomerRangeBill) gwu.Panel {
	p := gwu.NewPanel()
	infoTable := customerInfoTable(cust.Owner)
	p.Add(infoTable)
	p.AddVSpace(5)
	infoLabel := dateRangeInfoTable(cust.DateRange)
	p.Add(infoLabel)
	p.AddVSpace(10)
	resultTable := monthlyBillItemsTable(cust.LineItems)
	p.Add(resultTable)
	return p
}

func PanelOfCustomerRangeUsagePerSub(cust myJson.CustomerRangeCostPerSub) gwu.Panel {
	p := gwu.NewPanel()
	infoTable := customerInfoTable(cust.Owner)
	p.Add(infoTable)
	p.AddVSpace(5)
	infoLabel := dateRangeInfoTable(cust.DateRange)
	p.Add(infoLabel)
	p.AddVSpace(10)
	resultTable := subscriptionCostItemsTable(cust.LineItems)
	p.Add(resultTable)
	return p
}

func dateRangeInfoTable(date myJson.YearMonthRange) gwu.Table {
	infoTable := gwu.NewTable()
	infoTable.SetCellPadding(5)
	infoTable.SetBorder(1)

	startLabel := gwu.NewLabel("Start Date")
	startLabel.Style().SetFontWeight("bold")
	endLabel := gwu.NewLabel("End Date")
	endLabel.Style().SetFontWeight("bold")

	infoTable.Add(startLabel, 0, 0)
	infoTable.Add(dateInfoTable(date.StartDate), 0, 1)
	infoTable.Add(endLabel, 1, 0)
	infoTable.Add(dateInfoTable(date.EndDate), 1, 1)

	return infoTable
}

func dateInfoTable(date myJson.YearMonth) gwu.Table {
	infoTable := gwu.NewTable()
	infoTable.Add(gwu.NewLabel(fmt.Sprintf("Year:\t%d", date.Year)), 0, 0)
	infoTable.Add(gwu.NewLabel(fmt.Sprintf("Month:\t%s", time.Month(date.Month).String())), 0, 1)
	infoTable.SetCellPadding(5)
	return infoTable
}

func customerInfoTable(customer myJson.Customer) gwu.Table {
	infoTable := gwu.NewTable()
	infoTable.Add(gwu.NewLabel("Customer ID"), 0, 0)
	infoTable.Add(gwu.NewLabel("Customer Name"), 1, 0)
	infoTable.Add(gwu.NewLabel("Other Name(s)"), 2, 0)
	infoTable.Add(gwu.NewLabel(customer.CustomerId), 0, 1)
	infoTable.Add(gwu.NewLabel(customer.CustomerCompanyName), 1, 1)
	infoTable.Add(gwu.NewLabel(customer.FormerNames), 2, 1)
	infoTable.SetBorder(1)
	infoTable.SetCellPadding(5)
	return infoTable
}

func customerCostTable(ccs []myJson.CustomerCostItem) gwu.Table {
	resultTable := gwu.NewTable()

	resultTable.SetBorder(1)
	resultTable.SetCellPadding(5)
	resultTable.Add(gwu.NewLabel("Customer Company Name"), 0, 0)
	resultTable.Add(gwu.NewLabel("Cost($)"), 0, 1)

	cfmt := resultTable.RowFmt(0)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	var total float32
	totalRowNo := len(ccs) + 1

	for i, li := range ccs {
		resultTable.Add(gwu.NewLabel(li.Owner.CustomerCompanyName), i+1, 0)
		resultTable.Add(costLabel(li.Cost), i+1, 1)
		total += li.Cost
	}

	resultTable.Add(gwu.NewLabel("Total"), totalRowNo, 0)
	resultTable.Add(costLabel(total), totalRowNo, 1)
	cfmt = resultTable.RowFmt(totalRowNo)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	return resultTable
}

func monthlyBillItemsTable(lineItems []myJson.SubscriptionServiceCostItem) gwu.Table {
	resultTable := gwu.NewTable()

	resultTable.SetBorder(1)
	resultTable.SetCellPadding(5)

	resultTable.Add(gwu.NewLabel("Suscription"), 0, 0)
	resultTable.Add(gwu.NewLabel("ServiceName - Type"), 0, 1)
	resultTable.Add(gwu.NewLabel("Cost"), 0, 2)

	cfmt := resultTable.RowFmt(0)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	var total float32
	totalRowNo := len(lineItems) + 1
	for i, li := range lineItems {
		resultTable.Add(gwu.NewLabel(li.Suscription), i+1, 0)
		resultTable.Add(gwu.NewLabel(li.ServiceNameAndType), i+1, 1)
		resultTable.Add(costLabel(li.Cost), i+1, 2)
		total += li.Cost
	}
	resultTable.Add(gwu.NewLabel("Total"), totalRowNo, 1)
	resultTable.Add(costLabel(total), totalRowNo, 2)
	cfmt = resultTable.RowFmt(totalRowNo)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	return resultTable
}

func dataCostItemsTable(lineItems []myJson.DateCostItem) gwu.Table {
	resultTable := gwu.NewTable()

	resultTable.SetBorder(1)
	resultTable.SetCellPadding(5)

	resultTable.Add(gwu.NewLabel("Year"), 0, 0)
	resultTable.Add(gwu.NewLabel("Month"), 0, 1)
	resultTable.Add(gwu.NewLabel("Cost($)"), 0, 2)

	cfmt := resultTable.RowFmt(0)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	var total float32
	totalRowNo := len(lineItems) + 1
	for i, li := range lineItems {
		resultTable.Add(gwu.NewLabel(strconv.Itoa(li.Date.Year)), i+1, 0)
		resultTable.Add(gwu.NewLabel(time.Month(li.Date.Month).String()), i+1, 1)
		resultTable.Add(costLabel(li.Cost), i+1, 2)
		total += li.Cost
	}
	resultTable.Add(gwu.NewLabel("Total"), totalRowNo, 1)
	resultTable.Add(costLabel(total), totalRowNo, 2)
	cfmt = resultTable.RowFmt(totalRowNo)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	return resultTable
}

func dateCostSubsItemsTable(lineItems []myJson.DateCostSubItem) gwu.Table {
	resultTable := gwu.NewTable()

	resultTable.SetBorder(1)

	resultTable.SetCellPadding(5)
	csa := lineItems[0].CostPerSubs

	resultTable.Add(gwu.NewLabel("Year"), 0, 0)
	resultTable.Add(gwu.NewLabel("Month"), 0, 1)
	for i, cs := range csa {
		resultTable.Add(gwu.NewLabel(cs.Subscription), 0, 2+i)
	}
	totalIndex := 2 + len(csa)
	resultTable.Add(gwu.NewLabel("Total"), 0, totalIndex)

	cfmt := resultTable.RowFmt(0)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	var total float32
	totalRowNo := len(lineItems) + 1
	colTolal := make([]float32, len(csa))
	for i, li := range lineItems {
		resultTable.Add(gwu.NewLabel(strconv.Itoa(li.Date.Year)), i+1, 0)
		resultTable.Add(gwu.NewLabel(time.Month(li.Date.Month).String()), i+1, 1)
		var rowTotal float32
		for j, cs := range li.CostPerSubs {
			cst := cs.Cost
			rowTotal += cst
			colTolal[j] += cst
			resultTable.Add(costLabel(cst), i+1, 2+j)
		}
		resultTable.Add(costLabel(rowTotal), i+1, totalIndex)
	}
	resultTable.Add(gwu.NewLabel("Total"), totalRowNo, 1)
	for j := range csa {
		total += colTolal[j]
		resultTable.Add(costLabel(colTolal[j]), totalRowNo, 2+j)
	}
	resultTable.Add(costLabel(total), totalRowNo, totalIndex)

	cfmt = resultTable.RowFmt(totalRowNo)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	return resultTable
}

func subscriptionCostItemsTable(lineItems []myJson.SubscriptionCostItem) gwu.Table {
	resultTable := gwu.NewTable()

	resultTable.SetBorder(1)
	resultTable.SetCellPadding(5)

	resultTable.Add(gwu.NewLabel("Suscription"), 0, 0)
	resultTable.Add(gwu.NewLabel("Cost"), 0, 1)

	cfmt := resultTable.RowFmt(0)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	var total float32
	totalRowNo := len(lineItems) + 1
	for i, li := range lineItems {
		resultTable.Add(gwu.NewLabel(li.Suscription), i+1, 0)
		costValueLabel := costLabel(li.Cost)
		costValueLabel.Style().Set("text-align", gwu.HARight)
		resultTable.Add(costValueLabel, i+1, 1)
		total += li.Cost
	}
	totalLabel := gwu.NewLabel("Total")
	totalLabel.Style().Set("text-align", gwu.HARight)
	resultTable.Add(totalLabel, totalRowNo, 0)
	totalValueLabel := costLabel(total)
	totalValueLabel.Style().Set("text-align", gwu.HARight)
	resultTable.Add(totalValueLabel, totalRowNo, 1)
	cfmt = resultTable.RowFmt(totalRowNo)
	cfmt.Style().SetFontSize("20")
	cfmt.Style().SetFontWeight("Bold")
	return resultTable
}

func costLabel(cost float32) gwu.Label {
	s := strconv.FormatFloat(float64(cost), 'f', 2, 32)
	return gwu.NewLabel(kSeparator(s))

}

func kSeparator(s string) string {
	s2 := strings.Split(s, ".")
	s1 := s2[0]
	l := len(s1)
	if l <= 3 {
		return s
	}
	m := l % 3
	ss := make([]string, 0)
	if m == 0 {
		m = 3
	}
	ss = append(ss, s1[0:m])
	a := m

	for {
		if a == l {
			break
		}

		ss = append(ss, ",", s1[a:a+3])

		a += 3
	}
	if len(s2[1]) != 0 {
		ss = append(ss, ".", s2[1])
	}

	return strings.Join(ss, "")
}
