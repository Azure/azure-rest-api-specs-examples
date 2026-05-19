package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay/v2"
)

// Generated from example definition: 2024-01-01/NameSpaces/RelayNameSpaceGet.json
func ExampleNamespacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().Get(ctx, "RG-eg", "example-RelayRelayNamespace-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrelay.NamespacesClientGetResponse{
	// 	Namespace: &armrelay.Namespace{
	// 		Name: to.Ptr("example-RelayRelayNamespace-01"),
	// 		Type: to.Ptr("Microsoft.Relay/Namespaces"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/RG-eg/providers/Microsoft.Relay/namespaces/example-RelayRelayNamespace-01"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armrelay.NamespaceProperties{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-23T20:38:12.46Z"); return t}()),
	// 			MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-RelayRelayNamespace-01"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ServiceBusEndpoint: to.Ptr("https://example-RelayRelayNamespace-01.servicebus.windows.net:443/"),
	// 			UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-23T20:38:34.533Z"); return t}()),
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
