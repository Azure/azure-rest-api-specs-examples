package armdelegatednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/getOrchestrator.json
func ExampleOrchestratorInstanceServiceClient_GetDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdelegatednetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrchestratorInstanceServiceClient().GetDetails(ctx, "TestRG", "testk8s1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Orchestrator = armdelegatednetwork.Orchestrator{
	// 	Name: to.Ptr("testk8s1"),
	// 	Type: to.Ptr("Microsoft.DelegatedNetwork/orchestrators"),
	// 	ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/orchestrators/testk8s1"),
	// 	Identity: &armdelegatednetwork.OrchestratorIdentity{
	// 		Type: to.Ptr(armdelegatednetwork.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("7a2192d7-503f-477a-9cfe-4efc3ee2bd60"),
	// 		TenantID: to.Ptr("3e2192d7-503f-477a-9cfe-4efc3ee2bd60"),
	// 	},
	// 	Kind: to.Ptr(armdelegatednetwork.OrchestratorKindKubernetes),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armdelegatednetwork.OrchestratorResourceProperties{
	// 		APIServerEndpoint: to.Ptr("https://testk8s.cloudapp.net"),
	// 		ClusterRootCA: to.Ptr("ddsadsad344mfdsfdl"),
	// 		ControllerDetails: &armdelegatednetwork.ControllerDetails{
	// 			ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/controller/testcontroller"),
	// 		},
	// 		OrchestratorAppID: to.Ptr("546192d7-503f-477a-9cfe-4efc3ee2b6e1"),
	// 		OrchestratorTenantID: to.Ptr("da6192d7-503f-477a-9cfe-4efc3ee2b6c3"),
	// 		PrivateLinkResourceID: to.Ptr("/subscriptions/d21192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/plrg/Microsoft.Network/PrivateLink/pl1"),
	// 		ProvisioningState: to.Ptr(armdelegatednetwork.OrchestratorInstanceStateSucceeded),
	// 		ResourceGUID: to.Ptr("1b2192d7-503f-477a-9cfe-4efc3ee2bd60"),
	// 	},
	// }
}
