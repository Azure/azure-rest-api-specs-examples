package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1da7cbab8d4f554484dedb676ba7bdbdf6cdf78/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/Volumes_Create_MaximumSet_Gen.json
func ExampleVolumesClient_BeginCreate_volumesCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginCreate(ctx, "rgelasticsan", "ti7q-k952-1qB3J_5", "u_5I_1j4t3", "9132y", armelasticsan.Volume{
		Tags: map[string]*string{
			"key7423": to.Ptr("aaaa"),
		},
		Properties: &armelasticsan.VolumeProperties{
			CreationData: &armelasticsan.SourceCreationData{
				CreateSource: to.Ptr("None"),
				SourceURI:    to.Ptr("aaaaaa"),
			},
			SizeGiB: to.Ptr[int64](22),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Volume = armelasticsan.Volume{
	// 	Name: to.Ptr("aaaaaaaaaaaa"),
	// 	Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 	ID: to.Ptr("aaaaaaaa"),
	// 	Tags: map[string]*string{
	// 		"key7423": to.Ptr("aaaa"),
	// 	},
	// 	Properties: &armelasticsan.VolumeProperties{
	// 		CreationData: &armelasticsan.SourceCreationData{
	// 			CreateSource: to.Ptr("None"),
	// 			SourceURI: to.Ptr("aaaaaa"),
	// 		},
	// 		SizeGiB: to.Ptr[int64](22),
	// 		StorageTarget: &armelasticsan.IscsiTargetInfo{
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			Status: to.Ptr(armelasticsan.OperationalStatusInvalid),
	// 			TargetIqn: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 			TargetPortalHostname: to.Ptr("aaaaaaa"),
	// 			TargetPortalPort: to.Ptr[int32](23),
	// 		},
	// 		VolumeID: to.Ptr("aaaaaaaaaa"),
	// 	},
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-08T10:39:37.620Z"); return t}()),
	// 		CreatedBy: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-08T10:39:37.620Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// }
}
