package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: 2025-03-01-preview/AppliancesListBySubscription.json
func ExampleAppliancesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppliancesClient().NewListBySubscriptionPager(nil)
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
		// page = armresourceconnector.AppliancesClientListBySubscriptionResponse{
		// 	ApplianceListResult: armresourceconnector.ApplianceListResult{
		// 		Value: []*armresourceconnector.Appliance{
		// 			{
		// 				ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance01"),
		// 				Name: to.Ptr("appliance01"),
		// 				Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armresourceconnector.ApplianceProperties{
		// 					Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
		// 					Version: to.Ptr("1.0.1"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					Status: to.Ptr(armresourceconnector.StatusRunning),
		// 					PublicKey: to.Ptr("xxxxxxxx"),
		// 					InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
		// 						Provider: to.Ptr(armresourceconnector.ProviderVMWare),
		// 					},
		// 					NetworkProfile: &armresourceconnector.NetworkProfile{
		// 						ProxyConfiguration: &armresourceconnector.ProxyConfiguration{
		// 							Version: to.Ptr("1744221145"),
		// 						},
		// 						DNSConfiguration: &armresourceconnector.DNSConfiguration{
		// 							Version: to.Ptr("1744221145"),
		// 						},
		// 						GatewayConfiguration: &armresourceconnector.GatewayConfiguration{
		// 							Version: to.Ptr("1744221145"),
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armresourceconnector.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-09T17:52:25.0928001Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-09T17:52:25.0928001Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance02"),
		// 				Name: to.Ptr("appliance02"),
		// 				Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armresourceconnector.ApplianceProperties{
		// 					Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
		// 					Version: to.Ptr("1.0.1"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					Status: to.Ptr(armresourceconnector.StatusRunning),
		// 					PublicKey: to.Ptr("xxxxxxxx"),
		// 					InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
		// 						Provider: to.Ptr(armresourceconnector.ProviderVMWare),
		// 					},
		// 					NetworkProfile: &armresourceconnector.NetworkProfile{
		// 						ProxyConfiguration: &armresourceconnector.ProxyConfiguration{
		// 							Version: to.Ptr("1744221145"),
		// 						},
		// 						DNSConfiguration: &armresourceconnector.DNSConfiguration{
		// 							Version: to.Ptr("1744221145"),
		// 						},
		// 						GatewayConfiguration: &armresourceconnector.GatewayConfiguration{
		// 							Version: to.Ptr("1744221145"),
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armresourceconnector.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-09T17:52:25.0928001Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-09T17:52:25.0928001Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
