package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub/v2"
)

// Generated from example definition: 2026-05-01-preview/iothub_generateverificationcode.json
func ExampleCertificatesClient_GenerateVerificationCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
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
	// res = armiothub.CertificatesClientGenerateVerificationCodeResponse{
	// 	CertificateWithNonceDescription: armiothub.CertificateWithNonceDescription{
	// 		Name: to.Ptr("cert"),
	// 		Properties: &armiothub.CertificatePropertiesWithNonce{
	// 			Created: to.Ptr(time.Date(2017, time.October, 12, 19, 23, 50, 0, time.UTC)),
	// 			Expiry: to.Ptr(time.Date(2039, time.December, 31, 23, 59, 59, 0, time.UTC)),
	// 			IsVerified: to.Ptr(false),
	// 			Subject: to.Ptr("CN=andbucdevice1"),
	// 			Thumbprint: to.Ptr("##############################"),
	// 			Updated: to.Ptr(time.Date(2017, time.October, 12, 19, 26, 56, 0, time.UTC)),
	// 			VerificationCode: to.Ptr("##################################"),
	// 		},
	// 	},
	// }
}
