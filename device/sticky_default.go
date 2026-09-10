//go:build !linux

package device

import (
	"github.com/snehesht/wireguard-go/conn"
	"github.com/snehesht/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
