package main

import (
	"errors"
)

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type Email struct {
	EmailName string `json:"email"`
}

type ItemRepositoryInterface interface {
	GetAllItems([]Item, []Email) (map[string]int, error)
	GetAllItemsPeriodicSequence([]Item, []Email) (map[string]int, error)
}

type ItemService struct {
	ItemRepositoryInterface
}

type ItemRepository struct{}

func (r ItemRepository) GetAllItems(items []Item, emails []Email) (map[string]int, error) {
	finalResult, err := calculateFinal(items, emails)
	if err != nil {
		return nil, err
	}
	return finalResult, nil
}

func (r ItemRepository) GetAllItemsPeriodicSequence(items []Item, emails []Email) (map[string]int, error) {
	finalResult, err := calculateFinalInfiniteTitheCase(items, emails)
	if err != nil {
		return nil, err
	}
	return finalResult, nil
}

func mockData() ([]Item, []Email) {
	email := []Email{
		{"testing@hotmail.com"},
		{"testingTwo@gmail.com"},
	}
	// using cents too help in the manipulation
	users := []Item{
		{"uva", 1, 50},
		{"pessego", 1, 50},
	}
	return users, email
}

func mockDataToInfiniteTitheCase() ([]Item, []Email) {
	email := []Email{
		{"testing@hotmail.com"},
		{"testingTwo@gmail.com"},
		{"testingThird@gmail.com"},
	}
	// using cents too help in the manipulation
	users := []Item{
		{"uva", 1, 150},
		{"pessego", 1, 50},
	}
	return users, email
}

func calculateItemsSum(items []Item) (int, error) {
	var acc int
	for _, s := range items {
		acc += s.Price * s.Quantity
	}
	return acc, nil
}

func dividePriceByEmails(price int, emails []Email) ([]int, error) {
	var resp []int
	var value int
	var ok bool
	emailsLen := len(emails)
	isPeriodicSequence := price % emailsLen
	if isPeriodicSequence != 0 {
		ok = true
	}
	for range emails {
		value = price / emailsLen
		resp = append(resp, value)
	}
	if ok {
		resp[emailsLen-1] = (price / emailsLen) + 1
	}
	return resp, nil
}

func finalCalculation(value []int, emails []Email) (map[string]int, error) {
	s := make(map[string]int)
	for i, emails := range emails {
		s[emails.EmailName] = value[i]
	}
	return s, nil
}

func calculateFinal(items []Item, emails []Email) (map[string]int, error) {
	// TODO: implement err return on all function calls below to be catched on iferr
	if (len(items) | len(emails)) == 0 {
		return nil, errors.New("cannot work on empty item or email list")
	}
	itemsSummed, err := calculateItemsSum(items)
	if err != nil {
		return nil, err
	}
	finalValue, err := dividePriceByEmails(itemsSummed, emails)
	if err != nil {
		return nil, err
	}
	finalResult, err := finalCalculation(finalValue, emails)
	if err != nil {
		return nil, err
	}
	return finalResult, nil
}

func calculateFinalInfiniteTitheCase(items []Item, emails []Email) (map[string]int, error) {
	if (len(items) | len(emails)) == 0 {
		return nil, errors.New("cannot work on empty item or email list")
	}
	itemsSummed, err := calculateItemsSum(items)
	if err != nil {
		return nil, err
	}
	finalValue, err := dividePriceByEmails(itemsSummed, emails)
	if err != nil {
		return nil, err
	}
	finalResult, err := finalCalculation(finalValue, emails)
	if err != nil {
		return nil, err
	}
	return finalResult, nil
}

func main() {
	// TODO: Create todo on how to test the program by here
	// itemsInfinite, emailsInfinite := mockDataToInfiniteTitheCase()
	// items, emails := mockData()
	// repository := ItemRepository{}
	// service := ItemService{repository}
	// itemsNonInfinite, err := service.GetAllItems([]Item{}, []Email{})
	// itemsInfinite, _ := service.GetAllItemsPeriodicSequence(itemsInfinite, emailsInfinite)
	// fmt.Println(itemsNonInfinite)
}
