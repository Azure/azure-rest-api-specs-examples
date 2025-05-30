package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Job_Execution_Get.json
func ExampleContainerAppsAPIClient_JobExecution() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsAPIClient().JobExecution(ctx, "rg", "testcontainerappsjob0", "jobExecution1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobExecution = armappcontainers.JobExecution{
	// 	Name: to.Ptr("jobExecution1"),
	// 	Properties: &armappcontainers.JobExecutionProperties{
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:47:30.000Z"); return t}()),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-13T20:37:30.000Z"); return t}()),
	// 		Status: to.Ptr(armappcontainers.JobExecutionRunningStateRunning),
	// 		Template: &armappcontainers.JobExecutionTemplate{
	// 			Containers: []*armappcontainers.JobExecutionContainer{
	// 				{
	// 					Name: to.Ptr("testcontainerappsjob0"),
	// 					Image: to.Ptr("repo/testcontainerappsjob0:v4"),
	// 					Resources: &armappcontainers.ContainerResources{
	// 						CPU: to.Ptr[float64](0.5),
	// 						Memory: to.Ptr("1Gi"),
	// 					},
	// 			}},
	// 			InitContainers: []*armappcontainers.JobExecutionContainer{
	// 				{
	// 					Name: to.Ptr("testinitcontainerAppsJob0"),
	// 					Args: []*string{
	// 						to.Ptr("-c"),
	// 						to.Ptr("while true; do echo hello; sleep 10;done")},
	// 						Command: []*string{
	// 							to.Ptr("/bin/sh")},
	// 							Image: to.Ptr("repo/testcontainerappsjob0:v4"),
	// 							Resources: &armappcontainers.ContainerResources{
	// 								CPU: to.Ptr[float64](0.5),
	// 								Memory: to.Ptr("1Gi"),
	// 							},
	// 					}},
	// 				},
	// 			},
	// 		}
}
