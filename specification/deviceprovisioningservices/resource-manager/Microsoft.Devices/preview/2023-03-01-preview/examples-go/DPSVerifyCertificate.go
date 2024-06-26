package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSVerifyCertificate.json
func ExampleDpsCertificateClient_VerifyCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDpsCertificateClient().VerifyCertificate(ctx, "cert", "AAAAAAAADGk=", "myResourceGroup", "myFirstProvisioningService", armdeviceprovisioningservices.VerificationCodeRequest{
		Certificate: to.Ptr("#####################################"),
	}, &armdeviceprovisioningservices.DpsCertificateClientVerifyCertificateOptions{CertificateName1: nil,
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateResponse = armdeviceprovisioningservices.CertificateResponse{
	// 	Name: to.Ptr("cert"),
	// 	Type: to.Ptr("Microsoft.Devices/ProvisioningServices/Certificates"),
	// 	Etag: to.Ptr("AAAAAAExpTQ="),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/ProvisioningServices/myFirstProvisioningService/certificates/cert"),
	// 	Properties: &armdeviceprovisioningservices.CertificateProperties{
	// 		Certificate: []byte("MA=="),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:23:50.000Z"); return t}()),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2039-12-31T23:59:59.000Z"); return t}()),
	// 		IsVerified: to.Ptr(true),
	// 		Subject: to.Ptr("CN=andbucdevice1"),
	// 		Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:26:56.000Z"); return t}()),
	// 	},
	// }
}
