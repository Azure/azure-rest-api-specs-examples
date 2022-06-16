package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VmssPublicIpGet.json
func ExamplePublicIPAddressesClient_GetVirtualMachineScaleSetPublicIPAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewPublicIPAddressesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetVirtualMachineScaleSetPublicIPAddress(ctx,
		"<resource-group-name>",
		"<virtual-machine-scale-set-name>",
		"<virtualmachine-index>",
		"<network-interface-name>",
		"<ip-configuration-name>",
		"<public-ipaddress-name>",
		&armnetwork.PublicIPAddressesClientGetVirtualMachineScaleSetPublicIPAddressOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
