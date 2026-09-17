package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2026-07-01/Job_Stop_Multiple.json
func ExampleJobsClient_BeginStopMultipleExecutions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobsClient().BeginStopMultipleExecutions(ctx, "rg", "testcontainerAppsJob0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.JobsClientStopMultipleExecutionsResponse{
	// 	ContainerAppJobExecutions: armappcontainers.ContainerAppJobExecutions{
	// 		Value: []*armappcontainers.JobExecution{
	// 			{
	// 				Name: to.Ptr("jobExecution-27944453"),
	// 				Properties: &armappcontainers.JobExecutionProperties{
	// 					EndTime: to.Ptr(time.Date(2023, time.February, 13, 20, 47, 30, 0, time.UTC)),
	// 					StartTime: to.Ptr(time.Date(2023, time.February, 13, 20, 37, 30, 0, time.UTC)),
	// 					Status: to.Ptr(armappcontainers.JobExecutionRunningStateRunning),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("jobExecution-27944452"),
	// 				Properties: &armappcontainers.JobExecutionProperties{
	// 					EndTime: to.Ptr(time.Date(2023, time.February, 13, 21, 47, 30, 0, time.UTC)),
	// 					StartTime: to.Ptr(time.Date(2023, time.February, 13, 21, 37, 30, 0, time.UTC)),
	// 					Status: to.Ptr(armappcontainers.JobExecutionRunningStateRunning),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("jobExecution-27944453"),
	// 				Properties: &armappcontainers.JobExecutionProperties{
	// 					EndTime: to.Ptr(time.Date(2023, time.February, 13, 22, 47, 30, 0, time.UTC)),
	// 					StartTime: to.Ptr(time.Date(2023, time.February, 13, 22, 37, 30, 0, time.UTC)),
	// 					Status: to.Ptr(armappcontainers.JobExecutionRunningStateRunning),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
