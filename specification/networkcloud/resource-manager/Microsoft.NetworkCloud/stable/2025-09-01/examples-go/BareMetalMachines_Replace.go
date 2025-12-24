package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/46c51b03d99b113ecc3b38883e3cb2d395fe94a4/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/BareMetalMachines_Replace.json
func ExampleBareMetalMachinesClient_BeginReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBareMetalMachinesClient().BeginReplace(ctx, "resourceGroupName", "bareMetalMachineName", &armnetworkcloud.BareMetalMachinesClientBeginReplaceOptions{BareMetalMachineReplaceParameters: &armnetworkcloud.BareMetalMachineReplaceParameters{
		BmcCredentials: &armnetworkcloud.AdministrativeCredentials{
			Password: to.Ptr("https://keyvaultname.vault.azure.net/secrets/secretName"),
			Username: to.Ptr("bmcuser"),
		},
		BmcMacAddress:  to.Ptr("00:00:4f:00:57:ad"),
		BootMacAddress: to.Ptr("00:00:4e:00:58:af"),
		MachineName:    to.Ptr("name"),
		SafeguardMode:  to.Ptr(armnetworkcloud.BareMetalMachineReplaceSafeguardModeAll),
		SerialNumber:   to.Ptr("BM1219XXX"),
		StoragePolicy:  to.Ptr(armnetworkcloud.BareMetalMachineReplaceStoragePolicyDiscardAll),
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
