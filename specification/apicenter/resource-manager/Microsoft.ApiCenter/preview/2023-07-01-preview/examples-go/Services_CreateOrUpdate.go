package armapicenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apicenter/armapicenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apicenter/resource-manager/Microsoft.ApiCenter/preview/2023-07-01-preview/examples/Services_CreateOrUpdate.json
func ExampleServicesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("<subscription-id>", "<service-name>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().CreateOrUpdate(ctx, "contoso-resources", armapicenter.Service{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Service = armapicenter.Service{
	// 	Name: to.Ptr("contoso"),
	// 	Type: to.Ptr("Microsoft.ApiCenter/services"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armapicenter.ManagedServiceIdentity{
	// 		Type: to.Ptr(armapicenter.ManagedServiceIdentityType("SystemAssigned, UserAssigned")),
	// 		PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		UserAssignedIdentities: map[string]*armapicenter.UserAssignedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity": &armapicenter.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armapicenter.ServiceProperties{
	// 		ProvisioningState: to.Ptr(armapicenter.ProvisioningStateSucceeded),
	// 	},
	// }
}
