package main

import (
	"net/http"

	"github.com/SaladkevichM/go-cleanarchitecture/src/infrastructure"
	"github.com/SaladkevichM/go-cleanarchitecture/src/interfaces"
	"github.com/SaladkevichM/go-cleanarchitecture/src/usecases"
)

func main() {

	dbHandler := infrastructure.NewSqliteHandler("/var/tmp/production.sqlite")

	handlers := make(map[string]interfaces.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbCustomerRepo"] = dbHandler
	handlers["DbItemRepo"] = dbHandler
	handlers["DbOrderRepo"] = dbHandler

	orderInteractor := new(usecases.OrderInteractor)
	orderInteractor.UserRepository = interfaces.NewDbUserRepo(handlers)
	orderInteractor.ItemRepository = interfaces.NewDbItemRepo(handlers)
	orderInteractor.OrderRepository = interfaces.NewDbOrderRepo(handlers)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.OrderInteractor = orderInteractor

	http.HandleFunc("/all_orders", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowOrders(res, req)
	})

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowOrder(res, req)
	})

	http.ListenAndServe(":8080", nil)
}
