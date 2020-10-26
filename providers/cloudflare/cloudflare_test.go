package cloudflare

import (
	"net"
	"testing"
)

func TestCloudflare (t *testing.T) {
	cloudflareProvisioner := CloudflareProvisioner{
		ApiKey: "15091c568183e6c729ed5e2473f646b9967a8",
		Email: "ivanklee86@gmail.com",
		Domain: "aoach.tech",
		Names : []string{"*", ""},
	}

	cloudflareProvisioner.update(net.ParseIP("173.54.247.108"))
}