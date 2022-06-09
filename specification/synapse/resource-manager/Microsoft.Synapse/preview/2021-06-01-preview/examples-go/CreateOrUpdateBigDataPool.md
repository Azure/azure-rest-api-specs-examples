```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/CreateOrUpdateBigDataPool.json
func ExampleBigDataPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewBigDataPoolsClient("01234567-89ab-4def-0123-456789abcdef", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"ExampleResourceGroup",
		"ExampleWorkspace",
		"ExamplePool",
		armsynapse.BigDataPoolResourceInfo{
			Location: to.Ptr("West US 2"),
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
				DefaultSparkLogFolder: to.Ptr("/logs"),
				LibraryRequirements: &armsynapse.LibraryRequirements{
					Content:  to.Ptr(""),
					Filename: to.Ptr("requirements.txt"),
				},
				NodeCount:         to.Ptr[int32](4),
				NodeSize:          to.Ptr(armsynapse.NodeSizeMedium),
				NodeSizeFamily:    to.Ptr(armsynapse.NodeSizeFamilyMemoryOptimized),
				SparkEventsFolder: to.Ptr("/events"),
				SparkVersion:      to.Ptr("2.4"),
			},
		},
		&armsynapse.BigDataPoolsClientBeginCreateOrUpdateOptions{Force: nil})
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.5.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.
