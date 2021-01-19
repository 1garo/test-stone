package main

import (
	"errors"
	"fmt"
)

type item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type email struct {
	EmailName string `json:"email"`
}

type itemRepositoryInterface interface {
	GetAllItems([]item, []email) (map[string]int, error)
	GetAllItemsPeriodicSequence([]item, []email) (map[string]int, error)
}

type itemService struct {
	itemRepositoryInterface
}

type itemRepository struct{}

// GetAllItems -> Return a map with email as key and the price as value
func (r itemRepository) GetAllItems(items []item, emails []email) (map[string]int, error) {
	finalResult, err := calculateFinal(items, emails)
	if err != nil {
		return nil, err
	}
	return finalResult, nil
}

// GetAllItemsPeriodicSequence -> Return a map with email as key and the price as value that are a periodic sequence
func (r itemRepository) GetAllItemsPeriodicSequence(items []item, emails []email) (map[string]int, error) {
	finalResult, err := calculateFinalInfiniteTitheCase(items, emails)
	if err != nil {
		return nil, err
	}
	return finalResult, nil
}

func mockData() ([]item, []email) {
	email := []email{
		{"testing@hotmail.com"},
		{"testingTwo@gmail.com"},
	}
	// using cents too help in the manipulation
	item := []item{
		{"uva", 1, 50},
		{"pessego", 1, 50},
	}
	return item, email
}

func mockDataToInfiniteTitheCase() ([]item, []email) {
	email := []email{
		{"testing@hotmail.com"},
		{"testingTwo@gmail.com"},
		{"testingThird@gmail.com"},
		{"testingFouth@gmail.com"},
		{"testingFifth@gmail.com"},
		{"testingSixty@gmail.com"},
	}
	// using cents too help in the manipulation
	item := []item{
		{"uva", 1, 150},
		{"pessego", 1, 50},
	}
	return item, email
}

func calculateItemsSum(items []item) int {
	var acc int
	for _, s := range items {
		acc += s.Price * s.Quantity
	}
	return acc
}

func dividePriceByEmails(price int, emails []email) []int {
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
	return resp
}

func finalCalculation(value []int, emails []email) map[string]int {
	s := make(map[string]int, len(emails))
	for i, emails := range emails {
		s[emails.EmailName] = value[i]
	}
	return s
}

func calculateFinal(items []item, emails []email) (map[string]int, error) {
	if (len(items) | len(emails)) == 0 {
		return nil, errors.New("cannot work on empty item or email list")
	}

	itemsSummed := calculateItemsSum(items)
	finalValue := dividePriceByEmails(itemsSummed, emails)
	finalResult := finalCalculation(finalValue, emails)

	return finalResult, nil
}

func calculateFinalInfiniteTitheCase(items []item, emails []email) (map[string]int, error) {
	if (len(items) | len(emails)) == 0 {
		return nil, errors.New("cannot work on empty item or email list")
	}

	itemsSummed := calculateItemsSum(items)
	finalValue := dividePriceByEmails(itemsSummed, emails)
	finalResult := finalCalculation(finalValue, emails)

	return finalResult, nil
}

func main() {
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
}
