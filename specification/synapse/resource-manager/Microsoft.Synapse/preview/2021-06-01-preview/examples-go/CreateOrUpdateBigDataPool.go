package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/CreateOrUpdateBigDataPool.json
func ExampleBigDataPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBigDataPoolsClient().BeginCreateOrUpdate(ctx, "ExampleResourceGroup", "ExampleWorkspace", "ExamplePool", armsynapse.BigDataPoolResourceInfo{
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
			IsAutotuneEnabled:     to.Ptr(false),
			LibraryRequirements: &armsynapse.LibraryRequirements{
				Content:  to.Ptr(""),
				Filename: to.Ptr("requirements.txt"),
			},
			NodeCount:         to.Ptr[int32](4),
			NodeSize:          to.Ptr(armsynapse.NodeSizeMedium),
			NodeSizeFamily:    to.Ptr(armsynapse.NodeSizeFamilyMemoryOptimized),
			SparkEventsFolder: to.Ptr("/events"),
			SparkVersion:      to.Ptr("3.3"),
		},
	}, &armsynapse.BigDataPoolsClientBeginCreateOrUpdateOptions{Force: nil})
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
	// res.BigDataPoolResourceInfo = armsynapse.BigDataPoolResourceInfo{
	// 	Name: to.Ptr("ExamplePool"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/bigDataPools"),
	// 	ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/bigDataPools/ExamplePool"),
	// 	Location: to.Ptr("West US 2"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armsynapse.BigDataPoolResourceProperties{
	// 		AutoPause: &armsynapse.AutoPauseProperties{
	// 			DelayInMinutes: to.Ptr[int32](15),
	// 			Enabled: to.Ptr(true),
	// 		},
	// 		AutoScale: &armsynapse.AutoScaleProperties{
	// 			Enabled: to.Ptr(true),
	// 			MaxNodeCount: to.Ptr[int32](50),
	// 			MinNodeCount: to.Ptr[int32](3),
	// 		},
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t}()),
	// 		DefaultSparkLogFolder: to.Ptr("/logs"),
	// 		IsAutotuneEnabled: to.Ptr(false),
	// 		LastSucceededTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T10:00:00.000Z"); return t}()),
	// 		LibraryRequirements: &armsynapse.LibraryRequirements{
	// 			Content: to.Ptr(""),
	// 			Filename: to.Ptr("requirements.txt"),
	// 			Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t}()),
	// 		},
	// 		NodeCount: to.Ptr[int32](4),
	// 		NodeSize: to.Ptr(armsynapse.NodeSizeMedium),
	// 		NodeSizeFamily: to.Ptr(armsynapse.NodeSizeFamilyMemoryOptimized),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SparkEventsFolder: to.Ptr("/events"),
	// 		SparkVersion: to.Ptr("3.3"),
	// 	},
	// }
}
