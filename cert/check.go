package cert

//憑證相關
import (
	"crypto/tls"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wei840222/certchecker/bot"
	"github.com/wei840222/certchecker/db"
)

func checkHost(d *db.Domain) error {
	conn, err := tls.Dial("tcp", d.Host, nil)
	if err != nil {
		return fmt.Errorf("%s: %s https://%s", d.Name, d.Host, err)
	}
	defer conn.Close()

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
			if now.AddDate(0, 3, 0).After(cert.NotAfter) {
				expiresIn := int64(cert.NotAfter.Sub(now).Hours())
				if expiresIn <= 48 {
					return fmt.Errorf("%s: %s \nhttps://%s expires in %d hours", d.Name, d.Host, d.Host, expiresIn)
				}
				return fmt.Errorf("%s: %s \nhttps://%s expires in roughly %d days", d.Name, d.Host, d.Host, expiresIn/24)
			}
		}
	}
	return nil
}

func StartCertCheck() {
	for range time.NewTicker(5 * time.Minute).C {
		domains, _ := db.ListDomain()
		for _, d := range domains {
			if err := checkHost(d); err != nil {
				msg := tgbotapi.NewMessage(909503895, err.Error())
				bot.Bot.Send(msg)
			}
		}
	}
}
func StartCertDateCheck() {
	for range time.NewTicker(5 * time.Second).C {
		domains, _ := db.ListDomain()
		for _, m := range domains {
			if err := checkHost(m); err != nil {
				return
			}
		}
	}
}
