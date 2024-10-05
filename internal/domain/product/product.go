package product

import (
	"fmt"
	"time"
  "math/rand"
  "unicode/utf8"

	"github.com/oklog/ulid/v2"
)

type Product struct {
	id          ulid.ULID
	ownerID     string
	name        string
	description string
	price       int64
	stock       int
}

const (
  nameLengthMin = 1 
  nameLengthMax = 100
  descriptionLengthMin = 1 
  descriptionLengthMax = 1000
)
func newProduct(
	id ulid.ULID,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
  if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax { 
    return nil, fmt.Errorf("商品名の値が不正です 。")
  }
  if price < 1 {
		fmt.Println("価格が不正です")
	}
	return &Product{
		id:          id,
		ownerID:     ownerID,
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
	}, nil
}

func Reconatruct(
	id ulid.ULID,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(
		id,
		ownerID,
		name,
		description,
		price,
		stock,
	)
}

func NewProduc(
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
  t := time.Now()
  entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
  id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return newProduct(
    id,
    ownerID,
		name,
		description,
		price,
		stock,
	)
}
