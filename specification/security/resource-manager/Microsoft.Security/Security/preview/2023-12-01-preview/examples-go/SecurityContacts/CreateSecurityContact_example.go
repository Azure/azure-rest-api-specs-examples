package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2023-12-01-preview/SecurityContacts/CreateSecurityContact_example.json
func ExampleContactsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContactsClient().Create(ctx, armsecurity.SecurityContactNameDefault, armsecurity.Contact{
		Properties: &armsecurity.ContactProperties{
			Emails:    to.Ptr("john@contoso.com;jane@contoso.com"),
			IsEnabled: to.Ptr(true),
			NotificationsByRole: &armsecurity.ContactPropertiesNotificationsByRole{
				Roles: []*armsecurity.SecurityContactRole{
					to.Ptr(armsecurity.SecurityContactRoleOwner),
				},
				State: to.Ptr(armsecurity.StateOn),
			},
			NotificationsSources: []armsecurity.NotificationsSourceClassification{
				&armsecurity.NotificationsSourceAttackPath{
					MinimalRiskLevel: to.Ptr(armsecurity.MinimalRiskLevelCritical),
					SourceType:       to.Ptr(armsecurity.SourceTypeAttackPath),
				},
				&armsecurity.NotificationsSourceAlert{
					MinimalSeverity: to.Ptr(armsecurity.MinimalSeverityMedium),
					SourceType:      to.Ptr(armsecurity.SourceTypeAlert),
				},
			},
			Phone: to.Ptr("(214)275-4038"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.ContactsClientCreateResponse{
	// 	Contact: armsecurity.Contact{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Security/securityContact"),
	// 		ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/securityContacts/default"),
	// 		Properties: &armsecurity.ContactProperties{
	// 			Emails: to.Ptr("john@microsoft.com;jane@microsoft.com"),
	// 			IsEnabled: to.Ptr(true),
	// 			NotificationsByRole: &armsecurity.ContactPropertiesNotificationsByRole{
	// 				Roles: []*armsecurity.SecurityContactRole{
	// 					to.Ptr(armsecurity.SecurityContactRoleOwner),
	// 				},
	// 				State: to.Ptr(armsecurity.StateOn),
	// 			},
	// 			NotificationsSources: []armsecurity.NotificationsSourceClassification{
	// 				&armsecurity.NotificationsSourceAttackPath{
	// 					MinimalRiskLevel: to.Ptr(armsecurity.MinimalRiskLevelCritical),
	// 					SourceType: to.Ptr(armsecurity.SourceTypeAttackPath),
	// 				},
	// 				&armsecurity.NotificationsSourceAlert{
	// 					MinimalSeverity: to.Ptr(armsecurity.MinimalSeverityMedium),
	// 					SourceType: to.Ptr(armsecurity.SourceTypeAlert),
	// 				},
	// 			},
	// 			Phone: to.Ptr("(214)275-4038"),
	// 		},
	// 	},
	// }
}
