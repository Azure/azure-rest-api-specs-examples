package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceCreatePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsSecClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsSecClient().BeginCreateOrUpdate(ctx, "rgname", "service1", "myConnection", armm365securityandcompliance.PrivateEndpointConnection{
		Properties: &armm365securityandcompliance.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armm365securityandcompliance.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Auto-Approved"),
				Status:      to.Ptr(armm365securityandcompliance.PrivateEndpointServiceConnectionStatusApproved),
			},
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
	// res.PrivateEndpointConnection = armm365securityandcompliance.PrivateEndpointConnection{
	// 	Name: to.Ptr("myConnection"),
	// 	Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/service1/privateEndpointConnections/myConnection"),
	// 	Properties: &armm365securityandcompliance.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armm365securityandcompliance.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/peexample01"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armm365securityandcompliance.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armm365securityandcompliance.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armm365securityandcompliance.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// 	SystemData: &armm365securityandcompliance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		CreatedBy: to.Ptr("sove"),
	// 		CreatedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sove"),
	// 		LastModifiedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 	},
	// }
}
