package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/UpdateOrderItem.json
func ExampleManagementClient_BeginUpdateOrderItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("YourSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateOrderItem(ctx,
		"TestOrderItemName3",
		"YourResourceGroupName",
		armedgeorder.OrderItemUpdateParameter{
			Properties: &armedgeorder.OrderItemUpdateProperties{
				Preferences: &armedgeorder.Preferences{
					TransportPreferences: &armedgeorder.TransportPreferences{
						PreferredShipmentType: to.Ptr(armedgeorder.TransportShipmentTypesCustomerManaged),
					},
				},
			},
		},
		&armedgeorder.ManagementClientBeginUpdateOrderItemOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
