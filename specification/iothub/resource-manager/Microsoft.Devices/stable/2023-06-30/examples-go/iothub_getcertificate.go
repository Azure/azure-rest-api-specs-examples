package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_getcertificate.json
func ExampleCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().Get(ctx, "myResourceGroup", "testhub", "cert", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateDescription = armiothub.CertificateDescription{
	// 	Name: to.Ptr("cert"),
	// 	Type: to.Ptr("Microsoft.Devices/IotHubs/Certificates"),
	// 	Etag: to.Ptr("AAAAAAExpNs="),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/andbuc-hub/certificates/cert"),
	// 	Properties: &armiothub.CertificateProperties{
	// 		Certificate: to.Ptr("############################################"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:23:50.000Z"); return t}()),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2039-12-31T23:59:59.000Z"); return t}()),
	// 		IsVerified: to.Ptr(false),
	// 		Subject: to.Ptr("CN=testdevice1"),
	// 		Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2017-10-12T19:23:50.000Z"); return t}()),
	// 	},
	// }
}
