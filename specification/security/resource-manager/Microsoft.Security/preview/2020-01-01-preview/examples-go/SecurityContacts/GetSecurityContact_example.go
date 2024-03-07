package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/SecurityContacts/GetSecurityContact_example.json
func ExampleContactsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContactsClient().Get(ctx, "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Contact = armsecurity.Contact{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Security/securityContacts"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/securityContacts/default"),
	// 	Properties: &armsecurity.ContactProperties{
	// 		AlertNotifications: &armsecurity.ContactPropertiesAlertNotifications{
	// 			MinimalSeverity: to.Ptr(armsecurity.MinimalSeverityLow),
	// 			State: to.Ptr(armsecurity.State("On")),
	// 		},
	// 		Emails: to.Ptr("john@contoso.com;jane@contoso.com"),
	// 		NotificationsByRole: &armsecurity.ContactPropertiesNotificationsByRole{
	// 			Roles: []*armsecurity.Roles{
	// 				to.Ptr(armsecurity.RolesOwner)},
	// 				State: to.Ptr(armsecurity.State("On")),
	// 			},
	// 			Phone: to.Ptr("(214)275-4038"),
	// 		},
	// 	}
}
