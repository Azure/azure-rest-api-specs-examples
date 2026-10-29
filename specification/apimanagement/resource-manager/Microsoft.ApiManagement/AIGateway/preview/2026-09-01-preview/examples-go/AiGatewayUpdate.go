package armaigateway_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aigateway/armaigateway"
)

// Generated from example definition: 2026-09-01-preview/AiGatewayUpdate.json
func ExampleResourcesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armaigateway.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewResourcesClient().BeginUpdate(ctx, "example-rg", "example-gateway", "*", armaigateway.UpdateParameters{
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
		},
		SKU: &armaigateway.SKUUpdate{
			Name: to.Ptr("AIGateway"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armaigateway.ResourcesClientUpdateResponse{
	// 	Resource: armaigateway.Resource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.ApiManagement/aigateways/example-gateway"),
	// 		Name: to.Ptr("example-gateway"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/aigateways"),
	// 		Location: to.Ptr("westus2"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 		},
	// 		SKU: &armaigateway.SKU{
	// 			Name: to.Ptr("AIGateway"),
	// 		},
	// 		Identity: &armaigateway.ManagedServiceIdentity{
	// 			Type: to.Ptr(armaigateway.ManagedServiceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 		},
	// 		ETag: to.Ptr("\"AAAAAAAAAAI=\""),
	// 		Properties: &armaigateway.Properties{
	// 			ProvisioningState: to.Ptr(armaigateway.ProvisioningStateSucceeded),
	// 			Frontend: &armaigateway.Frontend{
	// 				DefaultHostname: to.Ptr("example-gateway.westus2.aigateway.azure-api.net"),
	// 			},
	// 		},
	// 	},
	// }
}
