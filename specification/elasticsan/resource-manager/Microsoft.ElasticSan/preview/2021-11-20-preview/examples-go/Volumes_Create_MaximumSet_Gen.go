package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/Volumes_Create_MaximumSet_Gen.json
func ExampleVolumesClient_BeginCreate_volumesCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armelasticsan.NewVolumesClient("aaaaaaaaaaaaaaaaaa", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "rgelasticsan", "ti7q-k952-1qB3J_5", "u_5I_1j4t3", "9132y", armelasticsan.Volume{
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
	// TODO: use response item
	_ = res
}
