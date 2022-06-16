package armdelegatednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/putOrchestrator.json
func ExampleOrchestratorInstanceServiceClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdelegatednetwork.NewOrchestratorInstanceServiceClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"TestRG",
		"testk8s1",
		armdelegatednetwork.Orchestrator{
			Identity: &armdelegatednetwork.OrchestratorIdentity{
				Type: to.Ptr(armdelegatednetwork.ResourceIdentityTypeSystemAssigned),
			},
			Kind:     to.Ptr(armdelegatednetwork.OrchestratorKindKubernetes),
			Location: to.Ptr("West US"),
			Properties: &armdelegatednetwork.OrchestratorResourceProperties{
				APIServerEndpoint: to.Ptr("https://testk8s.cloudapp.net"),
				ClusterRootCA:     to.Ptr("ddsadsad344mfdsfdl"),
				ControllerDetails: &armdelegatednetwork.ControllerDetails{
					ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/controller/testcontroller"),
				},
				OrchestratorAppID:     to.Ptr("546192d7-503f-477a-9cfe-4efc3ee2b6e1"),
				OrchestratorTenantID:  to.Ptr("da6192d7-503f-477a-9cfe-4efc3ee2b6c3"),
				PrivateLinkResourceID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.Network/privateLinkServices/plresource1"),
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
