package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationResultList = armmigrate.OperationResultList{
		// 	Value: []*armmigrate.Operation{
		// 		{
		// 			Name: to.Ptr("Read"),
		// 			Display: &armmigrate.OperationDisplay{
		// 				Description: to.Ptr("Reads a project"),
		// 				Operation: to.Ptr("Read"),
		// 				Provider: to.Ptr("Microsoft.Migrate"),
		// 				Resource: to.Ptr("Microsoft.Migrate/assessmentProjects"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Write"),
		// 			Display: &armmigrate.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates a project"),
		// 				Operation: to.Ptr("Write"),
		// 				Provider: to.Ptr("Microsoft.Migrate"),
		// 				Resource: to.Ptr("Microsoft.Migrate/assessmentProjects"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
