package armapicenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apicenter/armapicenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Services_ListByResourceGroup.json
func ExampleServicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListByResourceGroupPager("contoso-resources", nil)
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
		// page.ServiceListResult = armapicenter.ServiceListResult{
		// 	Value: []*armapicenter.Service{
		// 		{
		// 			Name: to.Ptr("contoso"),
		// 			Type: to.Ptr("Microsoft.ApiCenter/services"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armapicenter.ManagedServiceIdentity{
		// 				Type: to.Ptr(armapicenter.ManagedServiceIdentityType("SystemAssigned, UserAssigned")),
		// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				UserAssignedIdentities: map[string]*armapicenter.UserAssignedIdentity{
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity": &armapicenter.UserAssignedIdentity{
		// 						ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armapicenter.ServiceProperties{
		// 				ProvisioningState: to.Ptr(armapicenter.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
