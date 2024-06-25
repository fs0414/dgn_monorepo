package product

package (
	"fmt"
	"github.com/google/uuid"
)

type Product struct {
	id string
	ownerID string
	name string
	description string
	price int64
	stock int	
} 

func newProduct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	if !ulid.IsValid(ownerID) {
		return nil, fmt.Errorf("ownerIDの値が不正です")
	}
	if price < 1 {
		fmt.Println("価格が不正です")		
	}
	return &Product{
		id: id,
		ownerID: ownerID,
		name: name,
		description: description,
		price: price,
		stock: stock,
	}, nil
}

func Reconatruct(
	id string,
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
		stock
	)
}

func NewProduc(
	ownerID string	
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(
	    uuid.New(),
		name,
		description,
		price,
		stock
	)
}