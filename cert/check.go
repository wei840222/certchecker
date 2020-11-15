package cert

//憑證相關
import (
	_ "github.com/wei840222/certchecker/conf"
	"github.com/wei840222/certchecker/db"

	"crypto/tls"
	"fmt"
	"time"
)

func checkHost(d *db.Domain) {
	conn, err := tls.Dial("tcp", d.Host, nil)
	if err != nil {
		db.UpdateDomain(d.ID, &db.Domain{
			Error: fmt.Sprintf("%s: %s %s", d.Name, d.Host, err),
		})
		return
	}
	defer conn.Close()

	var errorString string
	now := time.Now()
	for chainID, chain := range conn.ConnectionState().VerifiedChains {
		for certID, cert := range chain {
			if chainID == 0 && certID == 0 {
				db.UpdateDomain(d.ID, &db.Domain{
					Since: &cert.NotBefore,
					End:   &cert.NotAfter,
				})
			}
			// Check the expiration.
			if now.AddDate(0, 0, 14).After(cert.NotAfter) {
				expiresIn := int64(cert.NotAfter.Sub(now).Hours())
				if expiresIn <= 48 {
					errorString = fmt.Sprintf("%s: %s \nhttps://%s expires in %d hours", d.Name, d.Host, d.Host, expiresIn)
					break
				}
				errorString = fmt.Sprintf("%s: %s \nhttps://%s expires in roughly %d days", d.Name, d.Host, d.Host, expiresIn/24)
				break
			}
		}
		if errorString != "" {
			db.UpdateDomain(d.ID, &db.Domain{
				Error: errorString,
			})
			return
		}
	}
	if errorString == "" {
		db.DeleteDomainError(d.ID)
	}
}

func StartCertCheck() {
	for range time.NewTicker(5 * time.Second).C {
		domains, _ := db.ListDomain()
		for _, d := range domains {
			go checkHost(d)
		}
	}
}
