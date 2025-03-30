package main

import (
	"fmt"
	"log"
	"time"
	"encoding/json"
)


type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductManager interface {
	GetName() string
	GetPrice() float64
}

func (p Product) GetName() string {
	return p.Name
}

func (p Product) GetPrice() float64 {
	return p.Price
}


func GetData(pm ProductManager) {
	fmt.Printf("Producto: %s, Precio: $%.2f\n", 
        pm.GetName(), 
        pm.GetPrice())
}


var products = []Product{
	{Id: 1, Name: "Laptop", Price: 999.99},
	{Id: 2, Name: "Mouse", Price: 25.50},
}

func FindProductByIDEnhanced(id int) (Product, error) {
	for _, p := range products {
		if p.Id == id {
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("producto con ID %d no existe", id)
}

func getProductFromDB(id int) (Product, error) {
	time.Sleep(100 * time.Millisecond) // Simula latencia
	return Product{Id: id, Name: "Laptop", Price: 999.99}, nil
}

func getStockFromExternalService(id int) (int, error) {
	time.Sleep(150 * time.Millisecond) // Simula latencia mayor
	return 42, nil
}

func fetchProductData(id int) (Product, int, error) {
	// Creamos channels para los resultados
	productChan := make(chan Product)
	stockChan := make(chan int)
	errChan := make(chan error)

	// Lanzamos las goroutines
	go func() {
			product, err := getProductFromDB(id)
			if err != nil {
					errChan <- err
					return
			}
			productChan <- product
	}()

	go func() {
			stock, err := getStockFromExternalService(id)
			if err != nil {
					errChan <- err
					return
			}
			stockChan <- stock
	}()

	// Esperamos los resultados
	var product Product
	var stock int
	var err error

	for i := 0; i < 2; i++ {
			select {
			case p := <-productChan:
					product = p
			case s := <-stockChan:
					stock = s
			case e := <-errChan:
					err = e
					return Product{}, 0, err
			}
	}

	return product, stock, nil
}

func main() {

	product := Product{
		Id: 1,
		Name: "Product 1",
		Price: 10.99,
	}


	var productManager ProductManager = Product{
		Id: 2,
		Name: "Product 2",
		Price: 20.99,
	}

	GetData(product)

	fmt.Println(product)
	fmt.Println(productManager)

	jsonProduct, err1 := json.Marshal(product)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("JSON:", string(jsonProduct))

	_, err := FindProductByIDEnhanced(99)
	if err != nil {
		fmt.Println("Error:", err) // Imprime: Error: producto con ID 99 no existe
	}

	start := time.Now()
	product, stock, err := fetchProductData(1)
	if err != nil {
			log.Fatal(err)
	}
	
	fmt.Printf("Producto: %+v, Stock: %d\n", product, stock)
	fmt.Printf("Tiempo total: %v\n", time.Since(start)) 



}
