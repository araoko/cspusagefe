package rest

import (
	"encoding/json"
	"fmt"

	myJson "github.com/araoko/cspusage/model/json"
	"github.com/go-resty/resty/v2"
)

type RestClient struct {
	server string
	port   int
	cert   string
	key    string
}

func NewRestClient(server string, port int, cert, key string) RestClient {
	return RestClient{
		server: server,
		port:   port,
		cert:   cert,
		key:    key,
	}
}

func (r RestClient) URL() string {
	return fmt.Sprintf("https://%s:%d/", r.server, r.port)
}

func (r RestClient) Server() string {
	return fmt.Sprintf("%s:%d", r.server, r.port)
}

func (r RestClient) GetCustomers() ([]myJson.Customer, error) {
	url := r.URL() + "list"
	result := []myJson.Customer{}
	err := r.fetch(url, nil, &result)
	return result, err
}

func (r RestClient) GetCustomerMontlyBill(c myJson.Customer, yr, mo int) (myJson.CustomerMonthlyBill, error) {
	url := r.URL() + "customermonthly"
	result := myJson.CustomerMonthlyBill{}
	cyd := myJson.CustomerIDAndDate{
		CustomerId: c.CustomerId,
		Date: myJson.YearMonth{
			Year:  yr,
			Month: mo,
		},
	}
	err := r.fetch(url, cyd, &result)
	return result, err
}

//GetCustomerMonthlyCostPerSub fetches a cutomer usage for a specific month
//breaking down the cost per subscription. it does this by calling the api server
// and returning the json response
func (r RestClient) GetCustomerMonthlyCostPerSub(c myJson.Customer, yr, mo int) (myJson.CustomerMonthlyCostPerSub, error) {
	url := r.URL() + "customermonthlypersub"
	result := myJson.CustomerMonthlyCostPerSub{}
	cyd := myJson.CustomerIDAndDate{
		CustomerId: c.CustomerId,
		Date: myJson.YearMonth{
			Year:  yr,
			Month: mo,
		},
	}
	err := r.fetch(url, cyd, &result)
	return result, err
}

func (r RestClient) GetMontlyBill(yr, mo int) (myJson.MonthlyBill, error) {
	url := r.URL() + "monthly"
	result := myJson.MonthlyBill{}
	cyd := myJson.YearMonth{
		Year:  yr,
		Month: mo,
	}
	err := r.fetch(url, cyd, &result)
	return result, err
}

func (r RestClient) GetCustomerRangeBill(c myJson.Customer, syr, smo, eyr, emo int) (myJson.CustomerRangeBill, error) {
	url := r.URL() + "customerbillrange"
	result := myJson.CustomerRangeBill{}
	cyd := myJson.CustomerIDAndDateRange{
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
	err := r.fetch(url, cyd, &result)
	return result, err
}

//GetCustomerRangeCostPerSub fetches a cutomer usage for a specified period
//breaking down the cost per subscription. it does this by calling the api server
// and returning the json response
func (r RestClient) GetCustomerRangeCostPerSub(c myJson.Customer, syr, smo, eyr, emo int) (myJson.CustomerRangeCostPerSub, error) {
	url := r.URL() + "customercostrangepersub"
	result := myJson.CustomerRangeCostPerSub{}
	cyd := myJson.CustomerIDAndDateRange{
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
	err := r.fetch(url, cyd, &result)
	return result, err
}

func (r RestClient) GetRangeBill(syr, smo, eyr, emo int) (myJson.RangeBill, error) {

	url := r.URL() + "billrange"
	result := myJson.RangeBill{}
	cyd := myJson.YearMonthRange{
		StartDate: myJson.YearMonth{
			Year:  syr,
			Month: smo,
		},
		EndDate: myJson.YearMonth{
			Year:  eyr,
			Month: emo,
		},
	}
	err := r.fetch(url, cyd, &result)
	return result, err
}

func (r RestClient) GetTrend(syr, smo, eyr, emo int) (myJson.MonthlyTrend, error) {

	url := r.URL() + "trend"
	result := myJson.MonthlyTrend{}
	cyd := myJson.YearMonthRange{
		StartDate: myJson.YearMonth{
			Year:  syr,
			Month: smo,
		},
		EndDate: myJson.YearMonth{
			Year:  eyr,
			Month: emo,
		},
	}
	err := r.fetch(url, cyd, &result)
	return result, err
}

func (r RestClient) GetCustomerTrend(cid string, syr, smo, eyr, emo int) (myJson.CustomerMonthlyTrend, error) {

	url := r.URL() + "customertrend"
	result := myJson.CustomerMonthlyTrend{}
	cyd := myJson.CustomerIDAndDateRange{
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
		CustomerId: cid,
	}
	err := r.fetch(url, cyd, &result)
	return result, err
}

func (r RestClient) GetCustomerPerSubTrend(cid string, syr, smo, eyr, emo int) (myJson.CustomerMonthlyPerSubTrend, error) {

	url := r.URL() + "customerpersubtrend"
	result := myJson.CustomerMonthlyPerSubTrend{}
	cyd := myJson.CustomerIDAndDateRange{
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
		CustomerId: cid,
	}
	err := r.fetch(url, cyd, &result)
	return result, err
}

func (r RestClient) Auth(userName, password string) (myJson.Login, error) {
	url := r.URL() + "auth"
	result := myJson.Login{}
	input := myJson.Login{
		UserName: userName,
		Password: password,
	}
	err := r.fetch(url, input, &result)
	return result, err
}

//fetch does the actual rest call to the rest server.
//it takes the server url, optional object to marshal as json body
//and a pointer to the object to store the json response
func (r RestClient) fetch(url string, body interface{}, result interface{}) error {
	client := resty.New()
	var resp *resty.Response
	var err error
	if body == nil {
		resp, err = client.R().SetResult(result).Get(url)

	} else {
		js, err := json.Marshal(body)
		if err != nil {
			return err
		}
		resp, err = client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(js).
			SetResult(result).
			Post(url)
	}

	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("Rest client Error: %v", resp.Error())
	}
	return nil
}
