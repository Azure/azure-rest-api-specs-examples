package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2024-12-01/Inventory_ListBySolutionConfiguration.json
func ExampleInventoryClient_NewListBySolutionConfigurationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInventoryClient().NewListBySolutionConfigurationPager("ymuj", "wsxt", nil)
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
		// page = armhybridconnectivity.InventoryClientListBySolutionConfigurationResponse{
		// 	InventoryResourceListResult: armhybridconnectivity.InventoryResourceListResult{
		// 		Value: []*armhybridconnectivity.InventoryResource{
		// 			{
		// 				Properties: &armhybridconnectivity.InventoryProperties{
		// 					CloudNativeType: to.Ptr(armhybridconnectivity.CloudNativeTypeEc2),
		// 					CloudNativeResourceID: to.Ptr("ljnxclzxficxhdkr"),
		// 					AzureResourceID: to.Ptr("ttzebbjzatugawuqxdupzmxkt"),
		// 					Status: to.Ptr(armhybridconnectivity.SolutionConfigurationStatusNew),
		// 					StatusDetails: to.Ptr("wxvnfzivtx"),
		// 					ProvisioningState: to.Ptr(armhybridconnectivity.ResourceProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/testSubcrptions/resourceGroups/testResourceGroup/providers/Microsoft.HybridConnectivity/SolutionConfigurations/qpwubemzmootxmtlxaerir/Inventory/xipjenocwvsqhhrplwmxwl"),
		// 				Name: to.Ptr("xipjenocwvsqhhrplwmxwl"),
		// 				Type: to.Ptr("jahwzrspsufypeouigsywjrx"),
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
