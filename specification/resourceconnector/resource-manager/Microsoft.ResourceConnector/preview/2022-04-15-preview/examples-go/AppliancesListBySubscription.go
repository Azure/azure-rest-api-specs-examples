package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a7de71ef99f5ea2aefe38bbd3c55db09c64547e8/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/AppliancesListBySubscription.json
func ExampleAppliancesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ApplianceListResult = armresourceconnector.ApplianceListResult{
		// 	Value: []*armresourceconnector.Appliance{
		// 		{
		// 			Name: to.Ptr("appliance01"),
		// 			Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance01"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armresourceconnector.ApplianceProperties{
		// 				Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
		// 				InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
		// 					Provider: to.Ptr(armresourceconnector.ProviderVMWare),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Status: to.Ptr(armresourceconnector.StatusConnected),
		// 				Version: to.Ptr("1.0.1"),
		// 			},
		// 			SystemData: &armresourceconnector.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("appliance02"),
		// 			Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance02"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armresourceconnector.ApplianceProperties{
		// 				Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
		// 				InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
		// 					Provider: to.Ptr(armresourceconnector.ProviderVMWare),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Status: to.Ptr(armresourceconnector.StatusConnected),
		// 				Version: to.Ptr("1.0.1"),
		// 			},
		// 			SystemData: &armresourceconnector.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 			},
		// 	}},
		// }
	}
}
