package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Jobs_Get-GET-example-101.json
func ExampleJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().Get(ctx, "DataTransformation", "jobdeffromtestcode1", "99ef93fe-36be-43e4-bebf-de6746730601", "ResourceGroupForSDKTest", "TestAzureSDKOperations", &armhybriddatamanager.JobsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armhybriddatamanager.Job{
	// 	Name: to.Ptr("99ef93fe-36be-43e4-bebf-de6746730601"),
	// 	Type: to.Ptr("Microsoft.HybridData/dataManagers/jobs"),
	// 	ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataServices/DataTransformation/jobDefinitions/jobdeffromtestcode1/jobs/99ef93fe-36be-43e4-bebf-de6746730601"),
	// 	Properties: &armhybriddatamanager.JobProperties{
	// 		BytesProcessed: to.Ptr[int64](0),
	// 		IsCancellable: to.Ptr(armhybriddatamanager.IsJobCancellableCancellable),
	// 		ItemsProcessed: to.Ptr[int64](0),
	// 		TotalBytesToProcess: to.Ptr[int64](0),
	// 		TotalItemsToProcess: to.Ptr[int64](0),
	// 	},
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-05T04:42:59.277Z"); return t}()),
	// 	Status: to.Ptr(armhybriddatamanager.JobStatusInProgress),
	// }
}
