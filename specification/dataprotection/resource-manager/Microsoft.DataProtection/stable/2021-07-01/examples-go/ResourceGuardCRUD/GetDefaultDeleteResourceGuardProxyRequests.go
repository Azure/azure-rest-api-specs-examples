package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/ResourceGuardCRUD/GetDefaultDeleteResourceGuardProxyRequests.json
func ExampleResourceGuardsClient_GetDefaultDeleteResourceGuardProxyRequestsObject() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewResourceGuardsClient("<subscription-id>", cred, nil)
	res, err := client.GetDefaultDeleteResourceGuardProxyRequestsObject(ctx,
		"<resource-group-name>",
		"<resource-guards-name>",
		"<request-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DppBaseResource.ID: %s\n", *res.ID)
}
