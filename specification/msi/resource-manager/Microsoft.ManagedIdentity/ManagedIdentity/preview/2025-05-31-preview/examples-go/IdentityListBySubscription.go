package armmsi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

// Generated from example definition: 2025-05-31-preview/IdentityListBySubscription.json
func ExampleUserAssignedIdentitiesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmsi.NewClientFactory("12345678-1234-5678-9012-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUserAssignedIdentitiesClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armmsi.UserAssignedIdentitiesClientListBySubscriptionResponse{
		// 	UserAssignedIdentitiesListResult: armmsi.UserAssignedIdentitiesListResult{
		// 		Value: []*armmsi.Identity{
		// 			{
		// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName"),
		// 				Location: to.Ptr("eastus"),
		// 				Name: to.Ptr("identityName"),
		// 				Properties: &armmsi.UserAssignedIdentityProperties{
		// 					ClientID: to.Ptr("4024ab25-56a8-4370-aea6-6389221caf29"),
		// 					PrincipalID: to.Ptr("25cc773c-7f05-40fc-a104-32d2300754ad"),
		// 					TenantID: to.Ptr("b6c948ef-f6b5-4384-8354-da3a15eca969"),
		// 					IsolationScope: to.Ptr(armmsi.IsolationScopeRegional),
		// 					AssignmentRestrictions: &armmsi.AssignmentRestrictions{
		// 						Providers: []*string{
		// 							to.Ptr("Microsoft.Compute"),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 					"key2": to.Ptr("value2"),
		// 				},
		// 				Type: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://serviceRoot/subscriptions/subId/providers/Microsoft.ManagedIdentity/userAssignedIdentities?api-version=2025-05-31-preview&$skiptoken=X'12345'"),
		// 	},
		// }
	}
}
