package mocks

import (
	"github.com/ParminCloud/go-proxmox/tests/mocks/config"
	"github.com/ParminCloud/go-proxmox/tests/mocks/pve6x"
	"github.com/ParminCloud/go-proxmox/tests/mocks/pve7x"
	"github.com/ParminCloud/go-proxmox/tests/mocks/pve8x"
	"github.com/h2non/gock"
)

func On(c config.Config) {
	ProxmoxVE7x(c) // default pve7
}

func Off() {
	gock.Off()
}

func ProxmoxVE8x(c config.Config) {
	config.C = c
	pve8x.Load()
}

func ProxmoxVE7x(c config.Config) {
	config.C = c
	pve7x.Load()
}

func ProxmoxVE6x(c config.Config) {
	config.C = c
	pve6x.Load()
}
