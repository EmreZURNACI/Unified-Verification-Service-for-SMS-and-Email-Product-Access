package Product

import (
	"fmt"
	"strings"

	c "ProductService/Connection"
	h "ProductService/Helpers"
	m "ProductService/Models"
)

var result string

func Products(bdy m.BodyProductRes) []m.Product {
	allProduct := []m.Product{}
	var query string = "SELECT * FROM product"
	if len(strings.Trim(bdy.Search, " ")) != 0 && bdy.Search != "" {
		query += fmt.Sprintf(" WHERE marka ILIKE '%%%s%%'", bdy.Search)
	}
	if bdy.Sorting != 0 {
		switch bdy.Sorting {
		case 1:
			query += " ORDER BY id ASC"
		case 2:
			query += " ORDER BY id DESC"
		}
	}
	if bdy.Offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", bdy.Offset)
	}
	if bdy.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", bdy.Limit)
	}
	row, err := c.Connection().Query(query)
	if err != nil {
		panic(err)
	}
	for row.Next() {
		var p m.Product
		err := row.Scan(&p.Id, &p.Marka, &p.Model, &p.IsletimSistemi)
		if err != nil {
			panic(err)
		}
		allProduct = append(allProduct, p)
	}
	return allProduct
}
func GetData(_id int) []m.Product {
	var p m.Product
	var query string = fmt.Sprintf("SELECT * FROM product WHERE id=%d", _id)
	row, err := c.Connection().Query(query)
	if err != nil {
		panic(err)
	}
	for row.Next() {
		var p m.Product
		err := row.Scan(&p.Id, &p.Marka, &p.Model, &p.IsletimSistemi)
		if err != nil {
			panic(err)
		}
	}
	return []m.Product{p}
}
func CreateProduct(_marka, _model, _os string) (bool, string) {
	c.Connection().QueryRow("SELECT createproduct($1,$2,$3);", _marka, _model, _os).Scan(&result)
	return h.ExtractStatuAndMessage(result)
}
func DeleteProduct(_id int) (bool, string) {
	c.Connection().QueryRow("SELECT deleteproduct($1);", _id).Scan(&result)
	return h.ExtractStatuAndMessage(result)
}
func UpdateProduct(_id int, _marka, _model, _isletimsistemi string) (bool, string) {
	c.Connection().QueryRow("Select updateproduct($1,$2,$3,$4)", _id, _marka, _model, _isletimsistemi).Scan(&result)
	return h.ExtractStatuAndMessage(result)
}
