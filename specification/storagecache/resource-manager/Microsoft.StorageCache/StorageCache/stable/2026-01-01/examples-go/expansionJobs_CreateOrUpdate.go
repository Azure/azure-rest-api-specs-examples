package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3ca2244e6becf79bdd1198e9d66f0e8e8d3a0977/specification/storagecache/resource-manager/Microsoft.StorageCache/StorageCache/stable/2026-01-01/examples/expansionJobs_CreateOrUpdate.json
func ExampleExpansionJobsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpansionJobsClient().BeginCreateOrUpdate(ctx, "scgroup", "fs1", "expansionjob1", armstoragecache.ExpansionJob{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"Dept": to.Ptr("ContosoAds"),
		},
		Properties: &armstoragecache.ExpansionJobProperties{
			NewStorageCapacityTiB: to.Ptr[float32](16),
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
	// res.ExpansionJob = armstoragecache.ExpansionJob{
	// 	Name: to.Ptr("expansionjob1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/expansionJob"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/expansionJobs/expansionjob1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"Dept": to.Ptr("ContosoAds"),
	// 	},
	// 	Properties: &armstoragecache.ExpansionJobProperties{
	// 		NewStorageCapacityTiB: to.Ptr[float32](16),
	// 		ProvisioningState: to.Ptr(armstoragecache.ExpansionJobPropertiesProvisioningStateSucceeded),
	// 		Status: &armstoragecache.ExpansionJobPropertiesStatus{
	// 			PercentComplete: to.Ptr[float32](85),
	// 			StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T17:25:43.511Z"); return t}()),
	// 			State: to.Ptr(armstoragecache.ExpansionJobStatusTypeInProgress),
	// 			StatusMessage: to.Ptr("Expansion is in progress"),
	// 		},
	// 	},
	// }
}
