package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-create.json
func ExampleJobsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewJobsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contosoresources", "contosomedia", "exampleTransform", "job1", armmediaservices.Job{
		Properties: &armmediaservices.JobProperties{
			CorrelationData: map[string]*string{
				"Key 2": to.Ptr("Value 2"),
				"key1":  to.Ptr("value1"),
			},
			Input: &armmediaservices.JobInputAsset{
				ODataType: to.Ptr("#Microsoft.Media.JobInputAsset"),
				AssetName: to.Ptr("job1-InputAsset"),
			},
			Outputs: []armmediaservices.JobOutputClassification{
				&armmediaservices.JobOutputAsset{
					ODataType: to.Ptr("#Microsoft.Media.JobOutputAsset"),
					AssetName: to.Ptr("job1-OutputAsset"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
