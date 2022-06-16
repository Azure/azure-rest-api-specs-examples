package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_CreateOrUpdate.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewVirtualMachinesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resourceGroupName",
		"{labName}",
		"{vmName}",
		armdevtestlabs.LabVirtualMachine{
			Location: to.Ptr("{location}"),
			Tags: map[string]*string{
				"tagName1": to.Ptr("tagValue1"),
			},
			Properties: &armdevtestlabs.LabVirtualMachineProperties{
				AllowClaim:              to.Ptr(true),
				DisallowPublicIPAddress: to.Ptr(true),
				GalleryImageReference: &armdevtestlabs.GalleryImageReference{
					Offer:     to.Ptr("UbuntuServer"),
					OSType:    to.Ptr("Linux"),
					Publisher: to.Ptr("Canonical"),
					SKU:       to.Ptr("16.04-LTS"),
					Version:   to.Ptr("Latest"),
				},
				LabSubnetName:       to.Ptr("{virtualNetworkName}Subnet"),
				LabVirtualNetworkID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}"),
				Password:            to.Ptr("{userPassword}"),
				Size:                to.Ptr("Standard_A2_v2"),
				StorageType:         to.Ptr("Standard"),
				UserName:            to.Ptr("{userName}"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
