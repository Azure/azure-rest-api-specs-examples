package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_PerformMaintenance_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginPerformMaintenance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginPerformMaintenance(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		&armcompute.VirtualMachineScaleSetsClientBeginPerformMaintenanceOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
			InstanceIDs: []*string{
				to.Ptr("aaaaaaaaaaaaaaaaa")},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
