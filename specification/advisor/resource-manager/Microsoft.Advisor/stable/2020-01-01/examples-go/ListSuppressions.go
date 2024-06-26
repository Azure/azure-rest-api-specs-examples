package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListSuppressions.json
func ExampleSuppressionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSuppressionsClient().NewListPager(&armadvisor.SuppressionsClientListOptions{Top: nil,
		SkipToken: nil,
	})
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
		// page.SuppressionContractListResult = armadvisor.SuppressionContractListResult{
		// 	Value: []*armadvisor.SuppressionContract{
		// 		{
		// 			Name: to.Ptr("suppressionName1"),
		// 			Type: to.Ptr("Microsoft.Advisor/suppressions"),
		// 			ID: to.Ptr("/resourceUri/providers/Microsoft.Advisor/recommendations/recommendationId/suppressions/suppressionName1"),
		// 			Properties: &armadvisor.SuppressionProperties{
		// 				ExpirationTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-24T22:24:43.321Z"); return t}()),
		// 				SuppressionID: to.Ptr("suppressionId1"),
		// 				TTL: to.Ptr("7.00:00:00"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("suppressionName2"),
		// 			Type: to.Ptr("Microsoft.Advisor/suppressions"),
		// 			ID: to.Ptr("/resourceUri/providers/Microsoft.Advisor/recommendations/recommendationId/suppressions/suppressionName2"),
		// 			Properties: &armadvisor.SuppressionProperties{
		// 				ExpirationTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-25T22:24:43.321Z"); return t}()),
		// 				SuppressionID: to.Ptr("suppressionId2"),
		// 				TTL: to.Ptr("7.00:00:00"),
		// 			},
		// 	}},
		// }
	}
}
