package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c0a12a75b702054cf1e7fcd8c014d0fc116dea6e/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/BareMetalMachines_PowerOff.json
func ExampleBareMetalMachinesClient_BeginPowerOff() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBareMetalMachinesClient().BeginPowerOff(ctx, "resourceGroupName", "bareMetalMachineName", &armnetworkcloud.BareMetalMachinesClientBeginPowerOffOptions{BareMetalMachinePowerOffParameters: &armnetworkcloud.BareMetalMachinePowerOffParameters{
		SkipShutdown: to.Ptr(armnetworkcloud.BareMetalMachineSkipShutdownTrue),
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
