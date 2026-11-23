package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: 2026-07-01-preview/NameSpaces/SBNameSpaceList.json
func ExampleNamespacesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
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
		// page = armservicebus.NamespacesClientListResponse{
		// 	SBNamespaceListResult: armservicebus.SBNamespaceListResult{
		// 		Value: []*armservicebus.SBNamespace{
		// 			{
		// 				Name: to.Ptr("NS-91f08e47-2b04-4943-b0cd-a5fb02b88f20"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-91f08e47-2b04-4943-b0cd-a5fb02b88f20"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 2, 40, 17, 270000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-91f08e47-2b04-4943-b0cd-a5fb02b88f20"),
		// 					MinimumTLSVersion: to.Ptr(armservicebus.TLSVersionOne2),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-91f08e47-2b04-4943-b0cd-a5fb02b88f20.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 7, 15, 30, 780000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-41dc63f4-0b08-4029-b3ef-535a131bfa65"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-41dc63f4-0b08-4029-b3ef-535a131bfa65"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 3, 50, 38, 980000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-41dc63f4-0b08-4029-b3ef-535a131bfa65"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-41dc63f4-0b08-4029-b3ef-535a131bfa65.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 10, 42, 58, 3000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-df52cf51-e831-4bf2-bd92-e9885f68a996"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-df52cf51-e831-4bf2-bd92-e9885f68a996"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 1, 17, 54, 997000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-df52cf51-e831-4bf2-bd92-e9885f68a996"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-df52cf51-e831-4bf2-bd92-e9885f68a996.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 6, 44, 39, 737000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("SBPremium"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/RapscallionResources/providers/Microsoft.ServiceBus/namespaces/SBPremium"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.October, 10, 22, 1, 0, 420000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:sbpremium"),
		// 					PremiumMessagingPartitions: to.Ptr[int32](1),
		// 					ProvisioningState: to.Ptr("Created"),
		// 					ServiceBusEndpoint: to.Ptr("https://SBPremium.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2016, time.October, 10, 22, 1, 0, 420000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNamePremium),
		// 					Capacity: to.Ptr[int32](1),
		// 					Tier: to.Ptr(armservicebus.SKUTierPremium),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("rrama-ns2"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/sadfsadfsadf/providers/Microsoft.ServiceBus/namespaces/rrama-ns2"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 4, 14, 0, 13000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:rrama-ns2"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://rrama-ns2.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 3, 22, 53, 32, 927000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-20e57600-29d0-4035-ac85-74f4c54dcda1"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-20e57600-29d0-4035-ac85-74f4c54dcda1"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 3, 30, 49, 160000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-20e57600-29d0-4035-ac85-74f4c54dcda1"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-20e57600-29d0-4035-ac85-74f4c54dcda1.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 4, 17, 58, 483000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-3e538a1a-58fb-4315-b2ce-76f5c944114c"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-3e538a1a-58fb-4315-b2ce-76f5c944114c"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 18, 7, 30, 50000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-3e538a1a-58fb-4315-b2ce-76f5c944114c"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-3e538a1a-58fb-4315-b2ce-76f5c944114c.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 10, 42, 57, 747000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("prem-ns123"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/prem-ns123"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 13, 0, 2, 39, 997000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:prem-ns123"),
		// 					PremiumMessagingPartitions: to.Ptr[int32](2),
		// 					ProvisioningState: to.Ptr("Created"),
		// 					ServiceBusEndpoint: to.Ptr("https://prem-ns123.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2016, time.September, 13, 0, 2, 39, 997000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNamePremium),
		// 					Capacity: to.Ptr[int32](4),
		// 					Tier: to.Ptr(armservicebus.SKUTierPremium),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-4e1bfdf1-0cff-4e86-ae80-cdcac4873039"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-4e1bfdf1-0cff-4e86-ae80-cdcac4873039"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 1, 1, 58, 730000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-4e1bfdf1-0cff-4e86-ae80-cdcac4873039"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-4e1bfdf1-0cff-4e86-ae80-cdcac4873039.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 3, 2, 59, 800000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-6b90b7f3-7aa0-48c9-bc30-b299dcb66c03"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-6b90b7f3-7aa0-48c9-bc30-b299dcb66c03"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 3, 22, 45, 327000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-6b90b7f3-7aa0-48c9-bc30-b299dcb66c03"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-6b90b7f3-7aa0-48c9-bc30-b299dcb66c03.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 6, 8, 1, 207000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-c05e9df3-7737-44ee-a321-15f6e0545b97"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-c05e9df3-7737-44ee-a321-15f6e0545b97"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 5, 3, 29, 19, 750000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-c05e9df3-7737-44ee-a321-15f6e0545b97"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-c05e9df3-7737-44ee-a321-15f6e0545b97.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 8, 10, 35, 527000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-dcb4152c-231b-4c16-a683-07cc6b38fa46"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-dcb4152c-231b-4c16-a683-07cc6b38fa46"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 5, 3, 34, 35, 363000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-dcb4152c-231b-4c16-a683-07cc6b38fa46"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-dcb4152c-231b-4c16-a683-07cc6b38fa46.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 5, 33, 0, 957000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-f501f5e6-1f24-439b-8982-9af665156d40"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-f501f5e6-1f24-439b-8982-9af665156d40"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 1, 25, 55, 707000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-f501f5e6-1f24-439b-8982-9af665156d40"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-f501f5e6-1f24-439b-8982-9af665156d40.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 7, 42, 59, 687000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-fe2ed660-2cd6-46f2-a9c3-7e11551a1f30"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-fe2ed660-2cd6-46f2-a9c3-7e11551a1f30"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 2, 32, 8, 227000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-fe2ed660-2cd6-46f2-a9c3-7e11551a1f30"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-fe2ed660-2cd6-46f2-a9c3-7e11551a1f30.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 6, 32, 57, 770000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-8a5e3b4e-4e97-4d85-9083-cd33536c9d71"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-8a5e3b4e-4e97-4d85-9083-cd33536c9d71"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 0, 54, 5, 103000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-8a5e3b4e-4e97-4d85-9083-cd33536c9d71"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-8a5e3b4e-4e97-4d85-9083-cd33536c9d71.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 10, 43, 50, 313000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-6520cc09-01ac-40a3-bc09-c5c431116e92"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-6520cc09-01ac-40a3-bc09-c5c431116e92"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 1, 49, 59, 243000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-6520cc09-01ac-40a3-bc09-c5c431116e92"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-6520cc09-01ac-40a3-bc09-c5c431116e92.servicebus.windows-int.net:443"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 8, 15, 36, 950000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-bfba6d5c-a425-42d9-85db-0f4da770e29a"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-bfba6d5c-a425-42d9-85db-0f4da770e29a"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 5, 3, 23, 32, 83000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-bfba6d5c-a425-42d9-85db-0f4da770e29a"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-bfba6d5c-a425-42d9-85db-0f4da770e29a.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 9, 2, 57, 433000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("SBPrem"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/RapscallionResources/providers/Microsoft.ServiceBus/namespaces/SBPrem"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.October, 10, 22, 16, 30, 870000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:sbprem"),
		// 					PremiumMessagingPartitions: to.Ptr[int32](2),
		// 					ProvisioningState: to.Ptr("Created"),
		// 					ServiceBusEndpoint: to.Ptr("https://SBPrem.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2016, time.October, 10, 22, 16, 30, 870000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNamePremium),
		// 					Capacity: to.Ptr[int32](1),
		// 					Tier: to.Ptr(armservicebus.SKUTierPremium),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-43b136b4-8716-40b2-97c5-0d77cac0062c"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-43b136b4-8716-40b2-97c5-0d77cac0062c"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 3, 14, 50, 577000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-43b136b4-8716-40b2-97c5-0d77cac0062c"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-43b136b4-8716-40b2-97c5-0d77cac0062c.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 9, 23, 1, 67000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-7c0443de-5f88-450c-b574-83f60a097dd1"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-7c0443de-5f88-450c-b574-83f60a097dd1"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 4, 7, 15, 397000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-7c0443de-5f88-450c-b574-83f60a097dd1"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-7c0443de-5f88-450c-b574-83f60a097dd1.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 4, 3, 3, 97000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-62dd7753-a5f9-42fd-a354-ca38a4505d69"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-62dd7753-a5f9-42fd-a354-ca38a4505d69"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 1, 33, 50, 450000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-62dd7753-a5f9-42fd-a354-ca38a4505d69"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-62dd7753-a5f9-42fd-a354-ca38a4505d69.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 5, 35, 33, 53000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-ae18a18c-97ab-4089-965d-8acbf4794091"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-ae18a18c-97ab-4089-965d-8acbf4794091"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 2, 43, 36, 517000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-ae18a18c-97ab-4089-965d-8acbf4794091"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-ae18a18c-97ab-4089-965d-8acbf4794091.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 12, 40, 30, 587000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-8e3b56c1-0ee8-4e13-ae88-5cadf6e2ce11"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-8e3b56c1-0ee8-4e13-ae88-5cadf6e2ce11"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 0, 46, 3, 773000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-8e3b56c1-0ee8-4e13-ae88-5cadf6e2ce11"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-8e3b56c1-0ee8-4e13-ae88-5cadf6e2ce11.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 4, 43, 54, 560000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-7ffca4b4-4728-4fb0-b2d0-1e7c016e3a44"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-7ffca4b4-4728-4fb0-b2d0-1e7c016e3a44"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 3, 59, 12, 100000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-7ffca4b4-4728-4fb0-b2d0-1e7c016e3a44"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-7ffca4b4-4728-4fb0-b2d0-1e7c016e3a44.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 6, 33, 52, 230000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-d9337efd-9b27-454c-b2a5-dcfea56920d9"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-d9337efd-9b27-454c-b2a5-dcfea56920d9"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 5, 3, 45, 9, 270000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-d9337efd-9b27-454c-b2a5-dcfea56920d9"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-d9337efd-9b27-454c-b2a5-dcfea56920d9.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 6, 20, 31, 863000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-ad5ae732-abea-4e62-9de0-c90de0ddec0a"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-ad5ae732-abea-4e62-9de0-c90de0ddec0a"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 2, 34, 36, 447000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-ad5ae732-abea-4e62-9de0-c90de0ddec0a"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-ad5ae732-abea-4e62-9de0-c90de0ddec0a.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 6, 15, 31, 607000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-d447fb03-c7da-40fe-b5eb-14f36888837b"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-d447fb03-c7da-40fe-b5eb-14f36888837b"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 5, 0, 53, 46, 697000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-d447fb03-c7da-40fe-b5eb-14f36888837b"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-d447fb03-c7da-40fe-b5eb-14f36888837b.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 11, 9, 41, 260000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("ReproSB"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/RapscallionResources/providers/Microsoft.ServiceBus/namespaces/ReproSB"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2017, time.February, 27, 19, 29, 34, 523000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:reprosb"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://ReproSB.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 27, 19, 29, 58, 640000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-4c90097f-19a8-42e7-bb3c-4ac088994719"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-4c90097f-19a8-42e7-bb3c-4ac088994719"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 17, 35, 32, 610000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-4c90097f-19a8-42e7-bb3c-4ac088994719"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-4c90097f-19a8-42e7-bb3c-4ac088994719.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 9, 13, 52, 270000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("rrama-1-23-17"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/rrama-1-23-17"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2017, time.January, 23, 22, 54, 40, 907000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:rrama-1-23-17"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://rrama-1-23-17.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 4, 0, 53, 28, 777000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-5191e541-8e4e-4229-9fdc-b89f6c3e7f12"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-5191e541-8e4e-4229-9fdc-b89f6c3e7f12"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 17, 43, 25, 710000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-5191e541-8e4e-4229-9fdc-b89f6c3e7f12"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-5191e541-8e4e-4229-9fdc-b89f6c3e7f12.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 11, 5, 31, 890000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-be903820-3533-46e8-90e4-72c132411848"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-be903820-3533-46e8-90e4-72c132411848"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 5, 3, 24, 1, 923000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-be903820-3533-46e8-90e4-72c132411848"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-be903820-3533-46e8-90e4-72c132411848.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 10, 9, 42, 513000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("rrama-namespace1"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/rrama-namespace1"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 5, 0, 47, 22, 963000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:rrama-namespace1"),
		// 					ProvisioningState: to.Ptr("Created"),
		// 					ServiceBusEndpoint: to.Ptr("https://rrama-namespace1.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2016, time.August, 5, 0, 47, 27, 297000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-a3c38e9b-32a3-4c51-85d7-263150a8dda9"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-a3c38e9b-32a3-4c51-85d7-263150a8dda9"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 0, 38, 2, 517000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-a3c38e9b-32a3-4c51-85d7-263150a8dda9"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-a3c38e9b-32a3-4c51-85d7-263150a8dda9.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 5, 3, 55, 960000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-70d3fa25-6bbe-4a6b-a381-a52cf0d539e6"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-70d3fa25-6bbe-4a6b-a381-a52cf0d539e6"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 23, 3, 42, 40, 10000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-70d3fa25-6bbe-4a6b-a381-a52cf0d539e6"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-70d3fa25-6bbe-4a6b-a381-a52cf0d539e6.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 6, 33, 2, 363000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-e6536f77-0d1b-4a6b-8f42-29cc15b2930a"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-e6536f77-0d1b-4a6b-8f42-29cc15b2930a"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 5, 4, 28, 10, 710000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-e6536f77-0d1b-4a6b-8f42-29cc15b2930a"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-e6536f77-0d1b-4a6b-8f42-29cc15b2930a.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 8, 43, 51, 587000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sdk-Namespace-2924"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-2924"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2017, time.May, 25, 22, 26, 36, 760000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:sdk-namespace-2924"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://sdk-Namespace-2924.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.May, 25, 22, 26, 59, 350000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("value1"),
		// 					"tag2": to.Ptr("value2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("rrama-sb1"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/rrama-sb1"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2017, time.May, 1, 21, 47, 34, 903000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:rrama-sb1"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://rrama-sb1.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.May, 2, 2, 10, 3, 83000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("WhackWhack"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/RapscallionResources/providers/Microsoft.ServiceBus/namespaces/WhackWhack"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.October, 10, 23, 39, 1, 347000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:whackwhack"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://WhackWhack.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 4, 0, 56, 32, 687000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-66ed32d6-611e-4bb0-8e1a-a6d0fc65427c"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-66ed32d6-611e-4bb0-8e1a-a6d0fc65427c"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.September, 16, 17, 51, 27, 730000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-66ed32d6-611e-4bb0-8e1a-a6d0fc65427c"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-66ed32d6-611e-4bb0-8e1a-a6d0fc65427c.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 8, 19, 43, 383000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("NS-e0cab401-6df8-465d-8d4a-da9a9e55cf0e"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/NS-e0cab401-6df8-465d-8d4a-da9a9e55cf0e"),
		// 				Location: to.Ptr("South Central US"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2016, time.August, 5, 1, 14, 25, 613000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:ns-e0cab401-6df8-465d-8d4a-da9a9e55cf0e"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://NS-e0cab401-6df8-465d-8d4a-da9a9e55cf0e.servicebus.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.February, 11, 12, 33, 1, 727000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("bn3-rrama-foo1"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/bn3-rrama-foo1"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2017, time.April, 28, 23, 54, 26, 927000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:bn3-rrama-foo1"),
		// 					ProvisioningState: to.Ptr("Created"),
		// 					ServiceBusEndpoint: to.Ptr("https://bn3-rrama-foo1.servicebus.int7.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.April, 28, 23, 54, 26, 927000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("bn3-rrama-foo3"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/bn3-rrama-foo3"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2017, time.April, 29, 0, 24, 9, 907000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:bn3-rrama-foo3"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://bn3-rrama-foo3.servicebus.int7.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.April, 29, 0, 24, 33, 233000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("bn3-rrama-foo2"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/bn3-rrama-foo2"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2017, time.April, 28, 23, 57, 40, 820000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:bn3-rrama-foo2"),
		// 					ProvisioningState: to.Ptr("Created"),
		// 					ServiceBusEndpoint: to.Ptr("https://bn3-rrama-foo2.servicebus.int7.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.April, 28, 23, 57, 40, 820000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("db3-rrama-foo2"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/Default-ServiceBus-SouthCentralUS/providers/Microsoft.ServiceBus/namespaces/db3-rrama-foo2"),
		// 				Location: to.Ptr("North Europe"),
		// 				Properties: &armservicebus.SBNamespaceProperties{
		// 					CreatedAt: to.Ptr(time.Date(2017, time.April, 29, 0, 10, 43, 463000000, time.UTC)),
		// 					DisableLocalAuth: to.Ptr(false),
		// 					MetricID: to.Ptr("5f750a97-50d9-4e36-8081-c9ee4c0210d4:db3-rrama-foo2"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceBusEndpoint: to.Ptr("https://db3-rrama-foo2.servicebus.int7.windows-int.net:443/"),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.April, 29, 0, 11, 9, 133000000, time.UTC)),
		// 				},
		// 				SKU: &armservicebus.SBSKU{
		// 					Name: to.Ptr(armservicebus.SKUNameStandard),
		// 					Tier: to.Ptr(armservicebus.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
