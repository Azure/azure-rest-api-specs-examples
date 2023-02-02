package armmsi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3d7a3848106b831a4a7f46976fe38aa605c4f44d/specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/IdentityUpdate.json
func ExampleUserAssignedIdentitiesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmsi.NewUserAssignedIdentitiesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "rgName", "resourceName", armmsi.IdentityUpdate{
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
	// res.Identity = armmsi.Identity{
	// 	Name: to.Ptr("identityName"),
	// 	Type: to.Ptr("Microsoft.ManagedIdentity/userAssignedIdentities"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rgName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armmsi.UserAssignedIdentityProperties{
	// 		ClientID: to.Ptr("4024ab25-56a8-4370-aea6-6389221caf29"),
	// 		PrincipalID: to.Ptr("25cc773c-7f05-40fc-a104-32d2300754ad"),
	// 		TenantID: to.Ptr("b6c948ef-f6b5-4384-8354-da3a15eca969"),
	// 	},
	// }
}
