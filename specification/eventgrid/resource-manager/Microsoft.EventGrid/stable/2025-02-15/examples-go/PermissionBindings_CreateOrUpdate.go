package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PermissionBindings_CreateOrUpdate.json
func ExamplePermissionBindingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPermissionBindingsClient().BeginCreateOrUpdate(ctx, "examplerg", "exampleNamespaceName1", "examplePermissionBindingName1", armeventgrid.PermissionBinding{
		Properties: &armeventgrid.PermissionBindingProperties{
			ClientGroupName: to.Ptr("exampleClientGroupName1"),
			Permission:      to.Ptr(armeventgrid.PermissionTypePublisher),
			TopicSpaceName:  to.Ptr("exampleTopicSpaceName1"),
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
	// res.PermissionBinding = armeventgrid.PermissionBinding{
	// 	Name: to.Ptr("examplePermissionBindingName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces/permissionBindings"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1/permissionBindings/examplePermissionBindingName1"),
	// 	Properties: &armeventgrid.PermissionBindingProperties{
	// 		ClientGroupName: to.Ptr("exampleClientGroupName1"),
	// 		Permission: to.Ptr(armeventgrid.PermissionTypePublisher),
	// 		ProvisioningState: to.Ptr(armeventgrid.PermissionBindingProvisioningStateSucceeded),
	// 		TopicSpaceName: to.Ptr("exampleTopicSpaceName1"),
	// 	},
	// }
}
