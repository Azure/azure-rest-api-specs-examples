package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armmigrationassessment.OperationListResult{
		// 	Value: []*armmigrationassessment.Operation{
		// 		{
		// 			Name: to.Ptr("AksAssessmentOperations_Get"),
		// 			ActionType: to.Ptr(armmigrationassessment.ActionTypeInternal),
		// 			Display: &armmigrationassessment.OperationDisplay{
		// 				Description: to.Ptr("AKSAssessment"),
		// 				Operation: to.Ptr("GET"),
		// 				Provider: to.Ptr("Microsoft.Migrate"),
		// 				Resource: to.Ptr("AKSAssessment"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armmigrationassessment.OriginUser),
		// 	}},
		// }
	}
}
