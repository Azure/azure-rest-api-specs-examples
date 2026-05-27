package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/UpdateJobAgentWithIdentity.json
func ExampleJobAgentsClient_BeginUpdate_updateAJobAgentSIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobAgentsClient().BeginUpdate(ctx, "group1", "server1", "agent1", armsql.JobAgentUpdate{
		Identity: &armsql.JobAgentIdentity{
			Type: to.Ptr(armsql.JobAgentIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armsql.JobAgentUserAssignedIdentity{
				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-umi": {},
			},
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
	// res = armsql.JobAgentsClientUpdateResponse{
	// 	JobAgent: armsql.JobAgent{
	// 		Name: to.Ptr("agent1"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/jobAgents"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1"),
	// 		Identity: &armsql.JobAgentIdentity{
	// 			Type: to.Ptr(armsql.JobAgentIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armsql.JobAgentUserAssignedIdentity{
	// 				"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-umi": &armsql.JobAgentUserAssignedIdentity{
	// 					ClientID: to.Ptr("e09c8507-0000-0000-97e2-18c5beec59dc"),
	// 					PrincipalID: to.Ptr("0c29d9b7-0ae2-4014-96ea-faf8e0cf2bc7"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("southeastasia"),
	// 		Properties: &armsql.JobAgentProperties{
	// 			DatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/databases/db1"),
	// 		},
	// 	},
	// }
}
