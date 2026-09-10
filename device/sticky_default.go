//go:build !linux

package device

import (
	"github.com/wirepeer/wireguard-go/conn"
	"github.com/wirepeer/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
