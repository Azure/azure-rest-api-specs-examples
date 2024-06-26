package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListProblemClassifications.json
func ExampleProblemClassificationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProblemClassificationsClient().NewListPager("service_guid", nil)
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
		// page.ProblemClassificationsListResult = armsupport.ProblemClassificationsListResult{
		// 	Value: []*armsupport.ProblemClassification{
		// 		{
		// 			Name: to.Ptr("problemClassification_guid_1"),
		// 			Type: to.Ptr("Microsoft.Support/problemClassifications"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/problemClassification_guid_1"),
		// 			Properties: &armsupport.ProblemClassificationProperties{
		// 				DisplayName: to.Ptr("Reservation Management / Exchanges and Refunds"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("problemClassification_guid_2"),
		// 			Type: to.Ptr("Microsoft.Support/problemClassifications"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/problemClassification_guid_2"),
		// 			Properties: &armsupport.ProblemClassificationProperties{
		// 				DisplayName: to.Ptr("Reservation Management / Request Invoices"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("problemClassification_guid_3"),
		// 			Type: to.Ptr("Microsoft.Support/problemClassifications"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/problemClassification_guid_3"),
		// 			Properties: &armsupport.ProblemClassificationProperties{
		// 				DisplayName: to.Ptr("Reservation Management / Other Iissues or Requests"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("problemClassification_guid_4"),
		// 			Type: to.Ptr("Microsoft.Support/problemClassifications"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/problemClassification_guid_4"),
		// 			Properties: &armsupport.ProblemClassificationProperties{
		// 				DisplayName: to.Ptr("Other General Billing Questions"),
		// 			},
		// 	}},
		// }
	}
}
