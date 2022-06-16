package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2020-03-01/examples/DPSVerifyCertificate.json
func ExampleDpsCertificateClient_VerifyCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeviceprovisioningservices.NewDpsCertificateClient("<subscription-id>", cred, nil)
	res, err := client.VerifyCertificate(ctx,
		"<certificate-name>",
		"<if-match>",
		"<resource-group-name>",
		"<provisioning-service-name>",
		armdeviceprovisioningservices.VerificationCodeRequest{
			Certificate: to.StringPtr("<certificate>"),
		},
		&armdeviceprovisioningservices.DpsCertificateClientVerifyCertificateOptions{CertificateName1: nil,
			CertificateIsVerified:    nil,
			CertificatePurpose:       nil,
			CertificateCreated:       nil,
			CertificateLastUpdated:   nil,
			CertificateHasPrivateKey: nil,
			CertificateNonce:         nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DpsCertificateClientVerifyCertificateResult)
}
