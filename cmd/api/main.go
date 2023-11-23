package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xandreafonso/gogo/internal/domain/entity"
)

func main() {
	// router := chi.NewRouter() // chi mantém compatibilidade com os handlers feitos de forma padrão
	// router.Use(middleware.Logger)
	// router.Get("/order", OrderHandler)

	// http.HandleFunc("/order", OrderHandler) // forma padrão de adicionar endpoints
	// http.ListenAndServe(":8888", nil) // forma padrão sem roteador
	// http.ListenAndServe(":8888", router)

	e := echo.New()
	e.GET("/order", OrderHandler)

	e.Logger.Fatal(e.Start(":8888"))
}

// func OrderHandler(w http.ResponseWriter, r *http.Request) {
// 	// if r.Method != http.MethodGet { // if necessário para quando não se tem roteadores
// 	// 	w.WriteHeader(http.StatusMethodNotAllowed)
// 	// 	return
// 	// }

// 	order, err := entity.NewOrder("1", 10.0, 1.0)

// 	if err != nil {
// 		panic(err)
// 	}

// 	err = order.CalculateFinalPrice()

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}

// 	json.NewEncoder(w).Encode(order)
// }

func OrderHandler(c echo.Context) error {
	order, err := entity.NewOrder("1", 10.0, 1.0)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error()) // return c.JSON(400, err)
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		return c.JSON(400, err) // return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
