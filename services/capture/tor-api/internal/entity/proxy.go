package entity

import (
	"errors"
)

var (
	ErrEmptyID   = errors.New("empty id")
	ErrEmptyIP   = errors.New("empty ip")
	ErrEmptyPort = errors.New("empty port")
)

type Proxy struct {
	ID     string `json:"id"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Active bool   `json:"is_active"`
}

type ProxyEntityProps struct {
	ID     string
	IP     string
	Port   int
}

func NewProxy(props ProxyEntityProps) (*Proxy, error) {
	proxy := &Proxy{
		ID:   props.ID,
		IP:   props.IP,
		Port: props.Port,
	}
	proxy.Enable()
	if err := proxy.isValid(); err != nil {
		return nil, err
	}
	return proxy, nil
}

func (p *Proxy) Enable() {
	p.Active = true
}

func (p *Proxy) Disable() {
	p.Active = false
}

func (p *Proxy) GetID() string {
	return p.ID
}

func (p *Proxy) GetIP() string {
	return p.IP
}

func (p *Proxy) GetPort() int {
	return p.Port
}

func (p *Proxy) IsActive() bool {
	return p.Active
}

func (p *Proxy) isValid() error {
	if p.ID == "" {
		return ErrEmptyID
	}
	if p.IP == "" {
		return ErrEmptyIP
	}
	if p.Port == 0 {
		return ErrEmptyPort
	}
	return nil
}
