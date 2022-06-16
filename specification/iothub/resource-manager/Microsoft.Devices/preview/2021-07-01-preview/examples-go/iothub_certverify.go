package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2021-07-01-preview/examples/iothub_certverify.json
func ExampleCertificatesClient_Verify() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiothub.NewCertificatesClient("<subscription-id>", cred, nil)
	res, err := client.Verify(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<certificate-name>",
		"<if-match>",
		armiothub.CertificateVerificationDescription{
			Certificate: to.StringPtr("<certificate>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CertificatesClientVerifyResult)
}
