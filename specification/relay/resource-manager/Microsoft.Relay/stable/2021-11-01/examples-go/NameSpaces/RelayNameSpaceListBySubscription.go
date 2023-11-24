package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/NameSpaces/RelayNameSpaceListBySubscription.json
func ExampleNamespacesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.NamespaceListResult = armrelay.NamespaceListResult{
		// 	Value: []*armrelay.Namespace{
		// 		{
		// 			Name: to.Ptr("example-RelayRelayNamespace-01"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/RG1-eg/providers/Microsoft.Relay/namespaces/example-RelayRelayNamespace-01"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-23T20:34:49.413Z"); return t}()),
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-RelayRelayNamespace-01"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayRelayNamespace-01.servicebus.windows.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-23T20:34:59.413Z"); return t}()),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("example-RelayRelayNamespace-02"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/RG1-eg/providers/Microsoft.Relay/namespaces/example-RelayRelayNamespace-02"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-23T20:34:39.413Z"); return t}()),
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-RelayRelayNamespace-02"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayRelayNamespace-02.servicebus.windows.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-23T20:34:49.413Z"); return t}()),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 	}},
		// }
	}
}
