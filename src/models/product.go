package models

import ()

type Product struct {
	id          int
	imageUrl    string
	name        string
	typ			string
	description string
	price       float32
}

func (this *Product) Id() int{
	return this.id
}
func (this *Product) setId(value int){
	this.id = value
}
func (this *Product) ImageUrl() string{
	return this.imageUrl
}
func (this *Product) setImateUrl(value string){
	this.imageUrl = value
}
func (this *Product) Name() string{
	return this.name
}
func (this *Product) setName(value string){
		this.name = value
}
func (this *Product) Typ() string{
	return this.typ
}
func (this *Product) setTyp(value string){
	this.typ = value
}
func (this *Product) Description() string{
	return this.description
}
func (this *Product) setDescription(value string){
	this.description = value
}
func (this *Product) Price() float32{
	return this.price
}
func (this *Product) setPrice(value float32){
	this.price = value
}
