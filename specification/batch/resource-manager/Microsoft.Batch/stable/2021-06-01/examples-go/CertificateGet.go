package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/CertificateGet.json
func ExampleCertificateClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewCertificateClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<certificate-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CertificateClientGetResult)
}
