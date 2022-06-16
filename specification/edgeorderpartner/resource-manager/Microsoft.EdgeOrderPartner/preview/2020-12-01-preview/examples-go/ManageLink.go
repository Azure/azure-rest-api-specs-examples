package armedgeorderpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/ManageLink.json
func ExampleAPISClient_ManageLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorderpartner.NewAPISClient("b783ea86-c85c-4175-b76d-3992656af50d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.ManageLink(ctx,
		"AzureStackEdge",
		"westus",
		"SerialNumber1",
		armedgeorderpartner.ManageLinkRequest{
			ManagementResourceArmID: to.Ptr("/subscriptions/c783ea86-c85c-4175-b76d-3992656af50d/resourceGroups/EdgeTestRG/providers/Microsoft.DataBoxEdge/DataBoxEdgeDevices/TestEdgeDeviceName1"),
			Operation:               to.Ptr(armedgeorderpartner.ManageLinkOperationLink),
			TenantID:                to.Ptr("a783ea86-c85c-4175-b76d-3992656af50d"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
