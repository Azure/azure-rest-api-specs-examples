package armaddons_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/addons/armaddons"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/addons/resource-manager/Microsoft.Addons/preview/2018-03-01/examples/Operations_List.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armaddons.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationListValue = armaddons.OperationListValue{
	// 	Value: []*armaddons.OperationsDefinition{
	// 		{
	// 			Name: to.Ptr("Microsoft.Addons/supportProviders/supportPlanTypes/read"),
	// 			Display: &armaddons.OperationsDisplayDefinition{
	// 				Description: to.Ptr("Get the specified Canonical support plan state."),
	// 				Operation: to.Ptr("Get Canonical support plan state"),
	// 				Provider: to.Ptr("Microsoft Addons"),
	// 				Resource: to.Ptr("supportPlanTypes"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Addons/supportProviders/supportPlanTypes/write"),
	// 			Display: &armaddons.OperationsDisplayDefinition{
	// 				Description: to.Ptr("Adds the Canonical support plan type specified."),
	// 				Operation: to.Ptr("Adds a Canonical support plan."),
	// 				Provider: to.Ptr("Microsoft Addons"),
	// 				Resource: to.Ptr("supportPlanTypes"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Addons/supportProviders/supportPlanTypes/delete"),
	// 			Display: &armaddons.OperationsDisplayDefinition{
	// 				Description: to.Ptr("Removes the specified Canonical support plan"),
	// 				Operation: to.Ptr("Removes the Canonical support plan"),
	// 				Provider: to.Ptr("Microsoft Addons"),
	// 				Resource: to.Ptr("supportPlanTypes"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Addons/supportProviders/canonical/supportPlanTypes/get"),
	// 			Display: &armaddons.OperationsDisplayDefinition{
	// 				Description: to.Ptr("Gets the available Canonical support plan types as well as some extra metadata on their enabled status."),
	// 				Operation: to.Ptr("Gets available Canonical support plan types."),
	// 				Provider: to.Ptr("Microsoft Addons"),
	// 				Resource: to.Ptr("supportProviders"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Addons/register/action"),
	// 			Display: &armaddons.OperationsDisplayDefinition{
	// 				Description: to.Ptr("Register the specified subscription with Microsoft.Addons"),
	// 				Operation: to.Ptr("Register for Microsoft.Addons"),
	// 				Provider: to.Ptr("Microsoft Addons"),
	// 				Resource: to.Ptr("register"),
	// 			},
	// 	}},
	// }
}
