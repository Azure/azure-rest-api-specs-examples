package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2024-12-01/SolutionConfigurations_List.json
func ExampleSolutionConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSolutionConfigurationsClient().NewListPager("ymuj", nil)
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
		// page = armhybridconnectivity.SolutionConfigurationsClientListResponse{
		// 	SolutionConfigurationListResult: armhybridconnectivity.SolutionConfigurationListResult{
		// 		Value: []*armhybridconnectivity.SolutionConfiguration{
		// 			{
		// 				Properties: &armhybridconnectivity.SolutionConfigurationProperties{
		// 					SolutionType: to.Ptr("Microsoft.AssetManagement"),
		// 					SolutionSettings: &armhybridconnectivity.SolutionSettings{
		// 					},
		// 					ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
		// 					StatusDetails: to.Ptr("Aws authorization validation pending in Aws account"),
		// 					LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-20T03:24:15.820Z"); return t}()),
		// 				},
		// 				ID: to.Ptr("/subscriptions/testSubcrptions/resourceGroups/testResourceGroup/providers/Microsoft.HybridConnectivity/SolutionConfigurations/solutionconfigurationtest"),
		// 				Name: to.Ptr("solutionconfigurationtest"),
		// 				Type: to.Ptr("microsoft.hybridconnectivity/solutionconfigurations"),
		// 				SystemData: &armhybridconnectivity.SystemData{
		// 					CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
		// 					CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("jidegyskxi"),
		// 					LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armhybridconnectivity.SolutionConfigurationProperties{
		// 					SolutionType: to.Ptr("Microsoft.HybridCompute"),
		// 					SolutionSettings: &armhybridconnectivity.SolutionSettings{
		// 					},
		// 					ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
		// 					StatusDetails: to.Ptr("Aws authorization validation succeeded in Aws account"),
		// 					LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-20T03:24:15.820Z"); return t}()),
		// 				},
		// 				ID: to.Ptr("/subscriptions/testSubcrptions/resourceGroups/testResourceGroup/providers/Microsoft.HybridConnectivity/SolutionConfigurations/solutionconfigurationtest2"),
		// 				Name: to.Ptr("solutionconfigurationtest2"),
		// 				Type: to.Ptr("microsoft.hybridconnectivity/solutionconfigurations"),
		// 				SystemData: &armhybridconnectivity.SystemData{
		// 					CreatedBy: to.Ptr("rpxzkcrobprrdvuoqxz"),
		// 					CreatedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("jidegyskxi"),
		// 					LastModifiedByType: to.Ptr(armhybridconnectivity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-18T22:52:07.890Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
