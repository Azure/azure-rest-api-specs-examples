package armedgeorderpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/ManageInventoryMetadata.json
func ExampleAPISClient_BeginManageInventoryMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorderpartner.NewAPISClient("b783ea86-c85c-4175-b76d-3992656af50d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginManageInventoryMetadata(ctx,
		"AzureStackEdge",
		"westus",
		"SerialNumber1",
		armedgeorderpartner.ManageInventoryMetadataRequest{
			ConfigurationOnDevice: &armedgeorderpartner.ConfigurationOnDevice{
				ConfigurationIdentifier: to.Ptr("EdgeP_High"),
			},
			InventoryMetadata: to.Ptr("InventoryMetadata"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
