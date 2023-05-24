package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a7af0df86022e5e6cc6e8f40ca1981c4557a4bc/specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/Job_Stop_Multiple.json
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
	poller, err := clientFactory.NewJobsClient().BeginStopMultipleExecutions(ctx, "rg", "testcontainerAppsJob0", armappcontainers.JobExecutionNamesCollection{
		Value: []*armappcontainers.JobExecutionBase{
			{
				Name: to.Ptr("jobExecution-27944453"),
			},
			{
				Name: to.Ptr("jobExecution-27944452"),
			},
			{
				Name: to.Ptr("jobExecution-27944451"),
			}},
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
	// res.ContainerAppJobExecutions = armappcontainers.ContainerAppJobExecutions{
	// 	Value: []*armappcontainers.JobExecution{
	// 		{
	// 			Name: to.Ptr("jobExecution-27944453"),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:47:30+00:00"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:37:30+00:00"); return t}()),
	// 			Status: to.Ptr(armappcontainers.JobExecutionRunningStateStopped),
	// 		},
	// 		{
	// 			Name: to.Ptr("jobExecution-27944452"),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:47:30+00:00"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:37:30+00:00"); return t}()),
	// 			Status: to.Ptr(armappcontainers.JobExecutionRunningStateStopped),
	// 		},
	// 		{
	// 			Name: to.Ptr("jobExecution-27944453"),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:47:30+00:00"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:37:30+00:00"); return t}()),
	// 			Status: to.Ptr(armappcontainers.JobExecutionRunningStateFailed),
	// 	}},
	// }
}
