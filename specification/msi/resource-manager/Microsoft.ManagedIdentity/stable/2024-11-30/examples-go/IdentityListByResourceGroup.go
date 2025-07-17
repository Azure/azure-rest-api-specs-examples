package armmsi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2024-11-30/examples/IdentityListByResourceGroup.json
func ExampleUserAssignedIdentitiesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmsi.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUserAssignedIdentitiesClient().NewListByResourceGroupPager("rgName", nil)
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
		// page.UserAssignedIdentitiesListResult = armmsi.UserAssignedIdentitiesListResult{
		// 	Value: []*armmsi.Identity{
		// 		{
		// 			Name: to.Ptr("identityName"),
		// 			Type: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armmsi.UserAssignedIdentityProperties{
		// 				ClientID: to.Ptr("4024ab25-56a8-4370-aea6-6389221caf29"),
		// 				IsolationScope: to.Ptr(armmsi.IsolationScopeNone),
		// 				PrincipalID: to.Ptr("25cc773c-7f05-40fc-a104-32d2300754ad"),
		// 				TenantID: to.Ptr("b6c948ef-f6b5-4384-8354-da3a15eca969"),
		// 			},
		// 	}},
		// }
	}
}
