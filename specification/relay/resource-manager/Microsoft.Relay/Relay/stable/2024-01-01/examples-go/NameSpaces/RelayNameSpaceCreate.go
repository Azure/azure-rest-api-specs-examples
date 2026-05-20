package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay/v2"
)

// Generated from example definition: 2024-01-01/NameSpaces/RelayNameSpaceCreate.json
func ExampleNamespacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginCreateOrUpdate(ctx, "resourcegroup", "example-RelayNamespace-5849", armrelay.Namespace{
		Location: to.Ptr("South Central US"),
		SKU: &armrelay.SKU{
			Name: to.Ptr(armrelay.SKUNameStandard),
			Tier: to.Ptr(armrelay.SKUTierStandard),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res = armrelay.NamespacesClientCreateOrUpdateResponse{
	// 	Namespace: &armrelay.Namespace{
	// 		Name: to.Ptr("example-RelayNamespace-5849"),
	// 		Type: to.Ptr("Microsoft.Relay/Namespaces"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-5849"),
	// 		Location: to.Ptr("South Central US"),
	// 		Properties: &armrelay.NamespaceProperties{
	// 			MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-5849"),
	// 			ProvisioningState: to.Ptr("Created"),
	// 			ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-5849.servicebus.windows-int.net:443/"),
	// 		},
	// 		SKU: &armrelay.SKU{
	// 			Name: to.Ptr(armrelay.SKUNameStandard),
	// 			Tier: to.Ptr(armrelay.SKUTierStandard),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
