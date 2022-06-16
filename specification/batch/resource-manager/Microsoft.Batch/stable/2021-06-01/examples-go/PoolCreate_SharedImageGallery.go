package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PoolCreate_SharedImageGallery.json
func ExamplePoolClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPoolClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pool-name>",
		armbatch.Pool{
			Properties: &armbatch.PoolProperties{
				DeploymentConfiguration: &armbatch.DeploymentConfiguration{
					VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
						ImageReference: &armbatch.ImageReference{
							ID: to.StringPtr("<id>"),
						},
						NodeAgentSKUID: to.StringPtr("<node-agent-skuid>"),
					},
				},
				VMSize: to.StringPtr("<vmsize>"),
			},
		},
		&armbatch.PoolClientCreateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PoolClientCreateResult)
}
