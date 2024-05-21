package model

import "gorm.io/gorm"

type BlindBox struct {
	gorm.Model
	BlindType       int    `json:"type" gorm:"column:blind_type"`
	Sale            int    `json:"sale" gorm:"default:0;column:sale"`
	Number          int    `json:"number" gorm:"default:0;column:number"`
	Original        string `json:"original" gorm:"column:original"`
	Discount        int    `json:"discount" gorm:"default:0" gorm:"column:discount"`
	CarbonReduction int    `json:"carbon_reduction" gorm:"column:carbon_reduction"`
	Image           string `json:"image" gorm:"column:image"`
	Status          int    `json:"status" gorm:"default:0" gorm:"column:status"`
}
type Business struct {
	//ID           int     `gorm:"primaryKey" json:"id"`
	gorm.Model
	ManagerID    string  `json:"manager_id" gorm:"column:manager_id"`
	Name         string  `json:"name" gorm:"column:name"`
	Avatar       string  `json:"avatar" gorm:"column:avatar"`
	Address      string  `json:"address" gorm:"column:address"`
	Openinghours string  `json:"opening_hours" gorm:"column:opening_hours"`
	Description  string  `json:"description" gorm:"column:description"`
	Image        string  `json:"image" gorm:"column:image"`
	Score        float64 `json:"score" gorm:"column:score"`
	Tele         string  `json:"tele" gorm:"column:tele"`
}
type BusinessScore struct {
	ID        int     `json:"id" gorm:"primaryKey;column:id"`
	ScoreNum  float64 `json:"score_num" gorm:"column:score_num"`
	PeopleNum int     `json:"people_num" gorm:"column:people_num"`
}
type BusinessAndBox struct {
	gorm.Model
	BusinessID int `json:"business_id" gorm:"column:business_id"`
	BoxID      int `json:"box_id" gorm:"column:box_id"`
}

func NewBlindBox() BlindBox {
	return BlindBox{}
}

func NewBusiness() Business {
	return Business{}
}
func NewBusinessAndBox() BusinessAndBox {
	return BusinessAndBox{}
}
func (b *BlindBox) AddOneTypeBox(btype int, num int, original string, caebonreduction int, image string) {
	b.BlindType = btype
	b.Number = num
	b.Original = original
	b.CarbonReduction = caebonreduction
	b.Image = image
}
func (b *BlindBox) UpdateSale(sale int) {
	b.Sale += sale
}
func (b *BlindBox) UpdateBoxNum(num int) {
	b.Number += num
}
func (b *BlindBox) UpdateBoxStatus(satatus int) {
	b.Status = satatus
}
func (b *BlindBox) UpdateDiscount(discount int) {
	b.Discount = discount
}
func (b *BlindBox) UpdateOriginal(original string) {
	b.Original = original
}
func (b *BlindBox) UpdateCarbonReduction(carbonreduction int) {
	b.CarbonReduction = carbonreduction
}
func (b *BlindBox) UpdateImage(image string) {
	b.Image = image
}
func (b *Business) AddBusiness(name string, avatar string, address string, openinghours string, description string, image string, tele string) {
	b.Name = name
	b.Avatar = avatar
	b.Address = address
	b.Openinghours = openinghours
	b.Description = description
	b.Image = image
	b.Tele = tele
}
func (b *Business) UpdateBusinessScore(score float64) {
	b.Score = score
}
func (b *Business) UpdateBusinessAvatar(avatar string) {
	b.Avatar = avatar
}
func (b *Business) UpdateBusinessAddress(address string) {
	b.Address = address
}
func (b *Business) UpdateBusinessOpeninghours(openinghours string) {
	b.Openinghours = openinghours
}
func (b *Business) UpdateBusinessDescription(description string) {
	b.Description = description
}
func (b *Business) UpdateBusinessImage(image string) {
	b.Image = image
}
func (b *Business) UpdateBusinessTele(tele string) {
	b.Tele = tele
}
func (b *Business) UpdateBusinessName(name string) {
	b.Name = name
}
func (bb *BusinessAndBox) AddOneEdge(bussinessid, boxid int) {
	bb.BusinessID = bussinessid
	bb.BoxID = boxid
}
