package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2023-12-01-preview/SecurityContacts/GetSecurityContactsSubscription_example.json
func ExampleContactsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
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
		// page = armsecurity.ContactsClientListResponse{
		// 	ContactList: armsecurity.ContactList{
		// 		Value: []*armsecurity.Contact{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.Security/securityContact"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/securityContact/default"),
		// 				Properties: &armsecurity.ContactProperties{
		// 					Emails: to.Ptr("john@contoso.com;Jane@contoso.com"),
		// 					IsEnabled: to.Ptr(true),
		// 					NotificationsByRole: &armsecurity.ContactPropertiesNotificationsByRole{
		// 						Roles: []*armsecurity.SecurityContactRole{
		// 							to.Ptr(armsecurity.SecurityContactRoleOwner),
		// 							to.Ptr(armsecurity.SecurityContactRole("Admin")),
		// 						},
		// 						State: to.Ptr(armsecurity.StateOn),
		// 					},
		// 					NotificationsSources: []armsecurity.NotificationsSourceClassification{
		// 						&armsecurity.NotificationsSourceAttackPath{
		// 							MinimalRiskLevel: to.Ptr(armsecurity.MinimalRiskLevelCritical),
		// 							SourceType: to.Ptr(armsecurity.SourceTypeAttackPath),
		// 						},
		// 						&armsecurity.NotificationsSourceAlert{
		// 							MinimalSeverity: to.Ptr(armsecurity.MinimalSeverityMedium),
		// 							SourceType: to.Ptr(armsecurity.SourceTypeAlert),
		// 						},
		// 					},
		// 					Phone: to.Ptr("(214)275-4038"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
