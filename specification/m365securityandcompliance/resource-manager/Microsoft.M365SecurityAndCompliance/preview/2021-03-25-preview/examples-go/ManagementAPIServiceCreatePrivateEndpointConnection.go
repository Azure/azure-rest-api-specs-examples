package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ManagementAPIServiceCreatePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsAdtAPIClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armm365securityandcompliance.NewPrivateEndpointConnectionsAdtAPIClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rgname",
		"service1",
		"myConnection",
		armm365securityandcompliance.PrivateEndpointConnection{
			Properties: &armm365securityandcompliance.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armm365securityandcompliance.PrivateLinkServiceConnectionState{
					Description: to.Ptr("Auto-Approved"),
					Status:      to.Ptr(armm365securityandcompliance.PrivateEndpointServiceConnectionStatusApproved),
				},
			},
		},
		nil)
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
