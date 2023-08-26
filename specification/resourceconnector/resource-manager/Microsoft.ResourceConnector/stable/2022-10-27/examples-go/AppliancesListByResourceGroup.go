package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListByResourceGroup.json
func ExampleAppliancesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppliancesClient().NewListByResourceGroupPager("testresourcegroup", nil)
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
		// 			SystemData: &armresourceconnector.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 			},
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
		// 		},
		// 		{
		// 			Name: to.Ptr("appliance02"),
		// 			Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance02"),
		// 			SystemData: &armresourceconnector.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 			},
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
		// 	}},
		// }
	}
}
