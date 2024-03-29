package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/listOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armguestconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationList = armguestconfiguration.OperationList{
		// 	Value: []*armguestconfiguration.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.GuestConfiguration/guestConfigurationAssignments/write"),
		// 			Display: &armguestconfiguration.OperationDisplay{
		// 				Description: to.Ptr("Create new guest configuration assignment."),
		// 				Operation: to.Ptr("Microsoft.GuestConfiguration/guestConfigurationAssignments/write"),
		// 				Provider: to.Ptr("Microsoft Guest Configuration"),
		// 				Resource: to.Ptr("Microsoft.GuestConfiguration/guestConfigurationAssignments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.GuestConfiguration/register/action"),
		// 			Display: &armguestconfiguration.OperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for the Microsoft.GuestConfiguration resource provider."),
		// 				Operation: to.Ptr("Registers the feature for Microsoft.GuestConfiguration."),
		// 				Provider: to.Ptr("Microsoft Guest Configuration"),
		// 				Resource: to.Ptr("Register"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.GuestConfiguration/guestConfigurationAssignments/read"),
		// 			Display: &armguestconfiguration.OperationDisplay{
		// 				Description: to.Ptr("Get guest configuration assignment."),
		// 				Operation: to.Ptr("Microsoft.GuestConfiguration/guestConfigurationAssignments/read"),
		// 				Provider: to.Ptr("Microsoft Guest Configuration"),
		// 				Resource: to.Ptr("Microsoft.GuestConfiguration/guestConfigurationAssignments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.GuestConfiguration/guestConfigurationAssignments/reports/read"),
		// 			Display: &armguestconfiguration.OperationDisplay{
		// 				Description: to.Ptr("Get guest configuration assignment report."),
		// 				Operation: to.Ptr("Microsoft.GuestConfiguration/guestConfigurationAssignments/reports/read"),
		// 				Provider: to.Ptr("Microsoft Guest Configuration"),
		// 				Resource: to.Ptr("Microsoft.GuestConfiguration/guestConfigurationAssignments"),
		// 			},
		// 	}},
		// }
	}
}
