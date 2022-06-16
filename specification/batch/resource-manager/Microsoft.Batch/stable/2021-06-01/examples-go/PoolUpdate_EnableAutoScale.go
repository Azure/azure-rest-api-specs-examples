package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PoolUpdate_EnableAutoScale.json
func ExamplePoolClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPoolClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pool-name>",
		armbatch.Pool{
			Properties: &armbatch.PoolProperties{
				ScaleSettings: &armbatch.ScaleSettings{
					AutoScale: &armbatch.AutoScaleSettings{
						Formula: to.StringPtr("<formula>"),
					},
				},
			},
		},
		&armbatch.PoolClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PoolClientUpdateResult)
}
