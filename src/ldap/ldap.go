package ldap

import (
	"crypto/tls"
	"fmt"
	"strings"

	goldap "gopkg.in/ldap.v2"
)

type LDAPType struct {
	Name           string
	AuthFilter     string
	AuthAttributes []string
}

var ldapTypes = []*LDAPType{
	ActiveDirectory,
	OpenLDAP,
}

func (l *LDAPType) String() string {
	return l.Name
}

func (l *LDAPType) ParseAuthFilter(values ...string) {
	m := make(map[string]string)
	for i := 0; i < len(values); i += 2 {
		m[values[i]] = values[i+1]
	}
	l.AuthFilterReplace(m)
}

func (l *LDAPType) AuthFilterReplace(match map[string]string) {
	for k, v := range match {
		l.AuthFilter = strings.Replace(l.AuthFilter, "{"+k+"}", v, -1)
	}
}

func ByLDAPType(typeName string) *LDAPType {
	for _, ldapType := range ldapTypes {
		if ldapType.Name == typeName {
			return ldapType
		}
	}
	return nil
}

var ActiveDirectory = &LDAPType{
	Name:           "ad",
	AuthFilter:     "(|(cn={user})(sAMAccountName={user})(userPrincipalName={user}@{domain})(userPrincipalName={name})(mail={name}))",
	AuthAttributes: []string{"sAMAccountName", "cn", "name", "mail", "department"},
}

var OpenLDAP = &LDAPType{
	Name:           "openldap",
	AuthFilter:     "(|(cn={name})(uid={name})(mail={name}))",
	AuthAttributes: []string{"uid", "cn", "name", "mail", "department"},
}

type LDAPClient struct {
	Conn     *goldap.Conn
	BaseDN   string
	Host     string
	Port     int
	UseTLS   bool
	UserName string
	Password string
	LdapType string
}

func NewLDAPClient(host string, port int, useTLS bool, baseDN, userName, password, ldapType string) *LDAPClient {
	return &LDAPClient{
		Host:     strings.TrimSpace(host),
		Port:     port,
		UseTLS:   useTLS,
		BaseDN:   strings.TrimSpace(baseDN),
		UserName: strings.TrimSpace(userName),
		Password: strings.TrimSpace(password),
		LdapType: strings.TrimSpace(ldapType),
	}
}

func (l *LDAPClient) Connect() (err error) {
	var lc *goldap.Conn
	ldapAddr := fmt.Sprintf("%s:%d", l.Host, l.Port)
	if l.UseTLS {
		config := &tls.Config{InsecureSkipVerify: true}
		if lc, err = goldap.DialTLS("tcp", ldapAddr, config); nil != err {
			return
		}
	} else {
		if lc, err = goldap.Dial("tcp", ldapAddr); nil != err {
			return
		}
	}

	if "" != l.UserName && "" != l.Password {
		if err = lc.Bind(l.UserName, l.Password); nil != err {
			return
		}
	} else {
		err = fmt.Errorf("%s", "Parameters cannot be empty.")
		return
	}
	l.Conn = lc
	return
}

func (l *LDAPClient) Close() {
	if nil != l.Conn {
		l.Conn.Close()
		l.Conn = nil
	}
}

func (l *LDAPClient) NewSearchRequest(searchRequestType string) (*goldap.SearchRequest, error) {
	lt := ByLDAPType(l.LdapType)
	switch searchRequestType {
	case "auth":
		sep := "@"
		if strings.Contains(l.UserName, sep) {
			userInfos := strings.Split(l.UserName, sep)
			if len(userInfos) == 2 {
				name := strings.TrimSpace(userInfos[0])
				domain := strings.TrimSpace(userInfos[1])
				lt.ParseAuthFilter("user", name, "name", l.UserName, "domain", domain)
			}
		} else {
			lt.ParseAuthFilter("user", l.UserName)
		}

		return goldap.NewSearchRequest(
			l.BaseDN,
			goldap.ScopeWholeSubtree, goldap.NeverDerefAliases, 0, 0, false,
			lt.AuthFilter,
			lt.AuthAttributes,
			nil,
		), nil
	default:
		return nil, fmt.Errorf("%s", "...")
	}
}

func (l *LDAPClient) Authenticate() (entries map[string]string, err error) {
	if err = l.Connect(); nil != err {
		return
	}

	searchRequest, err := l.NewSearchRequest("auth")
	if nil != err {
		return
	}

	sr, err := l.Conn.Search(searchRequest)
	if err != nil {
		return
	}

	if len(sr.Entries) > 0 {
		entries = make(map[string]string)
		for _, attr := range searchRequest.Attributes {
			entries[attr] = sr.Entries[0].GetAttributeValue(attr)
		}
		return
	}
	return
}
