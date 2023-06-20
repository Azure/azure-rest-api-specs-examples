package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSDeleteCertificate.json
func ExampleDpsCertificateClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDpsCertificateClient().Delete(ctx, "myResourceGroup", "AAAAAAAADGk=", "myFirstProvisioningService", "cert", &armdeviceprovisioningservices.DpsCertificateClientDeleteOptions{CertificateName1: nil,
		CertificateIsVerified:    nil,
		CertificatePurpose:       nil,
		CertificateCreated:       nil,
		CertificateLastUpdated:   nil,
		CertificateHasPrivateKey: nil,
		CertificateNonce:         nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
