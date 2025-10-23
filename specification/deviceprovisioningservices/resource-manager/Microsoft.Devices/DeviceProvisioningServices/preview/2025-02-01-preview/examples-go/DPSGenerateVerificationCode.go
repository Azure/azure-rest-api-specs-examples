package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: 2025-02-01-preview/DPSGenerateVerificationCode.json
func ExampleDpsCertificateClient_GenerateVerificationCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDpsCertificateClient().GenerateVerificationCode(ctx, "cert", "AAAAAAAADGk=", "myResourceGroup", "myFirstProvisioningService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceprovisioningservices.DpsCertificateClientGenerateVerificationCodeResponse{
	// 	VerificationCodeResponse: &armdeviceprovisioningservices.VerificationCodeResponse{
	// 		Name: to.Ptr("cert"),
	// 		Properties: &armdeviceprovisioningservices.VerificationCodeResponseProperties{
	// 			Certificate: []byte("MA=="),
	// 			Created: to.Ptr("Thu, 12 Oct 2017 19:23:50 GMT"),
	// 			Expiry: to.Ptr("Sat, 31 Dec 2039 23:59:59 GMT"),
	// 			IsVerified: to.Ptr(false),
	// 			Subject: to.Ptr("CN=andbucdevice1"),
	// 			Thumbprint: to.Ptr("##############################"),
	// 			Updated: to.Ptr("Thu, 12 Oct 2017 19:26:56 GMT"),
	// 			VerificationCode: to.Ptr("##################################"),
	// 		},
	// 	},
	// }
}
