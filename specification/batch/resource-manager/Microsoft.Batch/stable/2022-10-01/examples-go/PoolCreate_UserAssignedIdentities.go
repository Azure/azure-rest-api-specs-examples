package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_UserAssignedIdentities.json
func ExamplePoolClient_Create_createPoolUserAssignedIdentities() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Identity: &armbatch.PoolIdentity{
			Type: to.Ptr(armbatch.PoolIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armbatch.UserAssignedIdentities{
				"/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
				"/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
			},
		},
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					ImageReference: &armbatch.ImageReference{
						Offer:     to.Ptr("UbuntuServer"),
						Publisher: to.Ptr("Canonical"),
						SKU:       to.Ptr("18.04-LTS"),
						Version:   to.Ptr("latest"),
					},
					NodeAgentSKUID: to.Ptr("batch.node.ubuntu 18.04"),
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				AutoScale: &armbatch.AutoScaleSettings{
					EvaluationInterval: to.Ptr("PT5M"),
					Formula:            to.Ptr("$TargetDedicatedNodes=1"),
				},
			},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
