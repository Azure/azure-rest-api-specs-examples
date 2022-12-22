package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationvCenters_Create.json
func ExampleReplicationvCentersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationvCentersClient("MadhaviVault", "MadhaviVRG", "7c943c1b-5122-4097-90c8-861411bdd574", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "MadhaviFabric", "esx-78", armrecoveryservicessiterecovery.AddVCenterRequest{
		Properties: &armrecoveryservicessiterecovery.AddVCenterRequestProperties{
			FriendlyName:    to.Ptr("esx-78"),
			IPAddress:       to.Ptr("inmtest78"),
			Port:            to.Ptr("443"),
			ProcessServerID: to.Ptr("5A720CAB-39CB-F445-BD1662B0B33164B5"),
			RunAsAccountID:  to.Ptr("2"),
		},
	}, nil)
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
