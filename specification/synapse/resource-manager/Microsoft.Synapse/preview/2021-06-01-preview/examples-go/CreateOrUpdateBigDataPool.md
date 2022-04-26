Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.4.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/CreateOrUpdateBigDataPool.json
func ExampleBigDataPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewBigDataPoolsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<big-data-pool-name>",
		armsynapse.BigDataPoolResourceInfo{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
			Properties: &armsynapse.BigDataPoolResourceProperties{
				AutoPause: &armsynapse.AutoPauseProperties{
					DelayInMinutes: to.Ptr[int32](15),
					Enabled:        to.Ptr(true),
				},
				AutoScale: &armsynapse.AutoScaleProperties{
					Enabled:      to.Ptr(true),
					MaxNodeCount: to.Ptr[int32](50),
					MinNodeCount: to.Ptr[int32](3),
				},
				DefaultSparkLogFolder: to.Ptr("<default-spark-log-folder>"),
				LibraryRequirements: &armsynapse.LibraryRequirements{
					Content:  to.Ptr("<content>"),
					Filename: to.Ptr("<filename>"),
				},
				NodeCount:         to.Ptr[int32](4),
				NodeSize:          to.Ptr(armsynapse.NodeSizeMedium),
				NodeSizeFamily:    to.Ptr(armsynapse.NodeSizeFamilyMemoryOptimized),
				SparkEventsFolder: to.Ptr("<spark-events-folder>"),
				SparkVersion:      to.Ptr("<spark-version>"),
			},
		},
		&armsynapse.BigDataPoolsClientBeginCreateOrUpdateOptions{Force: nil,
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
