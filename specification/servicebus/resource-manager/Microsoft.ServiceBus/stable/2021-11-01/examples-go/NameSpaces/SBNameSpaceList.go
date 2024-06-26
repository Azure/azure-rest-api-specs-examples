package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/SBNameSpaceList.json
func ExampleNamespacesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.SBNamespaceListResult = armservicebus.SBNamespaceListResult{
		// 	Value: []*armservicebus.SBNamespace{
		// 		{
		// 			Name: to.Ptr("NS-91f08e47-2b04-4943-b0cd-a5fb02b88f20"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-91f08e47-2b04-4943-b0cd-a5fb02b88f20"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T02:40:17.270Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-91f08e47-2b04-4943-b0cd-a5fb02b88f20"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-91f08e47-2b04-4943-b0cd-a5fb02b88f20.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T07:15:30.780Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-41dc63f4-0b08-4029-b3ef-535a131bfa65"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-41dc63f4-0b08-4029-b3ef-535a131bfa65"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T03:50:38.980Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-41dc63f4-0b08-4029-b3ef-535a131bfa65"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-41dc63f4-0b08-4029-b3ef-535a131bfa65.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T10:42:58.003Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-df52cf51-e831-4bf2-bd92-e9885f68a996"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-df52cf51-e831-4bf2-bd92-e9885f68a996"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T01:17:54.997Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-df52cf51-e831-4bf2-bd92-e9885f68a996"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-df52cf51-e831-4bf2-bd92-e9885f68a996.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T06:44:39.737Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SBPremium"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/RapscallionResources/providers/Microsoft.ServiceBus/namespaces/SBPremium"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-10T22:01:00.420Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:sbpremium"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://SBPremium.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-10T22:01:00.420Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNamePremium),
		// 				Capacity: to.Ptr[int32](1),
		// 				Tier: to.Ptr(armservicebus.SKUTierPremium),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("rrama-ns2"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/sadfsadfsadf/providers/Microsoft.ServiceBus/namespaces/rrama-ns2"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T04:14:00.013Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:rrama-ns2"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://rrama-ns2.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-03T22:53:32.927Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-20e57600-29d0-4035-ac85-74f4c54dcda1"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-20e57600-29d0-4035-ac85-74f4c54dcda1"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T03:30:49.160Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-20e57600-29d0-4035-ac85-74f4c54dcda1"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-20e57600-29d0-4035-ac85-74f4c54dcda1.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T04:17:58.483Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-3e538a1a-58fb-4315-b2ce-76f5c944114c"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-3e538a1a-58fb-4315-b2ce-76f5c944114c"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T18:07:30.050Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-3e538a1a-58fb-4315-b2ce-76f5c944114c"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-3e538a1a-58fb-4315-b2ce-76f5c944114c.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T10:42:57.747Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("prem-ns123"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/prem-ns123"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-13T00:02:39.997Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:prem-ns123"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://prem-ns123.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-13T00:02:39.997Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNamePremium),
		// 				Capacity: to.Ptr[int32](1),
		// 				Tier: to.Ptr(armservicebus.SKUTierPremium),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-4e1bfdf1-0cff-4e86-ae80-cdcac4873039"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-4e1bfdf1-0cff-4e86-ae80-cdcac4873039"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T01:01:58.730Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-4e1bfdf1-0cff-4e86-ae80-cdcac4873039"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-4e1bfdf1-0cff-4e86-ae80-cdcac4873039.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T03:02:59.800Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-6b90b7f3-7aa0-48c9-bc30-b299dcb66c03"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-6b90b7f3-7aa0-48c9-bc30-b299dcb66c03"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T03:22:45.327Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-6b90b7f3-7aa0-48c9-bc30-b299dcb66c03"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-6b90b7f3-7aa0-48c9-bc30-b299dcb66c03.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T06:08:01.207Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-c05e9df3-7737-44ee-a321-15f6e0545b97"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-c05e9df3-7737-44ee-a321-15f6e0545b97"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T03:29:19.750Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-c05e9df3-7737-44ee-a321-15f6e0545b97"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-c05e9df3-7737-44ee-a321-15f6e0545b97.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T08:10:35.527Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-dcb4152c-231b-4c16-a683-07cc6b38fa46"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-dcb4152c-231b-4c16-a683-07cc6b38fa46"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T03:34:35.363Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-dcb4152c-231b-4c16-a683-07cc6b38fa46"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-dcb4152c-231b-4c16-a683-07cc6b38fa46.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T05:33:00.957Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-f501f5e6-1f24-439b-8982-9af665156d40"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-f501f5e6-1f24-439b-8982-9af665156d40"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T01:25:55.707Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-f501f5e6-1f24-439b-8982-9af665156d40"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-f501f5e6-1f24-439b-8982-9af665156d40.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T07:42:59.687Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-fe2ed660-2cd6-46f2-a9c3-7e11551a1f30"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-fe2ed660-2cd6-46f2-a9c3-7e11551a1f30"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T02:32:08.227Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-fe2ed660-2cd6-46f2-a9c3-7e11551a1f30"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-fe2ed660-2cd6-46f2-a9c3-7e11551a1f30.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T06:32:57.770Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-8a5e3b4e-4e97-4d85-9083-cd33536c9d71"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-8a5e3b4e-4e97-4d85-9083-cd33536c9d71"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T00:54:05.103Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-8a5e3b4e-4e97-4d85-9083-cd33536c9d71"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-8a5e3b4e-4e97-4d85-9083-cd33536c9d71.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T10:43:50.313Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-6520cc09-01ac-40a3-bc09-c5c431116e92"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-6520cc09-01ac-40a3-bc09-c5c431116e92"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T01:49:59.243Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-6520cc09-01ac-40a3-bc09-c5c431116e92"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-6520cc09-01ac-40a3-bc09-c5c431116e92.servicebus.windows-int.net:443"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T08:15:36.950Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-bfba6d5c-a425-42d9-85db-0f4da770e29a"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-bfba6d5c-a425-42d9-85db-0f4da770e29a"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T03:23:32.083Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-bfba6d5c-a425-42d9-85db-0f4da770e29a"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-bfba6d5c-a425-42d9-85db-0f4da770e29a.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T09:02:57.433Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SBPrem"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/RapscallionResources/providers/Microsoft.ServiceBus/namespaces/SBPrem"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-10T22:16:30.870Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:sbprem"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://SBPrem.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-10T22:16:30.870Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNamePremium),
		// 				Capacity: to.Ptr[int32](1),
		// 				Tier: to.Ptr(armservicebus.SKUTierPremium),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-43b136b4-8716-40b2-97c5-0d77cac0062c"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-43b136b4-8716-40b2-97c5-0d77cac0062c"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T03:14:50.577Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-43b136b4-8716-40b2-97c5-0d77cac0062c"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-43b136b4-8716-40b2-97c5-0d77cac0062c.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T09:23:01.067Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-7c0443de-5f88-450c-b574-83f60a097dd1"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-7c0443de-5f88-450c-b574-83f60a097dd1"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T04:07:15.397Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-7c0443de-5f88-450c-b574-83f60a097dd1"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-7c0443de-5f88-450c-b574-83f60a097dd1.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T04:03:03.097Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-62dd7753-a5f9-42fd-a354-ca38a4505d69"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-62dd7753-a5f9-42fd-a354-ca38a4505d69"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T01:33:50.450Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-62dd7753-a5f9-42fd-a354-ca38a4505d69"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-62dd7753-a5f9-42fd-a354-ca38a4505d69.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T05:35:33.053Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-ae18a18c-97ab-4089-965d-8acbf4794091"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-ae18a18c-97ab-4089-965d-8acbf4794091"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T02:43:36.517Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-ae18a18c-97ab-4089-965d-8acbf4794091"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-ae18a18c-97ab-4089-965d-8acbf4794091.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T12:40:30.587Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-8e3b56c1-0ee8-4e13-ae88-5cadf6e2ce11"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-8e3b56c1-0ee8-4e13-ae88-5cadf6e2ce11"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T00:46:03.773Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-8e3b56c1-0ee8-4e13-ae88-5cadf6e2ce11"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-8e3b56c1-0ee8-4e13-ae88-5cadf6e2ce11.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T04:43:54.560Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-7ffca4b4-4728-4fb0-b2d0-1e7c016e3a44"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-7ffca4b4-4728-4fb0-b2d0-1e7c016e3a44"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T03:59:12.100Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-7ffca4b4-4728-4fb0-b2d0-1e7c016e3a44"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-7ffca4b4-4728-4fb0-b2d0-1e7c016e3a44.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T06:33:52.230Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-d9337efd-9b27-454c-b2a5-dcfea56920d9"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-d9337efd-9b27-454c-b2a5-dcfea56920d9"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T03:45:09.270Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-d9337efd-9b27-454c-b2a5-dcfea56920d9"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-d9337efd-9b27-454c-b2a5-dcfea56920d9.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T06:20:31.863Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-ad5ae732-abea-4e62-9de0-c90de0ddec0a"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-ad5ae732-abea-4e62-9de0-c90de0ddec0a"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T02:34:36.447Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-ad5ae732-abea-4e62-9de0-c90de0ddec0a"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-ad5ae732-abea-4e62-9de0-c90de0ddec0a.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T06:15:31.607Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-d447fb03-c7da-40fe-b5eb-14f36888837b"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-d447fb03-c7da-40fe-b5eb-14f36888837b"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T00:53:46.697Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-d447fb03-c7da-40fe-b5eb-14f36888837b"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-d447fb03-c7da-40fe-b5eb-14f36888837b.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T11:09:41.260Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ReproSB"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/RapscallionResources/providers/Microsoft.ServiceBus/namespaces/ReproSB"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-27T19:29:34.523Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:reprosb"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://ReproSB.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-27T19:29:58.640Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-4c90097f-19a8-42e7-bb3c-4ac088994719"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-4c90097f-19a8-42e7-bb3c-4ac088994719"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T17:35:32.610Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-4c90097f-19a8-42e7-bb3c-4ac088994719"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-4c90097f-19a8-42e7-bb3c-4ac088994719.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T09:13:52.270Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("rrama-1-23-17"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/rrama-1-23-17"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-23T22:54:40.907Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:rrama-1-23-17"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://rrama-1-23-17.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-04T00:53:28.777Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-5191e541-8e4e-4229-9fdc-b89f6c3e7f12"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-5191e541-8e4e-4229-9fdc-b89f6c3e7f12"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T17:43:25.710Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-5191e541-8e4e-4229-9fdc-b89f6c3e7f12"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-5191e541-8e4e-4229-9fdc-b89f6c3e7f12.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T11:05:31.890Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-be903820-3533-46e8-90e4-72c132411848"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-be903820-3533-46e8-90e4-72c132411848"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T03:24:01.923Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-be903820-3533-46e8-90e4-72c132411848"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-be903820-3533-46e8-90e4-72c132411848.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T10:09:42.513Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("rrama-namespace1"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/rrama-namespace1"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T00:47:22.963Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:rrama-namespace1"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://rrama-namespace1.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T00:47:27.297Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-a3c38e9b-32a3-4c51-85d7-263150a8dda9"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-a3c38e9b-32a3-4c51-85d7-263150a8dda9"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T00:38:02.517Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-a3c38e9b-32a3-4c51-85d7-263150a8dda9"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-a3c38e9b-32a3-4c51-85d7-263150a8dda9.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T05:03:55.960Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-70d3fa25-6bbe-4a6b-a381-a52cf0d539e6"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-70d3fa25-6bbe-4a6b-a381-a52cf0d539e6"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-23T03:42:40.010Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-70d3fa25-6bbe-4a6b-a381-a52cf0d539e6"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-70d3fa25-6bbe-4a6b-a381-a52cf0d539e6.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T06:33:02.363Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-e6536f77-0d1b-4a6b-8f42-29cc15b2930a"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-e6536f77-0d1b-4a6b-8f42-29cc15b2930a"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T04:28:10.710Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-e6536f77-0d1b-4a6b-8f42-29cc15b2930a"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-e6536f77-0d1b-4a6b-8f42-29cc15b2930a.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T08:43:51.587Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sdk-Namespace-2924"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-2924"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T22:26:36.760Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:sdk-namespace-2924"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://sdk-Namespace-2924.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T22:26:59.350Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("rrama-sb1"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/rrama-sb1"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-01T21:47:34.903Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:rrama-sb1"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://rrama-sb1.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-02T02:10:03.083Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("WhackWhack"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/RapscallionResources/providers/Microsoft.ServiceBus/namespaces/WhackWhack"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-10T23:39:01.347Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:whackwhack"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://WhackWhack.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-04T00:56:32.687Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-66ed32d6-611e-4bb0-8e1a-a6d0fc65427c"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-66ed32d6-611e-4bb0-8e1a-a6d0fc65427c"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-16T17:51:27.730Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-66ed32d6-611e-4bb0-8e1a-a6d0fc65427c"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-66ed32d6-611e-4bb0-8e1a-a6d0fc65427c.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T08:19:43.383Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NS-e0cab401-6df8-465d-8d4a-da9a9e55cf0e"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-e0cab401-6df8-465d-8d4a-da9a9e55cf0e"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T01:14:25.613Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-e0cab401-6df8-465d-8d4a-da9a9e55cf0e"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://NS-e0cab401-6df8-465d-8d4a-da9a9e55cf0e.servicebus.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-11T12:33:01.727Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("bn3-rrama-foo1"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/bn3-rrama-foo1"),
		// 			Location: to.Ptr("East US 2"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-28T23:54:26.927Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:bn3-rrama-foo1"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://bn3-rrama-foo1.servicebus.int7.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-28T23:54:26.927Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("bn3-rrama-foo3"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/bn3-rrama-foo3"),
		// 			Location: to.Ptr("East US 2"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-29T00:24:09.907Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:bn3-rrama-foo3"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://bn3-rrama-foo3.servicebus.int7.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-29T00:24:33.233Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("bn3-rrama-foo2"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/bn3-rrama-foo2"),
		// 			Location: to.Ptr("East US 2"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-28T23:57:40.820Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:bn3-rrama-foo2"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://bn3-rrama-foo2.servicebus.int7.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-28T23:57:40.820Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("db3-rrama-foo2"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/db3-rrama-foo2"),
		// 			Location: to.Ptr("North Europe"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armservicebus.SBNamespaceProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-29T00:10:43.463Z"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:db3-rrama-foo2"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceBusEndpoint: to.Ptr("https://db3-rrama-foo2.servicebus.int7.windows-int.net:443/"),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-29T00:11:09.133Z"); return t}()),
		// 			},
		// 			SKU: &armservicebus.SBSKU{
		// 				Name: to.Ptr(armservicebus.SKUNameStandard),
		// 				Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 			},
		// 	}},
		// }
	}
}
