package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2023-12-01-preview/examples/SecurityContacts/GetSecurityContactsSubscription_example.json
func ExampleContactsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContactsClient().NewListPager(nil)
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
		// page.ContactList = armsecurity.ContactList{
		// 	Value: []*armsecurity.Contact{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Security/securityContact"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/securityContact/default"),
		// 			Properties: &armsecurity.ContactProperties{
		// 				Emails: to.Ptr("john@contoso.com;Jane@contoso.com"),
		// 				IsEnabled: to.Ptr(true),
		// 				NotificationsByRole: &armsecurity.ContactPropertiesNotificationsByRole{
		// 					Roles: []*armsecurity.SecurityContactRole{
		// 						to.Ptr(armsecurity.SecurityContactRoleOwner),
		// 						to.Ptr(armsecurity.SecurityContactRole("Admin"))},
		// 						State: to.Ptr(armsecurity.State("On")),
		// 					},
		// 					NotificationsSources: []armsecurity.NotificationsSourceClassification{
		// 						&armsecurity.NotificationsSourceAttackPath{
		// 							SourceType: to.Ptr(armsecurity.SourceTypeAttackPath),
		// 							MinimalRiskLevel: to.Ptr(armsecurity.MinimalRiskLevelCritical),
		// 						},
		// 						&armsecurity.NotificationsSourceAlert{
		// 							SourceType: to.Ptr(armsecurity.SourceTypeAlert),
		// 							MinimalSeverity: to.Ptr(armsecurity.MinimalSeverityMedium),
		// 					}},
		// 					Phone: to.Ptr("(214)275-4038"),
		// 				},
		// 		}},
		// 	}
	}
}
