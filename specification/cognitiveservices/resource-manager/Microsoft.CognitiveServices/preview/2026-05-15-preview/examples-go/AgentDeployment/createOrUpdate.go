package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/AgentDeployment/createOrUpdate.json
func ExampleAgentDeploymentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentDeploymentsClient().BeginCreateOrUpdate(ctx, "test-rg", "my-cognitive-services-account", "my-project", "agent-app-1", "deployment-1", armcognitiveservices.AgentDeployment{
		Properties: &armcognitiveservices.ManagedAgentDeployment{
			Agents: []*armcognitiveservices.VersionedAgentReference{
				{
					AgentID:      to.Ptr("agent-123"),
					AgentName:    to.Ptr("support-agent"),
					AgentVersion: to.Ptr("1.0.0"),
				},
			},
			DeploymentType: to.Ptr(armcognitiveservices.AgentDeploymentTypeManaged),
			DisplayName:    to.Ptr("Production Deployment"),
			Protocols: []*armcognitiveservices.AgentProtocolVersion{
				{
					Version:  to.Ptr("1.0"),
					Protocol: to.Ptr(armcognitiveservices.AgentProtocolAgent),
				},
			},
			State: to.Ptr(armcognitiveservices.AgentDeploymentStateStarting),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.AgentDeploymentsClientCreateOrUpdateResponse{
	// 	AgentDeployment: armcognitiveservices.AgentDeployment{
	// 		Name: to.Ptr("deployment-1"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/applications/agentDeployments"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveServices/accounts/my-cognitive-services-account/projects/my-project/applications/agent-app-1/agentDeployments/deployment-1"),
	// 		Properties: &armcognitiveservices.ManagedAgentDeployment{
	// 			Agents: []*armcognitiveservices.VersionedAgentReference{
	// 				{
	// 					AgentID: to.Ptr("agent-123"),
	// 					AgentName: to.Ptr("support-agent"),
	// 					AgentVersion: to.Ptr("1.0.0"),
	// 				},
	// 			},
	// 			DeploymentID: to.Ptr("550e8400-e29b-41d4-a716-446655440001"),
	// 			DeploymentType: to.Ptr(armcognitiveservices.AgentDeploymentTypeManaged),
	// 			DisplayName: to.Ptr("Production Deployment"),
	// 			Protocols: []*armcognitiveservices.AgentProtocolVersion{
	// 				{
	// 					Version: to.Ptr("1.0"),
	// 					Protocol: to.Ptr(armcognitiveservices.AgentProtocolAgent),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcognitiveservices.AgentDeploymentProvisioningStateSucceeded),
	// 			State: to.Ptr(armcognitiveservices.AgentDeploymentStateRunning),
	// 		},
	// 		SystemData: &armcognitiveservices.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-29T12:34:56.999Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-29T12:34:56.999Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
