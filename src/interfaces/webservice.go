package interfaces

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/SaladkevichM/go-cleanarchitecture/src/domain"

	"github.com/SaladkevichM/go-cleanarchitecture/src/usecases"
)

type OrderInteractor interface {
	Items(userId, orderId int) ([]usecases.Item, error)
	GetAll() ([]domain.Order, error)
	Add(userId, orderId, itemId int) error
}

type WebserviceHandler struct {
	OrderInteractor OrderInteractor
}

func (handler WebserviceHandler) ShowOrders(res http.ResponseWriter, req *http.Request) {

	orders, _ := handler.OrderInteractor.GetAll()
	for _, order := range orders {
		io.WriteString(res, fmt.Sprintf("order id: %d\n", order.Id))

	}
}

func (handler WebserviceHandler) ShowOrder(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))
	orderId, _ := strconv.Atoi(req.FormValue("orderId"))
	items, _ := handler.OrderInteractor.Items(userId, orderId)
	for _, item := range items {
		io.WriteString(res, fmt.Sprintf("item id: %d\n", item.Id))
		io.WriteString(res, fmt.Sprintf("item name: %v\n", item.Name))
		io.WriteString(res, fmt.Sprintf("item value: %f\n", item.Value))
	}
}
