package usecase

import (
	"services/capture/tor-api/internal/entity"
)

type CreateProxyUseCase struct {

}

func NewCreateProxyUseCase() *CreateProxyUseCase {
	return &CreateProxyUseCase{}
}

func (u *CreateProxyUseCase) Execute() {
	entityProps := entity.ProxyEntityProps{
		ID: "1",
		IP: "",
		Port: 8000,
	}
	proxyEntity, err := entity.NewProxy(entityProps)
	if err != nil {
		panic(err)
	}
}