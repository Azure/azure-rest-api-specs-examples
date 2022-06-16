package armcontainerservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-03-01/examples/AgentPoolsCreate_Snapshot.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<agent-pool-name>",
		armcontainerservice.AgentPool{
			Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
				Count: to.Ptr[int32](3),
				CreationData: &armcontainerservice.CreationData{
					SourceResourceID: to.Ptr("<source-resource-id>"),
				},
				EnableFIPS:          to.Ptr(true),
				OrchestratorVersion: to.Ptr("<orchestrator-version>"),
				OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
				VMSize:              to.Ptr("<vmsize>"),
			},
		},
		&armcontainerservice.AgentPoolsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
