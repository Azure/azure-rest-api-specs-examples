package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/Accounts/Accounts_Update.json
func ExampleAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginUpdate(ctx, "test-rg", "contoso", armdeviceupdate.AccountUpdate{
		Tags: map[string]*string{
			"tagKey": to.Ptr("tagValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armdeviceupdate.Account{
	// 	Name: to.Ptr("contoso"),
	// 	Type: to.Ptr("Microsoft.DeviceUpdate/accounts"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DeviceUpdate/accounts/contoso"),
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"tagKey": to.Ptr("tagValue"),
	// 	},
	// 	Properties: &armdeviceupdate.AccountProperties{
	// 		HostName: to.Ptr("contoso.api.adu.microsoft.com"),
	// 		Locations: []*armdeviceupdate.Location{
	// 			{
	// 				Name: to.Ptr("westus2"),
	// 				Role: to.Ptr(armdeviceupdate.RolePrimary),
	// 			},
	// 			{
	// 				Name: to.Ptr("westcentralus"),
	// 				Role: to.Ptr(armdeviceupdate.RoleFailover),
	// 		}},
	// 		PrivateEndpointConnections: []*armdeviceupdate.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("peexample01"),
	// 				Type: to.Ptr("Microsoft.DeviceUpdate/accounts/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DeviceUpdate/accounts/contoso/privateEndpointConnections/peexample01"),
	// 				Properties: &armdeviceupdate.PrivateEndpointConnectionProperties{
	// 					GroupIDs: []*string{
	// 						to.Ptr("groupId")},
	// 						PrivateEndpoint: &armdeviceupdate.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/peexample01"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armdeviceupdate.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Auto-Approved"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armdeviceupdate.PrivateEndpointServiceConnectionStatusApproved),
	// 						},
	// 						ProvisioningState: to.Ptr(armdeviceupdate.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			ProvisioningState: to.Ptr(armdeviceupdate.ProvisioningStateSucceeded),
	// 			SKU: to.Ptr(armdeviceupdate.SKUStandard),
	// 		},
	// 	}
}
