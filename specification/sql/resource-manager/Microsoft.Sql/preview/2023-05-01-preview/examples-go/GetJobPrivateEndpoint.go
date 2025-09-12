package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/GetJobPrivateEndpoint.json
func ExampleJobPrivateEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobPrivateEndpointsClient().Get(ctx, "group1", "server1", "agent1", "endpoint1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobPrivateEndpoint = armsql.JobPrivateEndpoint{
	// 	Name: to.Ptr("endpoint1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/jobAgents/privateEndpoints"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/privateEndpoints/endpoint1"),
	// 	Properties: &armsql.JobPrivateEndpointProperties{
	// 		PrivateEndpointID: to.Ptr("EJ_47e33188-85ff-4705-8d78-444444444444_endpoint1"),
	// 		TargetServerAzureResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/targetserver1"),
	// 	},
	// }
}
