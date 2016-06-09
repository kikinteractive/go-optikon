# go-optikon [![GoDoc][godoc image]][godoc] [![Build Status][travis image]][travis] [![Coverage Status][codecov image]][codecov] [![sourcegraph views][sourcegraph image]][sourcegraph]

Manipulating deep Go structures using simple relative path selectors. Useful for REST-ful services,
where you want to select/update/delete internal structures having a relative path to them. 

![Optikon Image][image]

## Installation

```bash
# get the package
$ go get github.com/rounds/go-optikon

# download package dependencies
$ go get -t ./...

# build
$ go build

# run tests
$ go test -v -cover 
```

### Dependencies
* [github.com/stretchr/testify/assert](https://github.com/stretchr/testify/assert)

## Usage

Add the following line in your Go file:
```go
import "github.com/rounds/go-optikon"
```

### Functions
```go
func Select(dataIn interface{}, path []string) (interface{}, error)
func UpdateJSON(dataIn interface{}, path []string, dataJSON json.RawMessage, opType OpType) error
```

### Examples

###### Our data model
```go
type OrderedProduct struct {
	SKU 		string	`json:"sku"`
	Quantity 	int		`json:"quantity"`
	SubTotal	float64 `json:"subtotal"`
}
type Order struct {
	ID 			string					`json:"id"`
	Date		string					`json:"date"`
	Products	[]OrderedProduct		`json:"products"`
	Total		float64					`json:"total"`
	ExtraProps	map[string]interface{}	`json:"extraprops"`
}
type Address struct {
	Street 	string	`json:"street"`
	City 	string	`json:"city"`
	Country string	`json:"country"`
}
type Customer struct {
	ID			string  `json:"id"`
	Name 		string	`json:"name"`
	AvatarURL 	string	`json:"avatar"`
	Address 	Address	`json:"address"`
	Orders		[]Order	`json:"orders"`
}

// Sample Customer instance.
customer := Customer {
	ID:				"a12345",
	Name: 			"John Doe", 
	AvatarURL: 		"http://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50",
	Address: Address {
		Street: 	"Pine St. 18",
		City: 		"Sherwood",
		Country:	"Narnia",
	},
	Orders: []Order{
		Order {
			ID: 	"o12345",
			Date: 	"2015-06-01T23:20:22Z",
			Products: []OrderedProduct {
				OrderedProduct{
					SKU: 		"1234-1234-1234",
					Quantity: 	1,
					SubTotal: 	12.40,
				},
				OrderedProduct{
					SKU: 		"3456-3456-3456",
					Quantity: 	3,
					SubTotal: 	54.60,
				},
			},
			Total: 67.0,
			ExtraProps: map[string]interface{}{
				"shipping company":	"EMS",
				"packaging cost":	12.20,
				"permits":	map[string]interface{}{
					"export permit": "XO-1221",
					"shipping permit": 4324234432,
				},
			},
		},
	},
}
```

###### Select
```go
// Getting the Address object:
// REST: GET /customers/a12345/address
addr, err = optikon.Select(customer, []string{"address"})

// Getting the first customer order
// REST: GET /customers/a12345/orders/0
order, err := optikon.Select(customer, []string{"orders", "0"})

// Getting the first product from the first customer order
// REST: GET /customers/a12345/orders/0/products/0
prod, err := optikon.Select(customer, []string{"orders", "0", "products", "0"})

// Getting export permit code for the first order
// REST: GET /customers/a12345/orders/0/extraprops/permits/export%20permit
code, err := optikon.Select(customer, []string{"orders", "0", "extraprops", "permits", "export permit"})
```

###### Update
```go
// Update address
// REST: PATCH /customers/a12345/address
addrJSON := `{
	"street": "another street",
	"city": "another city",
	"country": "another country"
}`
err = optikon.UpdateJSON(customer, []string{"address"}, []byte(addrJSON), optikon.UpdateOp)

// Update the first ordered product in the first customer order
// REST: PATCH /customers/a12345/orders/0/products/0
prodJSON := `{
	"sku": "4321-4321-4321",
	"quantity": 1,
	"subtotal": 12.40
}`
err = optikon.UpdateJSON(customer, []string{"orders", "0", "products", "0"}, []byte(prodJSON), optikon.UpdateOp)

// Update the export permit code in the first order
// REST: PATCH /customers/a12345/orders/0/extraprops/permits/export%20permit
err = optikon.UpdateJSON(customer, []string{"orders", "0", "extraprops", "permits", "export permit"}, []byte("XO-2222"), optikon.UpdateOp)
```

###### Create
```go
// Create a new order
// REST: POST /customers/a12345/orders
orderJSON := `{
	"id": "o12347",
	"date": "2015-06-01T10:10:10Z",
	"products": [
		{
			"sku": "2345-2345-2345",
			"quantity": 2,
			"subtotal": 100.0
		}
	],
	"total": 100.0,
	"extraprops": {
		"shipping company":	"USPost",
		"packaging cost":	0
	}
}`
err = optikon.UpdateJSON(customer, []string{"orders"}, []byte(orderJSON), optikon.CreateOp)

// Add a property to ExtraProps of the first order
// REST: POST /customers/a12345/orders/0/extra%20props
propJSON := `{
	"new property": "value"
}`
err = optikon.UpdateJSON(customer, []string{"orders", "0", "extraprops"}, []byte(orderJSON), optikon.CreateOp)
```

###### Delete
```go
// Delete the first product from the first customer's order
// REST: DELETE /customers/a12345/orders/0/products/0
err = optikon.UpdateJSON(customer, []string{"orders", "0", "products", "0"}, nil, optikon.DeleteOp)

// Delete the shipping permit property from permits in ExtraProps of the first order
// REST: DELETE /customers/a12345/orders/0/extra%20props/permits/shipping%20permit
err = optikon.UpdateJSON(customer, []string{"orders", "0", "extra props", "permits", "shipping permit"}, nil, optikon.DeleteOp)
```

## Contribute

Please check the [issues][issues] page which might have some TODOs.
Feel free to file new bugs and ask for improvements. We welcome pull requests!


[godoc]: https://godoc.org/github.com/rounds/go-optikon
[godoc image]: https://godoc.org/github.com/rounds/go-optikon?status.svg
[travis image]: https://travis-ci.org/rounds/go-optikon.svg
[travis]: https://travis-ci.org/rounds/go-optikon
[codecov image]: https://codecov.io/gh/rounds/go-optikon/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/rounds/go-optikon
[rounds]: http://rounds.com/
[image]: optikon-i45x.jpeg
[blog post]: http://rounds.com/blog/collecting-user-data-and-usage/
[gvm]: https://github.com/moovweb/gvm
[issues]: https://github.com/rounds/go-optikon/issues
[sourcegraph image]: https://sourcegraph.com/api/repos/github.com/rounds/go-optikon/.counters/views.svg
[sourcegraph]: https://sourcegraph.com/github.com/rounds/go-optikon
