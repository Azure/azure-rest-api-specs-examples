package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/ResourceGuardCRUD/PatchResourceGuard.json
func ExampleResourceGuardsClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewResourceGuardsClient("<subscription-id>", cred, nil)
	res, err := client.Patch(ctx,
		"<resource-group-name>",
		"<resource-guards-name>",
		armdataprotection.PatchResourceRequestInput{
			Tags: map[string]*string{
				"newKey": to.StringPtr("newVal"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ResourceGuardResource.ID: %s\n", *res.ID)
}
