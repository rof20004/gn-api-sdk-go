package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)

	

	body := map[string]interface{} {
		
		"pix": map[string]interface{}{
			    "receberSemChave": false,
				"chaves": map[string]interface{} {		
					"177f47c8-0ebd-4b9b-a941-f54babba72ba":map[string]interface{} {
						"recebimento":map[string]interface{} {
							"txidObrigatorio": false,
							"qrCodeEstatico":map[string]interface{} {
							"recusarTodos": false,

							},

						},

					},

			},
			
		},
		
	}

	res, err := gn.UpdateAccountConfig(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}