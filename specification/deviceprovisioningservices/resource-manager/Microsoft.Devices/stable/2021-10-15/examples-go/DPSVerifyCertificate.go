package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSVerifyCertificate.json
func ExampleDpsCertificateClient_VerifyCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdeviceprovisioningservices.NewDpsCertificateClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.VerifyCertificate(ctx,
		"<certificate-name>",
		"<if-match>",
		"<resource-group-name>",
		"<provisioning-service-name>",
		armdeviceprovisioningservices.VerificationCodeRequest{
			Certificate: to.Ptr("<certificate>"),
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
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
