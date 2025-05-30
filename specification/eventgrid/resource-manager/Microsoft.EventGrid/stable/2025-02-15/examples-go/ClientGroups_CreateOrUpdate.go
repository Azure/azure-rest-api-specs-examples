package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/ClientGroups_CreateOrUpdate.json
func ExampleClientGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClientGroupsClient().BeginCreateOrUpdate(ctx, "examplerg", "exampleNamespaceName1", "exampleClientGroupName1", armeventgrid.ClientGroup{
		Properties: &armeventgrid.ClientGroupProperties{
			Description: to.Ptr("This is a test client group"),
			Query:       to.Ptr("attributes.b IN ['a', 'b', 'c']"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClientGroup = armeventgrid.ClientGroup{
	// 	Name: to.Ptr("exampleClientGroupName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces/clientGroups"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1/clientGroups/exampleClientGroupName1"),
	// 	Properties: &armeventgrid.ClientGroupProperties{
	// 		Description: to.Ptr("This is a test client group"),
	// 		Query: to.Ptr("attributes.b IN ['a', 'b', 'c']"),
	// 	},
	// }
}
