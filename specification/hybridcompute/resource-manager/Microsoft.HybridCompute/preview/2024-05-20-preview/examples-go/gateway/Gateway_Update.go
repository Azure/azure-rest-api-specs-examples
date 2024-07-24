package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/gateway/Gateway_Update.json
func ExampleGatewaysClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewaysClient().Update(ctx, "myResourceGroup", "{gatewayName}", armhybridcompute.GatewayUpdate{
		Properties: &armhybridcompute.GatewayUpdateProperties{
			AllowedFeatures: []*string{
				to.Ptr("*")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gateway = armhybridcompute.Gateway{
	// 	Name: to.Ptr("{gatewayName}"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/gateways"),
	// 	ID: to.Ptr("/subscriptions/ffd506c8-3415-42d3-9612-fdb423fb17df/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/gateways/{gatewayName}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armhybridcompute.GatewayProperties{
	// 		AllowedFeatures: []*string{
	// 			to.Ptr("*")},
	// 			GatewayEndpoint: to.Ptr("https://uniqueValue.contoso.com"),
	// 			GatewayID: to.Ptr("<generated Guid>"),
	// 			GatewayType: to.Ptr(armhybridcompute.GatewayTypePublic),
	// 			ProvisioningState: to.Ptr(armhybridcompute.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
