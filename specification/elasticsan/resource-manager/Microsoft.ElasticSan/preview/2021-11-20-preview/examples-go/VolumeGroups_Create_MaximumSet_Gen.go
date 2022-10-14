package armelasticsans_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsans/armelasticsans"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/VolumeGroups_Create_MaximumSet_Gen.json
func ExampleVolumeGroupsClient_BeginCreate_volumeGroupsCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armelasticsans.NewVolumeGroupsClient("aaaaaaaaaaaaaaaaaa", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "rgelasticsan", "ti7q-k952-1qB3J_5", "u_5I_1j4t3", armelasticsans.VolumeGroup{
		Tags: map[string]*string{
			"key5933": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		},
		Properties: &armelasticsans.VolumeGroupProperties{
			Encryption: to.Ptr(armelasticsans.EncryptionTypeEncryptionAtRestWithPlatformKey),
			NetworkACLs: &armelasticsans.NetworkRuleSet{
				VirtualNetworkRules: []*armelasticsans.VirtualNetworkRule{
					{
						Action:                   to.Ptr("Allow"),
						VirtualNetworkResourceID: to.Ptr("aaaaaaaaaaaaaaaa"),
					}},
			},
			ProtocolType: to.Ptr(armelasticsans.StorageTargetTypeIscsi),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
