package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	arrEmail := strings.Split(email, "@")
	if len(arrEmail) != 2 {
		return ""
	}

	domain := arrEmail[1]
	arrParts := strings.Split(domain, ".")
	if len(arrParts) < 2 {
		return ""
	}

	domainName := arrParts[0]
	tld := strings.Join(arrParts[1:], ".")

	return fmt.Sprintf("Domain: %s dan TLD: %s", domainName, tld)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.com"))
}
