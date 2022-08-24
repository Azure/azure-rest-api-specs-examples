package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-07-02-preview/examples/AgentPoolsCreate_Update.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createUpdateAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count: to.Ptr[int32](3),
			Mode:  to.Ptr(armcontainerservice.AgentPoolModeUser),
			NodeLabels: map[string]*string{
				"key1": to.Ptr("val1"),
			},
			NodeTaints: []*string{
				to.Ptr("Key1=Value1:NoSchedule")},
			OrchestratorVersion:    to.Ptr(""),
			OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
			ScaleSetEvictionPolicy: to.Ptr(armcontainerservice.ScaleSetEvictionPolicyDelete),
			ScaleSetPriority:       to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
			Tags: map[string]*string{
				"name1": to.Ptr("val1"),
			},
			VMSize: to.Ptr("Standard_DS1_v2"),
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
