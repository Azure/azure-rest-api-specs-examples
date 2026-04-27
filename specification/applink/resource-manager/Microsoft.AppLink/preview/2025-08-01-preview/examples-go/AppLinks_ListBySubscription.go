package armappnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appnetwork/armappnetwork"
)

// Generated from example definition: 2025-08-01-preview/AppLinks_ListBySubscription.json
func ExampleAppLinksClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappnetwork.NewClientFactory("11809CA1-E126-4017-945E-AA795CD5C5A9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppLinksClient().NewListBySubscriptionPager(nil)
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
		// page = armappnetwork.AppLinksClientListBySubscriptionResponse{
		// 	AppLinkListResult: armappnetwork.AppLinkListResult{
		// 		Value: []*armappnetwork.AppLink{
		// 			{
		// 				Properties: &armappnetwork.AppLinkProperties{
		// 					ProvisioningState: to.Ptr(armappnetwork.ProvisioningStateSucceeded),
		// 				},
		// 				Identity: &armappnetwork.ManagedServiceIdentity{
		// 					Type: to.Ptr(armappnetwork.ManagedServiceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 					"team": to.Ptr("platform"),
		// 				},
		// 				Location: to.Ptr("westus2"),
		// 				ID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/rg1/providers/Microsoft.AppLink/appLinks/applink-prod-01"),
		// 				Name: to.Ptr("applink-prod-01"),
		// 				Type: to.Ptr("Microsoft.AppLink/appLinks"),
		// 				SystemData: &armappnetwork.SystemData{
		// 					CreatedBy: to.Ptr("user01"),
		// 					CreatedByType: to.Ptr(armappnetwork.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-19T00:28:48.610Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user02"),
		// 					LastModifiedByType: to.Ptr(armappnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-19T00:28:48.610Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armappnetwork.AppLinkProperties{
		// 					ProvisioningState: to.Ptr(armappnetwork.ProvisioningStateSucceeded),
		// 				},
		// 				Identity: &armappnetwork.ManagedServiceIdentity{
		// 					Type: to.Ptr(armappnetwork.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armappnetwork.UserAssignedIdentity{
		// 						"/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/rg2/providers/Microsoft.ManagedIdentity/userAssignedIdentities/applink-identity": &armappnetwork.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("22222222-3333-4444-5555-666666666666"),
		// 							ClientID: to.Ptr("77777777-8888-9999-aaaa-bbbbbbbbbbbb"),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("development"),
		// 					"cost-center": to.Ptr("engineering"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/rg2/providers/Microsoft.AppLink/appLinks/applink-dev-01"),
		// 				Name: to.Ptr("applink-dev-01"),
		// 				Type: to.Ptr("Microsoft.AppLink/appLinks"),
		// 				SystemData: &armappnetwork.SystemData{
		// 					CreatedBy: to.Ptr("dev-user"),
		// 					CreatedByType: to.Ptr(armappnetwork.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-10T14:22:30.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("dev-user"),
		// 					LastModifiedByType: to.Ptr(armappnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T09:15:45.000Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armappnetwork.AppLinkProperties{
		// 					ProvisioningState: to.Ptr(armappnetwork.ProvisioningStateProvisioning),
		// 				},
		// 				Identity: &armappnetwork.ManagedServiceIdentity{
		// 					Type: to.Ptr(armappnetwork.ManagedServiceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("33333333-4444-5555-6666-777777777777"),
		// 					TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("staging"),
		// 				},
		// 				Location: to.Ptr("northeurope"),
		// 				ID: to.Ptr("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/rg3/providers/Microsoft.AppLink/appLinks/applink-staging-01"),
		// 				Name: to.Ptr("applink-staging-01"),
		// 				Type: to.Ptr("Microsoft.AppLink/appLinks"),
		// 				SystemData: &armappnetwork.SystemData{
		// 					CreatedBy: to.Ptr("staging-user"),
		// 					CreatedByType: to.Ptr(armappnetwork.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T08:30:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("staging-user"),
		// 					LastModifiedByType: to.Ptr(armappnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T08:30:00.000Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/providers/Microsoft.AppLink/appLinks?api-version=2025-08-01-preview&$skiptoken=eyJjb250aW51YXRpb24iOiAiYXBwbGluay1zdGFnaW5nLTAyIn0%3D"),
		// 	},
		// }
	}
}
