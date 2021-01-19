## test-stone

1. You can test the solution with two approaches:

    a) go test -v *(it runs the unit test that is on the main_test.go file)*

    b) go run ./main.go 


2. For *b* option, you will need to follow the tips below: 

```go
// Everything that you are gonna see below are on the main function, you can just uncomment and run it

// Error Handling
 if err != nil {
 	fmt.Printf("err: %v\n", err.Error())
 	return
}

// Mock Data
// As the name says, one is a mock for normal data, and the other one is the periodic sequence case
itemsPeriodic, emailsPeriodic := mockDataToInfiniteTitheCase()
items, emails := mockData()

// if you want to use a different input, I have some tips for you
type item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
type email struct {
	EmailName string `json:"email"`
}

emails := []email{
    {"testing@something.com"},
    {"testingTwo@something.com"},
}

// Price field (third one) is using cents too help in the manipulation -> 100 are equivalent to R$ 1,00
items := []item{
    {"laranja", 1, 50},
    {"mexirica", 1, 50},
}

repository := itemRepository{}
service := itemService{repository}

// In the end but not least, you will need to call the function that does all the calculation
itemsInfinite, emailsInfinite := mockDataToInfiniteTitheCase()
items, emails := mockData()

repository := itemRepository{}
service := itemService{repository}

emptyList, err := service.GetAllItems([]item{}, []email{}) // -> Empty List
if err != nil {
    fmt.Printf("Empty List - err: %v\n", err.Error())
}
itemsPeriodic, err := service.GetAllItemsPeriodicSequence(itemsInfinite, emailsInfinite) //-> Infinite sequence
if err != nil {
    fmt.Printf("itemsPeriodic - err: %v\n", err.Error())
}
itemNotPeriodic, err := service.GetAllItems(items, emails) // -> Not periodic sequence
if err != nil {
    fmt.Printf("itemNotPeriodic - err: %v\n", err.Error())
}
fmt.Printf("Empty list: %v\nPeriodic sequence: %v\nNot Periodic sequence: %v\n", emptyList, itemsPeriodic, itemNotPeriodic)
```

## License
[MIT](https://choosealicense.com/licenses/mit/)

## Made by ♥ Alexandre Vardai 👋 
[Find me here](https://www.linkedin.com/in/alexandre-vardai-b8255b15b/)