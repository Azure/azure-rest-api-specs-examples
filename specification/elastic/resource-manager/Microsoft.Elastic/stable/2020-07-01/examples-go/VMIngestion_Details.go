package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/stable/2020-07-01/examples/VMIngestion_Details.json
func ExampleVMIngestionClient_Details() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armelastic.NewVMIngestionClient("<subscription-id>", cred, nil)
	_, err = client.Details(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
