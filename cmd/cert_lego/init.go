package cert_lego

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/http01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type LegoUser struct {
	Email        string
	Registration *registration.Resource
	Key          crypto.PrivateKey
}

func (u *LegoUser) GetEmail() string {
	return u.Email
}

func (u *LegoUser) GetRegistration() *registration.Resource {
	return u.Registration
}

func (u *LegoUser) GetPrivateKey() crypto.PrivateKey {
	return u.Key
}

const (
	HTTP01 = "http01"
	DNS01  = "dns01"
)

var (
	myUser LegoUser
)

func init() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic("failed to generate private key")
	}
	myUser = LegoUser{
		Email: "fezko1054@gmail.com",
		Key:   privateKey,
	}
}

type ConfigPayload struct {
	ServerName      []string `json:"server_name"`
	ChallengeMethod string   `json:"challenge_method"`
	DNSCredentialID int      `json:"dns_credential_id"`
}

var HttpChallengeCmd = &cobra.Command{
	Use:   "lego_http",
	Short: "set up challenge response",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("run lego_http")
		config := lego.NewConfig(&myUser)
		client, err := lego.NewClient(config)
		if err != nil {
			return err
		}

		err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("",
			"9980",
		),
		)
		if err != nil {
			fmt.Println("fail to setup http01 challenge server")
			return err
		}

		fmt.Println("end lego_http")
		return nil
	},
}

var RegisterCmd = &cobra.Command{
	Use:   "lego_register",
	Short: "leogo registration",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := lego.NewConfig(&myUser)
		config.CADirURL = "https://acme-staging-v02.api.letsencrypt.org/directory"
		config.Certificate.KeyType = certcrypto.RSA2048

		client, err := lego.NewClient(config)
		if err != nil {
			fmt.Println("fail to create new client of lego:" + err.Error())
			return err
		}

		reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
		if err != nil {
			fmt.Println("register failed:", err.Error())
			return err
		}

		err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("",
			"9980",
		),
		)
		if err != nil {
			fmt.Println("fail to setup http01 challenge server")
			return err
		}

		fmt.Println("end lego_http")

		fmt.Println("registration:", reg)

		request := certificate.ObtainRequest{
			Domains: []string{"ssl3.cloudfy669.xyz"},
			Bundle:  true,
		}

		certificates, err := client.Certificate.Obtain(request)
		if err != nil {
			fmt.Println("failed to obtain:", err.Error())
			return err
		}

		fmt.Println("get certificate success")
		saveDir := "/Users/tom/golang/example/temp"
		err = os.WriteFile(filepath.Join(saveDir, "fullchain.cer"), certificates.Certificate, os.ModePerm)
		if err != nil {
			fmt.Println("failed to write fullchain.cer")
			return err
		}

		fmt.Println("finished")
		return nil
	},
}

var ObtainCmd = cobra.Command{
	Use:   "leogo_obtain",
	Short: "获取",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
