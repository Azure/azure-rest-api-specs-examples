package armmanagednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ScopeAssignment/ScopeAssignmentsPut.json
func ExampleScopeAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScopeAssignmentsClient().CreateOrUpdate(ctx, "subscriptions/subscriptionC", "subscriptionCAssignment", armmanagednetwork.ScopeAssignment{
		Properties: &armmanagednetwork.ScopeAssignmentProperties{
			AssignedManagedNetwork: to.Ptr("/subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.ManagedNetwork/managedNetworks/myManagedNetwork"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScopeAssignment = armmanagednetwork.ScopeAssignment{
	// 	Name: to.Ptr("subscriptionCAssignment"),
	// 	Type: to.Ptr("Microsoft.ManagedNetwork/scopeAssignment"),
	// 	ID: to.Ptr("/subscriptions/subscriptionC/providers/Microsoft.ManagedNetwork/scopeAssignments/subscriptionCAssignment"),
	// 	Properties: &armmanagednetwork.ScopeAssignmentProperties{
	// 		Etag: to.Ptr("sadf-asdf-asdf-asdf"),
	// 		ProvisioningState: to.Ptr(armmanagednetwork.ProvisioningStateSucceeded),
	// 		AssignedManagedNetwork: to.Ptr("/subscriptions/subscriptionA/resourceGroups/myResourceGroup/providers/Microsoft.ManagedNetwork/managedNetworks/myManagedNetwork"),
	// 	},
	// }
}
