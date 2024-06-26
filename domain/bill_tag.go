// Code generated by gtag. DO NOT EDIT.
// See: https://github.com/wolfogre/gtag

//go:generate go run github.com/wolfogre/gtag/cmd/gtag -types Bill -tags bson .
package domain

import (
	"reflect"
	"strings"
)

var (
	valueOfBill = Bill{}
	typeOfBill  = reflect.TypeOf(valueOfBill)

	_                = valueOfBill.ID
	fieldOfBillID, _ = typeOfBill.FieldByName("ID")
	tagOfBillID      = fieldOfBillID.Tag

	_                    = valueOfBill.BookId
	fieldOfBillBookId, _ = typeOfBill.FieldByName("BookId")
	tagOfBillBookId      = fieldOfBillBookId.Tag

	_                   = valueOfBill.Money
	fieldOfBillMoney, _ = typeOfBill.FieldByName("Money")
	tagOfBillMoney      = fieldOfBillMoney.Tag

	_                  = valueOfBill.Type
	fieldOfBillType, _ = typeOfBill.FieldByName("Type")
	tagOfBillType      = fieldOfBillType.Tag

	_                      = valueOfBill.Category
	fieldOfBillCategory, _ = typeOfBill.FieldByName("Category")
	tagOfBillCategory      = fieldOfBillCategory.Tag

	_                     = valueOfBill.CrtUser
	fieldOfBillCrtUser, _ = typeOfBill.FieldByName("CrtUser")
	tagOfBillCrtUser      = fieldOfBillCrtUser.Tag

	_                  = valueOfBill.Time
	fieldOfBillTime, _ = typeOfBill.FieldByName("Time")
	tagOfBillTime      = fieldOfBillTime.Tag

	_                     = valueOfBill.CrtTime
	fieldOfBillCrtTime, _ = typeOfBill.FieldByName("CrtTime")
	tagOfBillCrtTime      = fieldOfBillCrtTime.Tag

	_                     = valueOfBill.UpdTime
	fieldOfBillUpdTime, _ = typeOfBill.FieldByName("UpdTime")
	tagOfBillUpdTime      = fieldOfBillUpdTime.Tag

	_                    = valueOfBill.Remark
	fieldOfBillRemark, _ = typeOfBill.FieldByName("Remark")
	tagOfBillRemark      = fieldOfBillRemark.Tag

	_                    = valueOfBill.Images
	fieldOfBillImages, _ = typeOfBill.FieldByName("Images")
	tagOfBillImages      = fieldOfBillImages.Tag
)

// BillTags indicate tags of type Bill
type BillTags struct {
	ID       string // `bson:"_id,omitempty" json:"_id"`
	BookId   string // `bson:"book_id" json:"book_id"`
	Money    string // `bson:"money" json:"money"`
	Type     string // `bson:"type" json:"type"`
	Category string // `bson:"category" json:"category"`
	CrtUser  string // `bson:"crt_user" json:"crt_user"`
	Time     string // `bson:"time" json:"time"`
	CrtTime  string // `bson:"crt_time" json:"crt_time"`
	UpdTime  string // `bson:"upd_time" json:"upd_time"`
	Remark   string // `bson:"remark" json:"remark"`
	Images   string // `bson:"images" json:"images"`
}

// Values return all tags of Bill as slice
func (t *BillTags) Values() []string {
	return []string{
		t.ID,
		t.BookId,
		t.Money,
		t.Type,
		t.Category,
		t.CrtUser,
		t.Time,
		t.CrtTime,
		t.UpdTime,
		t.Remark,
		t.Images,
	}
}

// Tags return specified tags of Bill
func (*Bill) Tags(tag string, convert ...func(string) string) BillTags {
	conv := func(in string) string { return strings.TrimSpace(strings.Split(in, ",")[0]) }
	if len(convert) > 0 {
		conv = convert[0]
	}
	if conv == nil {
		conv = func(in string) string { return in }
	}
	return BillTags{
		ID:       conv(tagOfBillID.Get(tag)),
		BookId:   conv(tagOfBillBookId.Get(tag)),
		Money:    conv(tagOfBillMoney.Get(tag)),
		Type:     conv(tagOfBillType.Get(tag)),
		Category: conv(tagOfBillCategory.Get(tag)),
		CrtUser:  conv(tagOfBillCrtUser.Get(tag)),
		Time:     conv(tagOfBillTime.Get(tag)),
		CrtTime:  conv(tagOfBillCrtTime.Get(tag)),
		UpdTime:  conv(tagOfBillUpdTime.Get(tag)),
		Remark:   conv(tagOfBillRemark.Get(tag)),
		Images:   conv(tagOfBillImages.Get(tag)),
	}
}

// TagsBson is alias of Tags("bson")
func (*Bill) TagsBson() BillTags {
	var v *Bill
	return v.Tags("bson")
}
