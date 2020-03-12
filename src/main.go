package main

import (
	"config"
	"fmt"
	"models"
	"entities"
)

func CallFindAll() {
	db, err := config.GetMysqlDb()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(products)
			fmt.Println("Products List")
			for _, product := range products {
				fmt.Println("ID : ", product.ID)
				fmt.Println("Name : ", product.Name)
				fmt.Println("Price : ", product.Price)
				fmt.Println("Quantity : ", product.Quantity)
				fmt.Println("Status : ", product.Status)
				fmt.Println("--------------------------")
			}
		}
	}
}

func CallFindByID() {
	db, err := config.GetMysqlDb()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err2 := productModel.FindById(1)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println()
			fmt.Println("Product Info For ID = ",product.ID)
			fmt.Println("ID: ", product.ID)
			fmt.Println("Name: ", product.Name)
			fmt.Println("Price: ", product.Price)
			fmt.Println("Quantity: ", product.Quantity)
			fmt.Println("Status: ", product.Status)
		}
	}
}

func CallInsert(){
	db, err := config.GetMysqlDb() //khởi tạo error
	if err != nil { //khởi tạo lỗi
		fmt.Println(err)
	}else{
		productModel := models.ProductModel{
			Db: db,
		}
		product := entities.Product{
			ID : 3,
			Name: "thit bo",
			Price: 4.5,
			Quantity: 19,
			Status: true,
		}
		err := productModel.Create(&product) //truyền địa chỉ + data product
		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Println("last id: ",product.ID)
		}
	}
}

func CallUpdate(){
	db, err := config.GetMysqlDb() //khởi tạo error
	if err != nil { //khởi tạo lỗi
		fmt.Println(err)
	}else{
		productModel := models.ProductModel{
			Db: db,
		}
		product := entities.Product{
			ID : 3,
			Name: "thit bo nhat ban",
			Price: 4.5,
			Quantity: 19,
			Status: true,
		}
		rows,err := productModel.Update(product) 
		if err != nil {
			fmt.Println(err)
		}else{
			if rows > 0 {
				fmt.Println("xong")
			}
		}
	}
}

func CallDelete(){
	db, err := config.GetMysqlDb() //khởi tạo error
	if err != nil { //khởi tạo lỗi
		fmt.Println(err)
	}else{
		productModel := models.ProductModel{
			Db: db,
		}
		rows,err := productModel.Delete(3) 
		if err != nil {
			fmt.Println(err)
		}else{
			if rows > 0 {
				fmt.Println("xong")
			}
		}
	}
}


func main() {
	CallFindAll()
	//CallFindById()
	//CallInsert()
	//CallUpdate()
	//CallDelete()
}
