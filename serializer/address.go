package serializer

import "gin-mall/model"

type AddressSerializer struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func BuildAddress(address *model.Address) *AddressSerializer {
	return &AddressSerializer{
		Address: address.Address,
		Name:    address.Name,
		Phone:   address.Phone,
	}
}

func BuildAddresses(items []*model.Address) (addresses []*AddressSerializer) {
	for _, item := range items {
		addresses = append(addresses, BuildAddress(item))
	}
	return addresses
}
