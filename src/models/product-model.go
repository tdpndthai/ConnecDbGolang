package models

import (
	"database/sql"
	"entities"
)

type ProductModel struct {
	Db *sql.DB 
}
/*FindAl*/
func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product")
	if err != nil {
		return nil, err
	} else {
		products := []entities.Product{} //khởi tạo mảng struc Product rỗng,mảng chứa các product 
		for rows.Next() {//khởi tạo các giá trị cần lấy ban đầu
			var id int64
			var name string
			var price float64
			var quantity int
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status) //gán từng giá trị lấy được trong rows vào id,name
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{id, name, price, quantity, status} //từng product tìm được trong row
				products = append(products, product)
			}
		}
		return products, nil
	}
}

/*FindByID*/
func (productModel ProductModel) FindById(id int64) (entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product where id =?", id)
	if err != nil {
		return entities.Product{}, err
	} else {
		product := entities.Product{}
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return entities.Product{}, err2
			} else {
			 	product = entities.Product{id, name, price, quantity, status}
			}
		}
		return product, nil
	}
}

/*Create*/
func (productModel ProductModel) Create(product *entities.Product) error { // triển khai interface error,con trỏ kiểu Product
	result, err := productModel.Db.Exec("insert into product (id,name,price,quantity,status) value (?,?,?,?,?)", product.ID,product.Name,product.Price,product.Quantity,product.Status )
	if err != nil {
		return err
	}else{
		product.ID,_ = result.LastInsertId() //sau khi tạo mới lấy id chèn cuối cùng
		return nil
	}
}

/*Update type:ProductModel,params: entities.Product,trả về số dòng và lỗi*/
func (productModel ProductModel) Update(product entities.Product) (int64,error){  
	result, err := productModel.Db.Exec("update product set name = ?,price=?,quantity =?,status=? where id=?",product.Name,product.Price,product.Quantity,product.Status,product.ID)
	if err != nil {
		return 0,err
	}else{
		//row := result.RowsAffected //trả về số dòng bị ảnh hưởng các truy vấn SELECT, INSERT, UPDATE, REPLACE, hoặc DELETE 
		return result.RowsAffected()
	}
}

func (productModel ProductModel) Delete(id int64) (int64,error){  
	result, err := productModel.Db.Exec("delete from product where id = ?",id)
	if err != nil {
		return 0,err
	}else{
		//row := result.RowsAffected //trả về số dòng bị ảnh hưởng các truy vấn SELECT, INSERT, UPDATE, REPLACE, hoặc DELETE 
		return result.RowsAffected()
	}
}
