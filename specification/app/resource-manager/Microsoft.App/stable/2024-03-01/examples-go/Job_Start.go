package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Job_Start.json
func ExampleJobsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobsClient().BeginStart(ctx, "rg", "testcontainerappsjob0", &armappcontainers.JobsClientBeginStartOptions{Template: &armappcontainers.JobExecutionTemplate{
		Containers: []*armappcontainers.JobExecutionContainer{
			{
				Name:  to.Ptr("testcontainerappsjob0"),
				Image: to.Ptr("repo/testcontainerappsjob0:v4"),
				Resources: &armappcontainers.ContainerResources{
					CPU:    to.Ptr[float64](0.5),
					Memory: to.Ptr("1Gi"),
				},
			}},
		InitContainers: []*armappcontainers.JobExecutionContainer{
			{
				Name: to.Ptr("testinitcontainerAppsJob0"),
				Args: []*string{
					to.Ptr("-c"),
					to.Ptr("while true; do echo hello; sleep 10;done")},
				Command: []*string{
					to.Ptr("/bin/sh")},
				Image: to.Ptr("repo/testcontainerappsjob0:v4"),
				Resources: &armappcontainers.ContainerResources{
					CPU:    to.Ptr[float64](0.5),
					Memory: to.Ptr("1Gi"),
				},
			}},
	},
	})
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
	// res.JobExecutionBase = armappcontainers.JobExecutionBase{
	// 	Name: to.Ptr("testcontainerappsjob0-pjxhsye"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/jobs/{containerAppsJobName}/executions/{jobExecutionName}"),
	// }
}
