package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-05-01-preview/HealthModels_ListBySubscription.json
func ExampleHealthModelsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHealthModelsClient().NewListBySubscriptionPager(nil)
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
		// page = armcloudhealth.HealthModelsClientListBySubscriptionResponse{
		// 	HealthModelListResult: armcloudhealth.HealthModelListResult{
		// 		Value: []*armcloudhealth.HealthModel{
		// 			{
		// 				Properties: &armcloudhealth.HealthModelProperties{
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 				},
		// 				Identity: &armcloudhealth.ManagedServiceIdentity{
		// 					Type: to.Ptr(armcloudhealth.ManagedServiceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 					"team": to.Ptr("online-store"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.CloudHealth/healthmodels/online-store"),
		// 				Name: to.Ptr("online-store"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T08:15:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armcloudhealth.HealthModelProperties{
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 				},
		// 				Identity: &armcloudhealth.ManagedServiceIdentity{
		// 					Type: to.Ptr(armcloudhealth.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armcloudhealth.UserAssignedIdentity{
		// 						"/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/online-store-identity": &armcloudhealth.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("33333333-3333-3333-3333-333333333333"),
		// 							ClientID: to.Ptr("44444444-4444-4444-4444-444444444444"),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("staging"),
		// 					"team": to.Ptr("online-store"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.CloudHealth/healthmodels/online-store-staging"),
		// 				Name: to.Ptr("online-store-staging"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T08:15:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
