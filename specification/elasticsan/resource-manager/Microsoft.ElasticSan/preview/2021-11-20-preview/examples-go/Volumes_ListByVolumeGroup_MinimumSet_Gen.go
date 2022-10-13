package armelasticsans_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsans/armelasticsans"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/Volumes_ListByVolumeGroup_MinimumSet_Gen.json
func ExampleVolumesClient_NewListByVolumeGroupPager_volumesListByVolumeGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armelasticsans.NewVolumesClient("aaaaaaaaaaaaaaaaaa", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByVolumeGroupPager("rgelasticsan", "ti7q-k952-1qB3J_5", "u_5I_1j4t3", nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
