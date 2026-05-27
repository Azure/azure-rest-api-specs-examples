package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/CreateOrUpdateJobAgent.json
func ExampleJobAgentsClient_BeginCreateOrUpdate_createOrUpdateAJobAgent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobAgentsClient().BeginCreateOrUpdate(ctx, "group1", "server1", "agent1", armsql.JobAgent{
		Location: to.Ptr("southeastasia"),
		Properties: &armsql.JobAgentProperties{
			DatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/databases/db1"),
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
	// res = armsql.JobAgentsClientCreateOrUpdateResponse{
	// 	JobAgent: armsql.JobAgent{
	// 		Name: to.Ptr("agent1"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/jobAgents"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1"),
	// 		Location: to.Ptr("southeastasia"),
	// 		Properties: &armsql.JobAgentProperties{
	// 			DatabaseID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/databases/db1"),
	// 		},
	// 		SKU: &armsql.SKU{
	// 			Name: to.Ptr("JA100"),
	// 			Capacity: to.Ptr[int32](100),
	// 		},
	// 	},
	// }
}
