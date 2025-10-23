package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: 2025-02-01-preview/DPSGetCertificate.json
func ExampleDpsCertificateClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDpsCertificateClient().Get(ctx, "cert", "myResourceGroup", "myFirstProvisioningService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceprovisioningservices.DpsCertificateClientGetResponse{
	// 	CertificateResponse: &armdeviceprovisioningservices.CertificateResponse{
	// 		Name: to.Ptr("cert"),
	// 		Type: to.Ptr("Microsoft.Devices/ProvisioningServices/Certificates"),
	// 		Etag: to.Ptr("AAAAAAExpNs="),
	// 		ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/andbuc-hub/certificates/cert"),
	// 		Properties: &armdeviceprovisioningservices.CertificateProperties{
	// 			Certificate: []byte("MA=="),
	// 			Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 			Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Sat, 31 Dec 2039 23:59:59 GMT"); return t}()),
	// 			IsVerified: to.Ptr(false),
	// 			Subject: to.Ptr("CN=testdevice1"),
	// 			Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 			Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 		},
	// 	},
	// }
}
