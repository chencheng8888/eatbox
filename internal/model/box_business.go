package model

import "gorm.io/gorm"

type BlindBox struct {
	gorm.Model
	BlindType       int    `json:"type"`
	Sale            int    `json:"sale" gorm:"default:0"`
	Number          int    `json:"number" gorm:"default:0"`
	Original        string `json:"original"`
	Discount        int    `json:"discount" gorm:"default:0"`
	CarbonReduction int    `json:"carbon_reduction"`
	Image           string `json:"image"`
	Status          int    `json:"status" gorm:"default:0"`
}
type Business struct {
	//ID           int     `gorm:"primaryKey" json:"id"`
	gorm.Model
	Name         string  `json:"name"`
	Avatar       string  `json:"avatar"`
	Address      string  `json:"address"`
	Openinghours string  `json:"openinghours"`
	Description  string  `json:"description"`
	Image        string  `json:"image"`
	Score        float64 `json:"score"`
	Tele         string  `json:"tele"`
}
type BusinessAndBox struct {
	gorm.Model
	BusinessID int `json:"business_id"`
	BoxID      int `json:"box_id"`
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
