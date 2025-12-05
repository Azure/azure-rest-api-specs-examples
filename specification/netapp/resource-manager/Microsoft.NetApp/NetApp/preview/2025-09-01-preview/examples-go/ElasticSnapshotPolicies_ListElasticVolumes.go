package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/ElasticSnapshotPolicies_ListElasticVolumes.json
func ExampleElasticSnapshotPoliciesClient_NewListElasticVolumesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticSnapshotPoliciesClient().NewListElasticVolumesPager("myRG", "account1", "snapshotPolicyName", nil)
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
		// page = armnetapp.ElasticSnapshotPoliciesClientListElasticVolumesResponse{
		// 	ElasticSnapshotPolicyVolumeList: armnetapp.ElasticSnapshotPolicyVolumeList{
		// 		Value: []*armnetapp.ElasticVolume{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticCapacityPools/pool1/elasticVolumes/volume1"),
		// 				Name: to.Ptr("account1/pool1/volume1"),
		// 				Type: to.Ptr("Microsoft.NetApp/elasticAccounts/elasticCapacityPools/elasticVolumes"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetapp.ElasticVolumeProperties{
		// 					FilePath: to.Ptr("my-unique-file-path"),
		// 					Size: to.Ptr[int64](107374182400),
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					ExportPolicy: &armnetapp.ElasticExportPolicy{
		// 						Rules: []*armnetapp.ElasticExportPolicyRule{
		// 							{
		// 								RuleIndex: to.Ptr[int32](1),
		// 								UnixAccessRule: to.Ptr(armnetapp.ElasticUnixAccessRuleReadOnly),
		// 								Nfsv3: to.Ptr(armnetapp.ElasticNfsv3AccessEnabled),
		// 								Nfsv4: to.Ptr(armnetapp.ElasticNfsv4AccessDisabled),
		// 								AllowedClients: []*string{
		// 									to.Ptr("0.0.0.0/0"),
		// 								},
		// 								RootAccess: to.Ptr(armnetapp.ElasticRootAccessDisabled),
		// 							},
		// 						},
		// 					},
		// 					ProtocolTypes: []*armnetapp.ElasticProtocolType{
		// 						to.Ptr(armnetapp.ElasticProtocolTypeNFSv3),
		// 					},
		// 					MountTargets: []*armnetapp.ElasticMountTargetProperties{
		// 						{
		// 							IPAddress: to.Ptr("10.1.0.9"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
