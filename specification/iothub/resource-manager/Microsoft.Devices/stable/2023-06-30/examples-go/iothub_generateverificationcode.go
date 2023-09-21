package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_generateverificationcode.json
func ExampleCertificatesClient_GenerateVerificationCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().GenerateVerificationCode(ctx, "myResourceGroup", "testHub", "cert", "AAAAAAAADGk=", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateWithNonceDescription = armiothub.CertificateWithNonceDescription{
	// 	Name: to.Ptr("cert"),
	// 	Properties: &armiothub.CertificatePropertiesWithNonce{
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Sat, 31 Dec 2039 23:59:59 GMT"); return t}()),
	// 		IsVerified: to.Ptr(false),
	// 		Subject: to.Ptr("CN=andbucdevice1"),
	// 		Thumbprint: to.Ptr("##############################"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:26:56 GMT"); return t}()),
	// 		VerificationCode: to.Ptr("##################################"),
	// 	},
	// }
}
