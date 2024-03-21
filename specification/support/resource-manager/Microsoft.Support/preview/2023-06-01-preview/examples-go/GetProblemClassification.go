package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/GetProblemClassification.json
func ExampleProblemClassificationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProblemClassificationsClient().Get(ctx, "service_guid", "problemClassification_guid", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProblemClassification = armsupport.ProblemClassification{
	// 	Name: to.Ptr("problemClassification_guid"),
	// 	Type: to.Ptr("Microsoft.Support/problemClassifications"),
	// 	ID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/problemClassification_guid"),
	// 	Properties: &armsupport.ProblemClassificationProperties{
	// 		DisplayName: to.Ptr("Reservation Management / Exchanges and Refunds"),
	// 		Metadata: map[string]*string{
	// 			"azureSubscriptionRequired": to.Ptr("true"),
	// 			"quotaautomationenabled": to.Ptr("true"),
	// 			"state": to.Ptr("public"),
	// 		},
	// 		ParentProblemClassification: &armsupport.ProblemClassification{
	// 			Name: to.Ptr("problemClassification_guid_1"),
	// 			Type: to.Ptr("Microsoft.Support/problemClassifications"),
	// 			ID: to.Ptr("/providers/Microsoft.Support/services/service_guid/problemClassifications/problemClassification_guid_1"),
	// 			Properties: &armsupport.ProblemClassificationProperties{
	// 				DisplayName: to.Ptr("Reservation Management"),
	// 				Metadata: map[string]*string{
	// 					"azureSubscriptionRequired": to.Ptr("true"),
	// 					"quotaautomationenabled": to.Ptr("true"),
	// 					"state": to.Ptr("public"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
