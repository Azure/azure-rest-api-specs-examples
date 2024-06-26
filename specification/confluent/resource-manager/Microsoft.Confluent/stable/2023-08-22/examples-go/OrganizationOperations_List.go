package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/OrganizationOperations_List.json
func ExampleOrganizationOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armconfluent.OperationListResult{
		// 	Value: []*armconfluent.OperationResult{
		// 		{
		// 			Name: to.Ptr("Microsoft.Confluent/organizations/Read"),
		// 			Display: &armconfluent.OperationDisplay{
		// 				Description: to.Ptr("Read organization"),
		// 				Operation: to.Ptr("Get/List organization resources"),
		// 				Provider: to.Ptr("Microsoft.Confluent"),
		// 				Resource: to.Ptr("organizations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Confluent/organizations/Write"),
		// 			Display: &armconfluent.OperationDisplay{
		// 				Description: to.Ptr("Write organization"),
		// 				Operation: to.Ptr("Create/Update organization resources"),
		// 				Provider: to.Ptr("Microsoft.Confluent"),
		// 				Resource: to.Ptr("organizations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Confluent/organizations/Delete"),
		// 			Display: &armconfluent.OperationDisplay{
		// 				Description: to.Ptr("Delete organization"),
		// 				Operation: to.Ptr("Delete organization resources"),
		// 				Provider: to.Ptr("Microsoft.Confluent"),
		// 				Resource: to.Ptr("organizations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
