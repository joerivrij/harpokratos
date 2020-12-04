package config

import (
	"flag"
	"github.com/ianschenck/envflag"
	"github.com/kpango/glg"
)

type HarpokratosConfig struct {
	VaultToken string
	VaultService string
	VaultUser string
	VaultPassword string
}

func Get() *HarpokratosConfig {
	vaultService := envflag.String("VAULT_SERVICE","http://127.0.0.1:8200", "location of the vault service")
	vaultUsername := envflag.String("VAULT_USERNAME","harpokratos", "vault username")
	vaultPassword := envflag.String("VAULT_PASSWORD","harpokratos", "vault password")
	vaultRootToken := envflag.String("VAULT_ROOT_TOKEN", "root", "root token for dev purposes")

	envflag.Parse()

	glg.Debugf("%s : %s", "VAULT_PASSWORD", *vaultPassword)
	glg.Debugf("%s : %s", "VAULT_USERNAME", *vaultUsername)
	glg.Debugf("%s : %s", "VAULT_SERVICE", *vaultService)
	glg.Debugf("%s : %s", "VAULT_ROOT_TOKEN", *vaultRootToken)

	config := &HarpokratosConfig{
		VaultToken:   *vaultRootToken,
		VaultService: *vaultService,
		VaultUser: *vaultUsername,
		VaultPassword: *vaultPassword,
	}

	flag.Parse()

	return config
}