package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a7af0df86022e5e6cc6e8f40ca1981c4557a4bc/specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/Job_Executions_Get.json
func ExampleJobsExecutionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsExecutionsClient().NewListPager("rg", "testcontainerAppsJob0", &armappcontainers.JobsExecutionsClientListOptions{Filter: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ContainerAppJobExecutions = armappcontainers.ContainerAppJobExecutions{
		// 	Value: []*armappcontainers.JobExecution{
		// 		{
		// 			Name: to.Ptr("testcontainerAppJob-27944454"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:47:30+00:00"); return t}()),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:37:30+00:00"); return t}()),
		// 			Status: to.Ptr(armappcontainers.JobExecutionRunningStateRunning),
		// 			Template: &armappcontainers.JobExecutionTemplate{
		// 				Containers: []*armappcontainers.JobExecutionContainer{
		// 					{
		// 						Name: to.Ptr("testcontainerAppsJob0"),
		// 						Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
		// 						Resources: &armappcontainers.ContainerResources{
		// 							CPU: to.Ptr[float64](0.2),
		// 							Memory: to.Ptr("100Mi"),
		// 						},
		// 				}},
		// 				InitContainers: []*armappcontainers.JobExecutionContainer{
		// 					{
		// 						Name: to.Ptr("testinitcontainerAppsJob0"),
		// 						Args: []*string{
		// 							to.Ptr("-c"),
		// 							to.Ptr("while true; do echo hello; sleep 10;done")},
		// 							Command: []*string{
		// 								to.Ptr("/bin/sh")},
		// 								Image: to.Ptr("repo/testcontainerAppsJob0:v4"),
		// 								Resources: &armappcontainers.ContainerResources{
		// 									CPU: to.Ptr[float64](0.2),
		// 									Memory: to.Ptr("100Mi"),
		// 								},
		// 						}},
		// 					},
		// 			}},
		// 		}
	}
}
