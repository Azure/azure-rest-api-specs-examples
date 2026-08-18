package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/NetworkManagerCommitList.json
func ExampleCommitsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommitsClient().NewListPager("myResourceGroup", "testNetworkManager", nil)
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
		// page = armnetwork.CommitsClientListResponse{
		// 	CommitListResult: armnetwork.CommitListResult{
		// 		Value: []*armnetwork.Commit{
		// 			{
		// 				Name: to.Ptr("myTestCommit"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/commits/myTestCommit"),
		// 				Type: to.Ptr("Microsoft.Network/networkManagers/commits"),
		// 				SystemData: &armnetwork.SystemData{
		// 					CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 					CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2025, time.January, 11, 18, 52, 27, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 					LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.January, 11, 18, 52, 27, 0, time.UTC)),
		// 				},
		// 				Properties: &armnetwork.CommitProperties{
		// 					Description: to.Ptr("Sample Commit"),
		// 					TargetLocations: []*string{
		// 						to.Ptr("useast"),
		// 					},
		// 					ActiveLocations: []*string{
		// 						to.Ptr("useast"),
		// 					},
		// 					ConfigurationIDs: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resoureGroupSample/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/SampleSecurityAdminConfig"),
		// 					},
		// 					CommitType: to.Ptr(armnetwork.ConfigurationTypeSecurityAdmin),
		// 					ForceUpdateTag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkManagers/testNetworkManager/commits?api-version=2025-07-01&$skipToken=10"),
		// 	},
		// }
	}
}
