package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2dcad6d6e9a96882eb6d317e7500a94be007a9c6/specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomanage.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armautomanage.OperationListResult{
		// 	Value: []*armautomanage.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/register/action"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for the Automanage Resource Provider"),
		// 				Operation: to.Ptr("Register the Automanage Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Automanage Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/configurationProfileAssignments/write"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Create new configuration profile assignment."),
		// 				Operation: to.Ptr("Microsoft.Automanage/configurationProfileAssignments/write"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/configurationProfileAssignments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/configurationProfileAssignments/read"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Get configuration profile assignment."),
		// 				Operation: to.Ptr("Microsoft.Automanage/configurationProfileAssignments/read"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/configurationProfileAssignments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/configurationProfileAssignments/delete"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Delete configuration profile assignment."),
		// 				Operation: to.Ptr("Microsoft.Automanage/configurationProfileAssignments/delete"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/configurationProfileAssignments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/bestPractices/read"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Get Automanage bestPractice."),
		// 				Operation: to.Ptr("Microsoft.Automanage/bestPractices/read"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/bestPractices"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/bestPractices/versions/read"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Get Automanage bestPractice version."),
		// 				Operation: to.Ptr("Microsoft.Automanage/bestPractices/versions/read"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/bestPractices/versions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/configurationProfiles/write"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Create new Automanage Configuration Profile ."),
		// 				Operation: to.Ptr("Microsoft.Automanage/configurationProfiles/write"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/configurationProfiles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/configurationProfiles/read"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Get Automanage Configuration Profile."),
		// 				Operation: to.Ptr("Microsoft.Automanage/configurationProfiles/read"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/configurationProfiles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/configurationProfiles/delete"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Delete Automanage Configuration Profile."),
		// 				Operation: to.Ptr("Microsoft.Automanage/configurationProfiles/delete"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/configurationProfiles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/configurationProfiles/versions/write"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Create new Automanage Configuration Profile version."),
		// 				Operation: to.Ptr("Microsoft.Automanage/configurationProfiles/versions/write"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/configurationProfiles/versions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/configurationProfiles/versions/read"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Get Automanage Configuration Profile version."),
		// 				Operation: to.Ptr("Microsoft.Automanage/configurationProfiles/versions/read"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/configurationProfiles/versions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/configurationProfiles/versions/delete"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Delete Automanage Configuration Profile version."),
		// 				Operation: to.Ptr("Microsoft.Automanage/configurationProfiles/versions/delete"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/configurationProfiles/versions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/configurationProfileAssignments/reports/read"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Get report for configuration profile assignment."),
		// 				Operation: to.Ptr("Microsoft.Automanage/configurationProfileAssignments/reports/read"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/configurationProfileAssignments/reports"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Automanage/servicePrincipals/read"),
		// 			Display: &armautomanage.OperationDisplay{
		// 				Description: to.Ptr("Read the Automanage AAD first party service principal id and authorizationSet for the subscription. This service principal id is used to grant the Contributor RBAC permission to Automanage AAD first party Application."),
		// 				Operation: to.Ptr("Microsoft.Automanage/servicePrincipals/read"),
		// 				Provider: to.Ptr("Microsoft Automanage"),
		// 				Resource: to.Ptr("Microsoft.Automanage/servicePrincipals"),
		// 			},
		// 	}},
		// }
	}
}
