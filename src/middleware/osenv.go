package middleware

import (
	"utils"

	iris "gopkg.in/kataras/iris.v6"
)

type osEnvMw struct {
	LdapType string
	LdapAddr string
	LdapPort string
	LdapDC   string

	VMImagesDir string
}

func (s *osEnvMw) Serve(ctx *iris.Context) {
	ctx.Set("LdapType", s.LdapType)
	ctx.Set("LdapAddr", s.LdapAddr)
	ctx.Set("LdapPort", s.LdapPort)
	ctx.Set("LdapDC", s.LdapDC)
	ctx.Set("VMImagesDir", s.VMImagesDir)
	ctx.Next()
}

//NewOSEnv *
func NewOSEnv() iris.HandlerFunc {

	osEnv := &osEnvMw{}

	ldapAddr := utils.GetENV("LDAP_ADDR")
	if len(ldapAddr) == 0 {
		ldapAddr = "10.0.99.100"
	}

	osEnv.LdapAddr = ldapAddr

	ldapPort := utils.GetENV("LDAP_PORT")
	if len(ldapPort) == 0 {
		ldapPort = "636"
	}
	osEnv.LdapPort = ldapPort

	ldapDC := utils.GetENV("LDAP_DC")
	if len(ldapDC) == 0 {
		ldapDC = "dc=ronglian,dc=com"
	}

	osEnv.LdapDC = ldapDC

	ldapType := utils.GetENV("LDAP_TYPE")
	if len(ldapType) == 0 {
		ldapType = "ad"
	}

	osEnv.LdapType = ldapType

	vmImagesDir := utils.GetENV("VM_IMAGES_DIR")
	if len(vmImagesDir) == 0 {
		vmImagesDir = "/opt/images/vm"
	}

	osEnv.VMImagesDir = vmImagesDir

	return osEnv.Serve
}
