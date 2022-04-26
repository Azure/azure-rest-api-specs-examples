Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.6.0/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-create.json
func ExampleJobsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmediaservices.NewJobsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<transform-name>",
		"<job-name>",
		armmediaservices.Job{
			Properties: &armmediaservices.JobProperties{
				CorrelationData: map[string]*string{
					"Key 2": to.Ptr("Value 2"),
					"key1":  to.Ptr("value1"),
				},
				Input: &armmediaservices.JobInputAsset{
					ODataType: to.Ptr("<odata-type>"),
					AssetName: to.Ptr("<asset-name>"),
				},
				Outputs: []armmediaservices.JobOutputClassification{
					&armmediaservices.JobOutputAsset{
						ODataType: to.Ptr("<odata-type>"),
						AssetName: to.Ptr("<asset-name>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
```
