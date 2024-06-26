package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListProblemClassifications.json
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
		// 				SecondaryConsentEnabled: []*armsupport.SecondaryConsentEnabled{
		// 					{
		// 						Type: to.Ptr("DatabricksConsent"),
		// 						Description: to.Ptr("For faster resolution, allow Microsoft and Databricks to temporarily have read and write access to your Databricks workspace. We will only access to read and write to your cluster for the purpose of resolving your support issue and in conformance with Microsoft's Privacy Policy."),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("problemClassification_guid_2"),
		// 			Type: to.Ptr("Microsoft.Support/problemClassifications"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/problemClassification_guid_2"),
		// 			Properties: &armsupport.ProblemClassificationProperties{
		// 				DisplayName: to.Ptr("Reservation Management / Request Invoices"),
		// 				SecondaryConsentEnabled: []*armsupport.SecondaryConsentEnabled{
		// 					{
		// 						Type: to.Ptr("DatabricksConsent"),
		// 						Description: to.Ptr("For faster resolution, allow Microsoft and Databricks to temporarily have read and write access to your Databricks workspace. We will only access to read and write to your cluster for the purpose of resolving your support issue and in conformance with Microsoft's Privacy Policy."),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("problemClassification_guid_3"),
		// 			Type: to.Ptr("Microsoft.Support/problemClassifications"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/problemClassification_guid_3"),
		// 			Properties: &armsupport.ProblemClassificationProperties{
		// 				DisplayName: to.Ptr("Reservation Management / Other Iissues or Requests"),
		// 				SecondaryConsentEnabled: []*armsupport.SecondaryConsentEnabled{
		// 					{
		// 						Type: to.Ptr("DatabricksConsent"),
		// 						Description: to.Ptr("For faster resolution, allow Microsoft and Databricks to temporarily have read and write access to your Databricks workspace. We will only access to read and write to your cluster for the purpose of resolving your support issue and in conformance with Microsoft's Privacy Policy."),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("problemClassification_guid_4"),
		// 			Type: to.Ptr("Microsoft.Support/problemClassifications"),
		// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/problemClassification_guid_4"),
		// 			Properties: &armsupport.ProblemClassificationProperties{
		// 				DisplayName: to.Ptr("Other General Billing Questions"),
		// 				SecondaryConsentEnabled: []*armsupport.SecondaryConsentEnabled{
		// 					{
		// 						Type: to.Ptr("DatabricksConsent"),
		// 						Description: to.Ptr("For faster resolution, allow Microsoft and Databricks to temporarily have read and write access to your Databricks workspace. We will only access to read and write to your cluster for the purpose of resolving your support issue and in conformance with Microsoft's Privacy Policy."),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
