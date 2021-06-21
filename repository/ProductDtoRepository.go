package repository

import (
	"errors"
	"ocg.com/product/model/dto"
)

type ProductDtoRepo struct {
	productsDto map[int64]*dto.ProductDto
	autoID      int64
}

var ProductsDto ProductDtoRepo

func init() {
	ProductsDto = ProductDtoRepo{autoID: 0}
	ProductsDto.productsDto = make(map[int64]*dto.ProductDto)
}

func (r *ProductDtoRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ProductDtoRepo) CreateProductDto(productDto *dto.ProductDto) *dto.ProductDto {
	nextID := r.getAutoID()
	productDto.Id = nextID
	r.productsDto[nextID] = productDto
	return productDto
}

func (r *ProductDtoRepo) FindProductDtoById(id int64) (*dto.ProductDto, error) {
	if productDto, ok := r.productsDto[id]; ok {
		return productDto, nil
	} else {
		return nil, errors.New("product not found")
	}
}
