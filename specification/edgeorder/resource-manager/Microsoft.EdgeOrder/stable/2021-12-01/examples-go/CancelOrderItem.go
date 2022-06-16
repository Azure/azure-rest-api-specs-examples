package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CancelOrderItem.json
func ExampleManagementClient_CancelOrderItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("YourSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CancelOrderItem(ctx,
		"TestOrderItemName3",
		"YourResourceGroupName",
		armedgeorder.CancellationReason{
			Reason: to.Ptr("Order cancelled"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
