package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0d41e635294dce73dfa99b07f3da4b68a9c9e29c/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/TrafficManagerUserMetricsKeys-GET.json
func ExampleUserMetricsKeysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtrafficmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserMetricsKeysClient().Get(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UserMetricsModel = armtrafficmanager.UserMetricsModel{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Network/trafficManagerUserMetricsKeys"),
	// 	ID: to.Ptr("/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default"),
	// 	Properties: &armtrafficmanager.UserMetricsProperties{
	// 		Key: to.Ptr("9ea056eb38f145a0891b5d5dc15e9aa2"),
	// 	},
	// }
}
