package armredhatopenshifthcp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshifthcp/armredhatopenshifthcp"
)

// Generated from example definition: 2026-09-01-preview/NodePools_ListByParent_MaximumSet_Gen.json
func ExampleNodePoolsClient_NewListByParentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshifthcp.NewClientFactory("FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNodePoolsClient().NewListByParentPager("rgopenapi", "hcpCluster-name", nil)
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
		// page = armredhatopenshifthcp.NodePoolsClientListByParentResponse{
		// 	NodePoolListResult: armredhatopenshifthcp.NodePoolListResult{
		// 		Value: []*armredhatopenshifthcp.NodePool{
		// 			{
		// 				Properties: &armredhatopenshifthcp.NodePoolProperties{
		// 					ProvisioningState: to.Ptr(armredhatopenshifthcp.ProvisioningStateSucceeded),
		// 					Version: &armredhatopenshifthcp.NodePoolVersionProfile{
		// 						ChannelGroup: to.Ptr("stable"),
		// 						ID: to.Ptr("4.12"),
		// 					},
		// 					Platform: &armredhatopenshifthcp.NodePoolPlatformProfile{
		// 						SubnetID: to.Ptr("/subscriptions/FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/hcp-network-example/subnets/example-subnet"),
		// 						VMSize: to.Ptr("Standard_D2s_v3"),
		// 						AvailabilityZone: to.Ptr("australiaeast-az1"),
		// 						EnableEncryptionAtHost: to.Ptr(true),
		// 						OSDisk: &armredhatopenshifthcp.OsDiskProfile{
		// 							SizeGiB: to.Ptr[int32](64),
		// 							DiskStorageAccountType: to.Ptr(armredhatopenshifthcp.DiskStorageAccountTypePremiumLRS),
		// 							EncryptionSetID: to.Ptr("/subscriptions/FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D/resourceGroups/resourceGroupName/providers/Microsoft.Compute/diskEncryptionSets/hcp-disk-encryption-set-example"),
		// 							DiskType: to.Ptr(armredhatopenshifthcp.OsDiskTypeManaged),
		// 						},
		// 					},
		// 					AutoScaling: &armredhatopenshifthcp.NodePoolAutoScaling{
		// 						Min: to.Ptr[int32](6),
		// 						Max: to.Ptr[int32](29),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key": to.Ptr("value"),
		// 				},
		// 				Location: to.Ptr("mqewzbuvnyxnwbmir"),
		// 				ID: to.Ptr("/subscriptions/FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D/resourceGroups/resourceGroupName/providers/Microsoft.RedHatOpenShift/resourceType/resourceName"),
		// 				Name: to.Ptr("riywfucwvfwoepzliopnphdfjw"),
		// 				Type: to.Ptr("znmdhkzcopsephiyom"),
		// 				SystemData: &armredhatopenshifthcp.SystemData{
		// 					CreatedBy: to.Ptr("iiqgrciyremxtwbrkjqtvcjkn"),
		// 					CreatedByType: to.Ptr(armredhatopenshifthcp.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2024, time.March, 25, 11, 14, 17, 555000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("ylhwjaq"),
		// 					LastModifiedByType: to.Ptr(armredhatopenshifthcp.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2024, time.March, 25, 11, 14, 17, 555000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
