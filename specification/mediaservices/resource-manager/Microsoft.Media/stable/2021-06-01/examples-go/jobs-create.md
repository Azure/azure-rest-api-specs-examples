```go
package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/jobs-create.json
func ExampleJobsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewJobsClient("<subscription-id>", cred, nil)
	_, err = client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<transform-name>",
		"<job-name>",
		armmediaservices.Job{
			Properties: &armmediaservices.JobProperties{
				CorrelationData: map[string]*string{
					"Key 2": to.StringPtr("Value 2"),
					"key1":  to.StringPtr("value1"),
				},
				Input: &armmediaservices.JobInputAsset{
					ODataType: to.StringPtr("<odata-type>"),
					AssetName: to.StringPtr("<asset-name>"),
				},
				Outputs: []armmediaservices.JobOutputClassification{
					&armmediaservices.JobOutputAsset{
						ODataType: to.StringPtr("<odata-type>"),
						AssetName: to.StringPtr("<asset-name>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.3.1/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.
