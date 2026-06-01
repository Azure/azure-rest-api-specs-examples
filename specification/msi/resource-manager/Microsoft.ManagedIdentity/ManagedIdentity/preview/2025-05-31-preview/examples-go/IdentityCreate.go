package armmsi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

// Generated from example definition: 2025-05-31-preview/IdentityCreate.json
func ExampleUserAssignedIdentitiesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmsi.NewClientFactory("12345678-1234-5678-9012-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserAssignedIdentitiesClient().CreateOrUpdate(ctx, "rgName", "resourceName", armmsi.Identity{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmsi.UserAssignedIdentitiesClientCreateOrUpdateResponse{
	// 	Identity: armmsi.Identity{
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName"),
	// 		Location: to.Ptr("eastus"),
	// 		Name: to.Ptr("identityName"),
	// 		Properties: &armmsi.UserAssignedIdentityProperties{
	// 			ClientID: to.Ptr("4024ab25-56a8-4370-aea6-6389221caf29"),
	// 			PrincipalID: to.Ptr("25cc773c-7f05-40fc-a104-32d2300754ad"),
	// 			TenantID: to.Ptr("b6c948ef-f6b5-4384-8354-da3a15eca969"),
	// 			IsolationScope: to.Ptr(armmsi.IsolationScopeNone),
	// 			AssignmentRestrictions: &armmsi.AssignmentRestrictions{
	// 				Providers: []*string{
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 		Type: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities"),
	// 	},
	// }
}
