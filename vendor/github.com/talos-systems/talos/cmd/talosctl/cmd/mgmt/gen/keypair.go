// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gen

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/spf13/cobra"
	"github.com/talos-systems/crypto/x509"

	"github.com/talos-systems/talos/pkg/cli"
)

var genKeypairCmdFlags struct {
	ip           string
	organization string
}

// genKeypairCmd represents the `gen keypair` command.
var genKeypairCmd = &cobra.Command{
	Use:   "keypair",
	Short: "Generates an X.509 Ed25519 key pair",
	Long:  ``,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := []x509.Option{}
		if genKeypairCmdFlags.ip != "" {
			parsed := net.ParseIP(genKeypairCmdFlags.ip)
			if parsed == nil {
				return fmt.Errorf("invalid IP: %s", genKeypairCmdFlags.ip)
			}
			ips := []net.IP{parsed}
			opts = append(opts, x509.IPAddresses(ips))
		}
		if genKeypairCmdFlags.organization != "" {
			opts = append(opts, x509.Organization(genKeypairCmdFlags.organization))
		}
		ca, err := x509.NewSelfSignedCertificateAuthority(opts...)
		if err != nil {
			return fmt.Errorf("error generating CA: %s", err)
		}
		if err := ioutil.WriteFile(genKeypairCmdFlags.organization+".crt", ca.CrtPEM, 0o600); err != nil {
			return fmt.Errorf("error writing certificate: %s", err)
		}
		if err := ioutil.WriteFile(genKeypairCmdFlags.organization+".key", ca.KeyPEM, 0o600); err != nil {
			return fmt.Errorf("error writing key: %s", err)
		}

		return nil
	},
}

func init() {
	genKeypairCmd.Flags().StringVar(&genKeypairCmdFlags.ip, "ip", "", "generate the certificate for this IP address")
	genKeypairCmd.Flags().StringVar(&genKeypairCmdFlags.organization, "organization", "", "X.509 distinguished name for the Organization")
	cli.Should(cobra.MarkFlagRequired(genKeypairCmd.Flags(), "organization"))

	Cmd.AddCommand(genKeypairCmd)
}
