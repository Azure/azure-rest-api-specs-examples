package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay/v2"
)

// Generated from example definition: 2024-01-01/HybridConnection/RelayHybridConnectionCreate.json
func ExampleHybridConnectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHybridConnectionsClient().CreateOrUpdate(ctx, "resourcegroup", "example-RelayNamespace-01", "example-Relay-Hybrid-01", armrelay.HybridConnection{
		Properties: &armrelay.HybridConnectionProperties{
			RequiresClientAuthorization: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrelay.HybridConnectionsClientCreateOrUpdateResponse{
	// 	HybridConnection: &armrelay.HybridConnection{
	// 		Name: to.Ptr("example-Relay-Hybrid-01"),
	// 		Type: to.Ptr("Microsoft.Relay/Namespaces/HybridConnections"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-01/HybridConnections/example-Relay-Hybrid-01"),
	// 		Properties: &armrelay.HybridConnectionProperties{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-23T20:34:49.4131724Z"); return t}()),
	// 			RequiresClientAuthorization: to.Ptr(true),
	// 			UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-23T20:34:49.4131724Z"); return t}()),
	// 		},
	// 	},
	// }
}
