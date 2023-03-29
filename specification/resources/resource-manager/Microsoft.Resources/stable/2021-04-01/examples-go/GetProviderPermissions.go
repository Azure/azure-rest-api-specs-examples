package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fd842fb73656039ec94ce367bcedee25a57bd18/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/GetProviderPermissions.json
func ExampleProvidersClient_ProviderPermissions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().ProviderPermissions(ctx, "Microsoft.TestRP", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProviderPermissionListResult = armresources.ProviderPermissionListResult{
	// 	Value: []*armresources.ProviderPermission{
	// 		{
	// 			ApplicationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ProviderAuthorizationConsentState: to.Ptr(armresources.ProviderAuthorizationConsentStateConsented),
	// 			RoleDefinition: &armresources.RoleDefinition{
	// 				Name: to.Ptr("Contoso service role"),
	// 				ID: to.Ptr("00000000000000000000000000000000"),
	// 				IsServiceRole: to.Ptr(true),
	// 				Permissions: []*armresources.Permission{
	// 					{
	// 						Actions: []*string{
	// 							to.Ptr("Microsoft.Contoso/*")},
	// 							DataActions: []*string{
	// 							},
	// 							NotActions: []*string{
	// 							},
	// 							NotDataActions: []*string{
	// 							},
	// 					}},
	// 					Scopes: []*string{
	// 						to.Ptr("/")},
	// 					},
	// 			}},
	// 		}
}
