package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	paymentToken := "428d7f3b2dc49117552ace464078557832c4ae4f";

	customer := map[string]interface{}{
		"name": "Gorbadoc Oldbuck",
		"cpf": "73504596295",
		"phone_number": "5144916523",
	}

	billingAddress := map[string]interface{} {
		"street": "Av JK",
		"number": 909,
		"neighborhood": "Bauxita",
		"zipcode": "35400000",
		"city": "Ouro Preto",
		"state": "MG",
	}

	body := map[string]interface{} {
		"payment": map[string]interface{} {
			"credit_card": map[string]interface{} {
				"billing_address": billingAddress,
				"payment_token": paymentToken,
				"customer": customer,
			},
		},
	}

	res, err := gn.DefineSubscriptionPayMethod(1, body) // no lugar do 1 coloque o subscription_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}