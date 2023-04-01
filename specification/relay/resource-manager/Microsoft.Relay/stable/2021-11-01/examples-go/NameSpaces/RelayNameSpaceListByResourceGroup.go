package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/NameSpaces/RelayNameSpaceListByResourceGroup.json
func ExampleNamespacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListByResourceGroupPager("resourcegroup", nil)
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
		// 			Name: to.Ptr("example-RelayNamespace-3054"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-3054"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-3054"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-3054.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("oaisdjfoiasdjfoiajsdfoijasd"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/oaisdjfoiasdjfoiajsdfoijasd"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:oaisdjfoiasdjfoiajsdfoijasd"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://oaisdjfoiasdjfoiajsdfoijasd.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("abc123"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/abc123"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:abc123"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://abc123.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("example-RelayNamespace-5849"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-5849"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-5849"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-5849.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("example-RelayNamespace-4984"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-4984"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-4984"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-4984.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("example-RelayNamespace-5606"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-5606"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-5606"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-5606.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("example-RelayNamespace-7703"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-7703"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-7703"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-7703.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("zzzzzzzzzzzzzzzzzzzzzz-00001"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/zzzzzzzzzzzzzzzzzzzzzz-00001"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:zzzzzzzzzzzzzzzzzzzzzz-00001"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://zzzzzzzzzzzzzzzzzzzzzz-00001.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("example-RelayNamespace-3919"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-3919"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-3919"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-3919.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("aiosdjfaoidjasdoijasdfoijasdfofijsdf"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/aiosdjfaoidjasdoijasdfoijasdfofijsdf"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:aiosdjfaoidjasdoijasdfoijasdfofijsdf"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://aiosdjfaoidjasdoijasdfoijasdfofijsdf.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("example-RelayNamespace-3413"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-3413"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-3413"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-3413.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("example-RelayNamespace-8695"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-8695"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-8695"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-8695.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("example-RelayNamespace-4801"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-4801"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-4801"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-4801.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("example-RelayNamespace-2812"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-2812"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armrelay.NamespaceProperties{
		// 				MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relaynamespace-2812"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://example-RelayNamespace-2812.servicebus.windows-int.net:443/"),
		// 			},
		// 			SKU: &armrelay.SKU{
		// 				Name: to.Ptr(armrelay.SKUNameStandard),
		// 				Tier: to.Ptr(armrelay.SKUTierStandard),
		// 			},
		// 	}},
		// }
	}
}
