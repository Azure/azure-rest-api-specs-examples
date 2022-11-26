package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/SecurityContacts/CreateSecurityContact_example.json
func ExampleContactsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewContactsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default", armsecurity.Contact{
		Properties: &armsecurity.ContactProperties{
			AlertNotifications: &armsecurity.ContactPropertiesAlertNotifications{
				MinimalSeverity: to.Ptr(armsecurity.MinimalSeverityLow),
				State:           to.Ptr(armsecurity.State("On")),
			},
			Emails: to.Ptr("john@contoso.com;jane@contoso.com"),
			NotificationsByRole: &armsecurity.ContactPropertiesNotificationsByRole{
				Roles: []*armsecurity.Roles{
					to.Ptr(armsecurity.RolesOwner)},
				State: to.Ptr(armsecurity.State("On")),
			},
			Phone: to.Ptr("(214)275-4038"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
