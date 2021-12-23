Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.1.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/CreateOrUpdateBigDataPool.json
func ExampleBigDataPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewBigDataPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<big-data-pool-name>",
		armsynapse.BigDataPoolResourceInfo{
			TrackedResource: armsynapse.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"key": to.StringPtr("value"),
				},
			},
			Properties: &armsynapse.BigDataPoolResourceProperties{
				AutoPause: &armsynapse.AutoPauseProperties{
					DelayInMinutes: to.Int32Ptr(15),
					Enabled:        to.BoolPtr(true),
				},
				AutoScale: &armsynapse.AutoScaleProperties{
					Enabled:      to.BoolPtr(true),
					MaxNodeCount: to.Int32Ptr(50),
					MinNodeCount: to.Int32Ptr(3),
				},
				DefaultSparkLogFolder: to.StringPtr("<default-spark-log-folder>"),
				LibraryRequirements: &armsynapse.LibraryRequirements{
					Content:  to.StringPtr("<content>"),
					Filename: to.StringPtr("<filename>"),
				},
				NodeCount:         to.Int32Ptr(4),
				NodeSize:          armsynapse.NodeSizeMedium.ToPtr(),
				NodeSizeFamily:    armsynapse.NodeSizeFamilyMemoryOptimized.ToPtr(),
				SparkEventsFolder: to.StringPtr("<spark-events-folder>"),
				SparkVersion:      to.StringPtr("<spark-version>"),
			},
		},
		&armsynapse.BigDataPoolsBeginCreateOrUpdateOptions{Force: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BigDataPoolResourceInfo.ID: %s\n", *res.ID)
}
```
