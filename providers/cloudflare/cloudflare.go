package cloudflare

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"log"
	"net"
)

type CloudflareProvisioner struct {
	ApiKey string   // Cloudflare API key.
	Email string    // Cloudflare email.
	Domain string   // Domain name to update.
	Names []string  // Names of A and AAAA records to update.  Note: This should be the subdomain only i.e. * and not *.test.com
}

func (cp CloudflareProvisioner) update(ip net.IP) {
	var (
		FullDnsNames []string  // Full DNS record name (imputed from Names & Domain).
		DnsType string         // DNS type
	)

	// Set up variables.
	for _, s := range cp.Names {
		if len(s) > 0 {
			FullDnsNames = append(FullDnsNames, fmt.Sprintf("%s.%s", s, cp.Domain))
		} else {
			FullDnsNames = append(FullDnsNames, cp.Domain)
		}
	}

	if ip.To4() != nil {
		DnsType = "A"
	} else {
		DnsType = "AAAA"
	}

	// Initialize Cloudflare API.
	api, err := cloudflare.New(cp.ApiKey, cp.Email)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the zone ID.
	ZoneId, err := api.ZoneIDByName(cp.Domain) // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		log.Fatal(err)
	}

	// Update DNS records.
	for _, DnsName := range FullDnsNames {
		records, err := api.DNSRecords(ZoneId, cloudflare.DNSRecord{Type: DnsType, Name: DnsName})
		if err != nil {
			log.Fatal(err)
		}

		DnsRecord := records[0]
		err = api.UpdateDNSRecord(ZoneId, DnsRecord.ID, cloudflare.DNSRecord{Content: ip.String()})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Updated %s (type %s) to IP %s", DnsName, DnsType, ip.String())
	}
}
