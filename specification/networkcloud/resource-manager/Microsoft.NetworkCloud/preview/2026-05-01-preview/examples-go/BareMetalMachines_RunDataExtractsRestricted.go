package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: 2026-05-01-preview/BareMetalMachines_RunDataExtractsRestricted.json
func ExampleBareMetalMachinesClient_BeginRunDataExtractsRestricted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("123e4567-e89b-12d3-a456-426655440000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBareMetalMachinesClient().BeginRunDataExtractsRestricted(ctx, "resourceGroupName", "bareMetalMachineName", armnetworkcloud.BareMetalMachineRunDataExtractsParameters{
		Commands: []*armnetworkcloud.BareMetalMachineCommandSpecification{
			{
				Arguments: []*string{
					to.Ptr("--min-severity=8"),
				},
				Command: to.Ptr("cluster-cve-report"),
			},
		},
		LimitTimeSeconds: to.Ptr[int64](60),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
