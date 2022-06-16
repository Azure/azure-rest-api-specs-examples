package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/ApplicationPackageActivate.json
func ExampleApplicationPackageClient_Activate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewApplicationPackageClient("<subscription-id>", cred, nil)
	res, err := client.Activate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<application-name>",
		"<version-name>",
		armbatch.ActivateApplicationPackageParameters{
			Format: to.StringPtr("<format>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ApplicationPackageClientActivateResult)
}
