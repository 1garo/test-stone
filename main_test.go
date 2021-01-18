package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
	"github.com/stretchr/testify/mock"
)

type ItemRepositoryMock struct {
	mock.Mock
}

func (r *ItemRepositoryMock) GetAllItems(items []Item, emails []Email) (map[string]int, error) {
	_ = r.Called()
	ret, err := calculateFinal(items, emails)
	return ret, err
}

func (r *ItemRepositoryMock) GetAllItemsPeriodicSequence(items []Item, emails []Email) (map[string]int, error) {
	_ = r.Called()
	ret, err := calculateFinalInfiniteTitheCase(items, emails)
	return ret, err
}

func TestService_GetItemsCorrect(t *testing.T) {
	repository := ItemRepositoryMock{}
	repository.On("GetAllItems").Return([]int{}, nil)

	service := ItemService{&repository}
	items, emails := mockData()
	respItems, _ := service.GetAllItems(items, emails)
	for email := range respItems {
		assert.Equal(t, respItems[email], 50)
		fmt.Printf("-> Non periodic sequence items: %s | %+v\n", email, respItems[email])
	}
}

func TestService_GetItemsPeriodicSequence(t *testing.T) {
	repository := ItemRepositoryMock{}
	repository.On("GetAllItemsPeriodicSequence").Return([]int{}, nil)

	service := ItemService{&repository}
	items, emails := mockDataToInfiniteTitheCase()
	respItemsInfinite, _ := service.GetAllItemsPeriodicSequence(items, emails)
	for email := range respItemsInfinite {
		assert.NotEqual(t, respItemsInfinite[email], 50)
		fmt.Printf("-> Periodic sequence items: %s | %+v\n", email, respItemsInfinite[email])
	}
}

func TestService_GetItemsEmpty(t *testing.T) {
	repository := ItemRepositoryMock{}
	repository.On("GetAllItemsPeriodicSequence").Return([]int{}, nil)

	service := ItemService{&repository}
	_, err := service.GetAllItemsPeriodicSequence([]Item{}, []Email{})
	if err != nil {
		assert.Equal(t, err.Error(), "cannot work on empty item or email list")
	}
	fmt.Println("-> Empty items")
}
