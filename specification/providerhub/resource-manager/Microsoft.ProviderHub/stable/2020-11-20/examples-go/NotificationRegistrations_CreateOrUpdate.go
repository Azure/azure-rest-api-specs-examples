package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/NotificationRegistrations_CreateOrUpdate.json
func ExampleNotificationRegistrationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNotificationRegistrationsClient().CreateOrUpdate(ctx, "Microsoft.Contoso", "fooNotificationRegistration", armproviderhub.NotificationRegistration{
		Properties: &armproviderhub.NotificationRegistrationProperties{
			IncludedEvents: []*string{
				to.Ptr("*/write"),
				to.Ptr("Microsoft.Contoso/employees/delete")},
			MessageScope: to.Ptr(armproviderhub.MessageScopeRegisteredSubscriptions),
			NotificationEndpoints: []*armproviderhub.NotificationEndpoint{
				{
					Locations: []*string{
						to.Ptr(""),
						to.Ptr("East US")},
					NotificationDestination: to.Ptr("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-eastus/providers/Microsoft.EventHub/namespaces/unitedstates-mgmtexpint/eventhubs/armlinkednotifications"),
				},
				{
					Locations: []*string{
						to.Ptr("North Europe")},
					NotificationDestination: to.Ptr("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-northeurope/providers/Microsoft.EventHub/namespaces/europe-mgmtexpint/eventhubs/armlinkednotifications"),
				}},
			NotificationMode: to.Ptr(armproviderhub.NotificationModeEventHub),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NotificationRegistration = armproviderhub.NotificationRegistration{
	// 	Name: to.Ptr("fooNotificationRegistration"),
	// 	Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/notificationregistrations"),
	// 	ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/notificationregistrations/fooNotificationRegistration"),
	// 	Properties: &armproviderhub.NotificationRegistrationProperties{
	// 		IncludedEvents: []*string{
	// 			to.Ptr("*/write"),
	// 			to.Ptr("Microsoft.Contoso/employees/delete")},
	// 			MessageScope: to.Ptr(armproviderhub.MessageScopeRegisteredSubscriptions),
	// 			NotificationEndpoints: []*armproviderhub.NotificationEndpoint{
	// 				{
	// 					Locations: []*string{
	// 						to.Ptr(""),
	// 						to.Ptr("East US")},
	// 						NotificationDestination: to.Ptr("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-eastus/providers/Microsoft.EventHub/namespaces/unitedstates-mgmtexpint/eventhubs/armlinkednotifications"),
	// 					},
	// 					{
	// 						Locations: []*string{
	// 							to.Ptr("North Europe")},
	// 							NotificationDestination: to.Ptr("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-northeurope/providers/Microsoft.EventHub/namespaces/europe-mgmtexpint/eventhubs/armlinkednotifications"),
	// 					}},
	// 					NotificationMode: to.Ptr(armproviderhub.NotificationModeEventHub),
	// 				},
	// 			}
}
