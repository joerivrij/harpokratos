package main

import (
	"github.com/ianschenck/envflag"
	"github.com/kpango/glg"
	harpokratos "harpokratos/pkg/impl"
	"net/http"
	"os"
	"strconv"
)

var customTag = "harpokratos"
var customErrTag = "CRIT"

func init() {
	errlog := glg.FileWriter("/tmp/error.log", 0666)
	defer errlog.Close()

	glg.Get().
		SetMode(glg.BOTH).
		AddLevelWriter(glg.ERR, errlog)
}

func main() {
	vaultService := envflag.String("VAULT_SERVICE","http://127.0.0.1", "location of the vault service")
	vaultUsername := envflag.String("VAULT_USERNAME","admin", "vault username")
	vaultPassword := envflag.String("VAULT_PASSWORD","password", "vault password")
	port := envflag.String("PORT",":5000", "port")

	envflag.Parse()

	glg.Info("welcome to harpokratos")
	glg.Info("starting up.....")
	glg.Debug("starting up and getting env variables")
	glg.Debugf("%s : %s", "VAULT_PASSWORD", *vaultPassword)
	glg.Debugf("%s : %s", "VAULT_USERNAME", *vaultUsername)
	glg.Debugf("%s : %s", "VAULT_SERVICE", *vaultService)

	glg.Debug("connecting to vault to establish health")
	vaultHealth := harpokratos.VaultHealth()

	if !vaultHealth {
		glg.Errorf("%s : %s", "vault healthy", strconv.FormatBool(vaultHealth))
		glg.Info("waiting for vault to become healthy")

		os.Exit(1)
	} else {
		glg.Infof("%s : %s", "vault healthy", strconv.FormatBool(vaultHealth))
	}

	srv := harpokratos.InitRoutes()

	glg.Infof("%s : %s", "running on port", *port)
	err := http.ListenAndServe(*port, srv)
	if err != nil {
		panic(err)
	}
}
