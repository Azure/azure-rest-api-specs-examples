package armedgeorderpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/SearchInventories.json
func ExampleAPISClient_NewSearchInventoriesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorderpartner.NewAPISClient("b783ea86-c85c-4175-b76d-3992656af50d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewSearchInventoriesPager(armedgeorderpartner.SearchInventoriesRequest{
		FamilyIdentifier: to.Ptr("AzureStackEdge"),
		SerialNumber:     to.Ptr("SerialNumber1"),
	},
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
