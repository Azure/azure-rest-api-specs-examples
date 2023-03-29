package armanalysisservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/listServersInResourceGroup.json
func ExampleServersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armanalysisservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServersClient().NewListByResourceGroupPager("TestRG", nil)
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
		// page.Servers = armanalysisservices.Servers{
		// 	Value: []*armanalysisservices.Server{
		// 		{
		// 			Name: to.Ptr("azsdktest"),
		// 			ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.AnalysisServices/servers/azsdktest"),
		// 			Location: to.Ptr("West US"),
		// 			SKU: &armanalysisservices.ResourceSKU{
		// 				Name: to.Ptr("S1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Tier: to.Ptr(armanalysisservices.SKUTierStandard),
		// 			},
		// 			Tags: map[string]*string{
		// 				"testKey": to.Ptr("testValue"),
		// 			},
		// 			Properties: &armanalysisservices.ServerProperties{
		// 				AsAdministrators: &armanalysisservices.ServerAdministrators{
		// 					Members: []*string{
		// 						to.Ptr("azsdktest@microsoft.com")},
		// 					},
		// 					ProvisioningState: to.Ptr(armanalysisservices.ProvisioningStateSucceeded),
		// 					ServerFullName: to.Ptr("asazure://nightly1.asazure-int.windows.net/azsdktest"),
		// 					State: to.Ptr(armanalysisservices.StateProvisioning),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("azsdktest2"),
		// 				ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.AnalysisServices/servers/azsdktest"),
		// 				Location: to.Ptr("West US"),
		// 				SKU: &armanalysisservices.ResourceSKU{
		// 					Name: to.Ptr("S0"),
		// 					Capacity: to.Ptr[int32](1),
		// 					Tier: to.Ptr(armanalysisservices.SKUTierStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 					"testKey": to.Ptr("testValue"),
		// 				},
		// 				Properties: &armanalysisservices.ServerProperties{
		// 					AsAdministrators: &armanalysisservices.ServerAdministrators{
		// 						Members: []*string{
		// 							to.Ptr("azsdktest@microsoft.com")},
		// 						},
		// 						ProvisioningState: to.Ptr(armanalysisservices.ProvisioningStateSucceeded),
		// 						ServerFullName: to.Ptr("asazure://nightly1.asazure-int.windows.net/azsdktest"),
		// 						State: to.Ptr(armanalysisservices.StateProvisioning),
		// 					},
		// 			}},
		// 		}
	}
}
