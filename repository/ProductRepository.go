package repository

import (
	"errors"
	"ocg.com/product/model/entity"
	"time"
)

type ProductRepo struct {
	products map[int64]*entity.Product
	autoID int64
}

var Products ProductRepo

func init(){
	Products = ProductRepo{autoID: 0}
	Products.products = make(map[int64]*entity.Product)
}

func (r *ProductRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ProductRepo) CreateProduct(product *entity.Product) *entity.Product {
	nextID := r.getAutoID()
	product.Id = nextID
	product.CreatedAt = time.Now().Unix()
	product.ModifiedAt = time.Now().Unix()
	product.Rating = 0
	r.products[nextID] = product
	return product
}

func (r *ProductRepo) FindAllProduct() map[int64]*entity.Product {
	return r.products
}

func (r *ProductRepo) FindProductById(id int64) (*entity.Product, error) {
	if product, ok := r.products[id]; ok {
		return product, nil
	} else {
		return nil, errors.New("product not found")
	}
}

func (r *ProductRepo) UpdateProductById(id int64, product *entity.Product) (*entity.Product, error) {
	if _, ok := r.products[id]; ok {
		product.Id = id
		product.CreatedAt = r.products[id].CreatedAt
		product.ModifiedAt = time.Now().Unix()
		r.products[id] = product
		return product, nil
	} else {
		return nil, errors.New("product not found")
	}
}

func (r *ProductRepo) DeleteProductById(id int64) error {
	if _, ok := r.products[id]; ok {
		delete(r.products, id)
		return nil
	} else {
		return errors.New("product not found")
	}
}

