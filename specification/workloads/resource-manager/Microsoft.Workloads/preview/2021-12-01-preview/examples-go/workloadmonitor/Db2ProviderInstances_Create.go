package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/Db2ProviderInstances_Create.json
func ExampleProviderInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armworkloads.NewProviderInstancesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"myResourceGroup",
		"mySapMonitor",
		"myProviderInstance",
		armworkloads.ProviderInstance{
			Properties: &armworkloads.ProviderInstanceProperties{
				ProviderSettings: &armworkloads.DB2ProviderInstanceProperties{
					ProviderType:  to.Ptr("Db2"),
					DbName:        to.Ptr("dbName"),
					DbPassword:    to.Ptr("password"),
					DbPasswordURI: to.Ptr(""),
					DbPort:        to.Ptr("dbPort"),
					DbUsername:    to.Ptr("username"),
					Hostname:      to.Ptr("hostname"),
					SapSid:        to.Ptr("SID"),
				},
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
