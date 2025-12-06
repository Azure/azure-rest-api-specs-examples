package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/ElasticVolumes_Get.json
func ExampleElasticVolumesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewElasticVolumesClient().Get(ctx, "myRG", "account1", "pool1", "volume1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.ElasticVolumesClientGetResponse{
	// 	ElasticVolume: &armnetapp.ElasticVolume{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticCapacityPools/pool1/elasticVolumes/volume1"),
	// 		Name: to.Ptr("account1/pool1/volume1"),
	// 		Type: to.Ptr("Microsoft.NetApp/elasticAccounts/elasticCapacityPools/elasticVolumes"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.ElasticVolumeProperties{
	// 			FilePath: to.Ptr("my-unique-file-path"),
	// 			Size: to.Ptr[int64](107374182400),
	// 			ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 			ExportPolicy: &armnetapp.ElasticExportPolicy{
	// 				Rules: []*armnetapp.ElasticExportPolicyRule{
	// 					{
	// 						RuleIndex: to.Ptr[int32](1),
	// 						UnixAccessRule: to.Ptr(armnetapp.ElasticUnixAccessRuleReadOnly),
	// 						Nfsv3: to.Ptr(armnetapp.ElasticNfsv3AccessEnabled),
	// 						Nfsv4: to.Ptr(armnetapp.ElasticNfsv4AccessDisabled),
	// 						AllowedClients: []*string{
	// 							to.Ptr("0.0.0.0/0"),
	// 						},
	// 						RootAccess: to.Ptr(armnetapp.ElasticRootAccessDisabled),
	// 					},
	// 				},
	// 			},
	// 			ProtocolTypes: []*armnetapp.ElasticProtocolType{
	// 				to.Ptr(armnetapp.ElasticProtocolTypeNFSv3),
	// 			},
	// 			MountTargets: []*armnetapp.ElasticMountTargetProperties{
	// 				{
	// 					IPAddress: to.Ptr("10.1.0.9"),
	// 				},
	// 			},
	// 			RestorationState: to.Ptr(armnetapp.ElasticVolumeRestorationStateRestoring),
	// 		},
	// 	},
	// }
}
