package armoep_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2022-04-04-preview/examples/EnergyServices_RemovePartition_MaximumSet_Gen.json
func ExampleEnergyServicesClient_BeginRemovePartition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoep.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEnergyServicesClient().BeginRemovePartition(ctx, "rgoep", "aaaaaaa", &armoep.EnergyServicesClientBeginRemovePartitionOptions{Body: &armoep.DataPartitionAddOrRemoveRequest{
		Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
