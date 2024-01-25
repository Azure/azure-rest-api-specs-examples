package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Volumes_Delete_MaximumSet_Gen.json
func ExampleVolumesClient_BeginDelete_volumesDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginDelete(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "volumename", &armelasticsan.VolumesClientBeginDeleteOptions{XMSDeleteSnapshots: to.Ptr(armelasticsan.XMSDeleteSnapshotsTrue),
		XMSForceDelete: to.Ptr(armelasticsan.XMSForceDeleteTrue),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
