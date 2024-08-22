package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Job_Stop_Multiple.json
func ExampleJobsClient_BeginStopMultipleExecutions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobsClient().BeginStopMultipleExecutions(ctx, "rg", "testcontainerappsjob0", nil)
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
	// res.ContainerAppJobExecutions = armappcontainers.ContainerAppJobExecutions{
	// 	Value: []*armappcontainers.JobExecution{
	// 		{
	// 			Name: to.Ptr("jobExecution-27944453"),
	// 			Properties: &armappcontainers.JobExecutionProperties{
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:47:30.000Z"); return t}()),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:37:30.000Z"); return t}()),
	// 				Status: to.Ptr(armappcontainers.JobExecutionRunningStateRunning),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("jobExecution-27944452"),
	// 			Properties: &armappcontainers.JobExecutionProperties{
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T21:47:30.000Z"); return t}()),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T21:37:30.000Z"); return t}()),
	// 				Status: to.Ptr(armappcontainers.JobExecutionRunningStateRunning),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("jobExecution-27944453"),
	// 			Properties: &armappcontainers.JobExecutionProperties{
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T22:47:30.000Z"); return t}()),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T22:37:30.000Z"); return t}()),
	// 				Status: to.Ptr(armappcontainers.JobExecutionRunningStateRunning),
	// 			},
	// 	}},
	// }
}
