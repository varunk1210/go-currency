package visa

import (
	"encoding/json"
	"fmt"
	"io"
	"mycurrencynotifier/datatypes"
	"net/http"
)

func GetCurrency(appId string) (*datatypes.Rate, error) {

	apiUrl := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s", appId)
	var response VisaResponse
	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyInformationBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		error := json.Unmarshal(bodyInformationBytes, &response)
		if error != nil {
			return nil, error
		}
	} else {
		return nil, fmt.Errorf("status Code received is %v", res.StatusCode)
	}

	// price, err := strconv.ParseFloat(, 64)
	// if err != nil {
	// 	return nil, err
	// }

	rate := datatypes.Rate{Currency: "INR", Price: response.Rates.Inr}
	return &rate, nil

}
