package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/NotificationRegistrations_ListByProviderRegistration.json
func ExampleNotificationRegistrationsClient_NewListByProviderRegistrationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNotificationRegistrationsClient().NewListByProviderRegistrationPager("Microsoft.Contoso", nil)
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
		// page = armproviderhub.NotificationRegistrationsClientListByProviderRegistrationResponse{
		// 	NotificationRegistrationArrayResponseWithContinuation: armproviderhub.NotificationRegistrationArrayResponseWithContinuation{
		// 		Value: []*armproviderhub.NotificationRegistration{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/notificationRegistration/2020week01"),
		// 				Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/notificationRegistration"),
		// 				Name: to.Ptr("fooNotificationRegistration"),
		// 				SystemData: &armproviderhub.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 				},
		// 				Properties: &armproviderhub.NotificationRegistrationProperties{
		// 					NotificationMode: to.Ptr(armproviderhub.NotificationModeEventHub),
		// 					MessageScope: to.Ptr(armproviderhub.MessageScopeRegisteredSubscriptions),
		// 					IncludedEvents: []*string{
		// 						to.Ptr("*/write"),
		// 						to.Ptr("Microsoft.Contoso/employees/delete"),
		// 					},
		// 					NotificationEndpoints: []*armproviderhub.NotificationEndpoint{
		// 						{
		// 							NotificationDestination: to.Ptr("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-eastus/providers/Microsoft.EventHub/namespaces/unitedstates-mgmtexpint/eventhubs/armlinkednotifications"),
		// 							Locations: []*string{
		// 								to.Ptr(""),
		// 								to.Ptr("East US"),
		// 							},
		// 						},
		// 						{
		// 							NotificationDestination: to.Ptr("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-northeurope/providers/Microsoft.EventHub/namespaces/europe-mgmtexpint/eventhubs/armlinkednotifications"),
		// 							Locations: []*string{
		// 								to.Ptr("North Europe"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/notificationRegistration/2020week01"),
		// 				Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/notificationRegistration"),
		// 				Name: to.Ptr("barNotificationRegistration"),
		// 				SystemData: &armproviderhub.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 				},
		// 				Properties: &armproviderhub.NotificationRegistrationProperties{
		// 					NotificationMode: to.Ptr(armproviderhub.NotificationModeEventHub),
		// 					MessageScope: to.Ptr(armproviderhub.MessageScopeRegisteredSubscriptions),
		// 					IncludedEvents: []*string{
		// 						to.Ptr("*/delete"),
		// 					},
		// 					NotificationEndpoints: []*armproviderhub.NotificationEndpoint{
		// 						{
		// 							NotificationDestination: to.Ptr("/subscriptions/ac6bcfb5-3dc1-491f-95a6-646b89bf3e88/resourceGroups/mgmtexp-eastus/providers/Microsoft.EventHub/namespaces/unitedstates-mgmtexpint/eventhubs/armlinkednotifications"),
		// 							Locations: []*string{
		// 								to.Ptr(""),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
