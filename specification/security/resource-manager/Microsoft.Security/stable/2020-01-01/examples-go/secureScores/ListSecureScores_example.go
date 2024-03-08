package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScores/ListSecureScores_example.json
func ExampleSecureScoresClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecureScoresClient().NewListPager(nil)
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
		// page.SecureScoresList = armsecurity.SecureScoresList{
		// 	Value: []*armsecurity.SecureScoreItem{
		// 		{
		// 			Name: to.Ptr("ascScore"),
		// 			Type: to.Ptr("Microsoft.Security/secureScores"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/secureScores/ascScore"),
		// 			Properties: &armsecurity.SecureScoreItemProperties{
		// 				DisplayName: to.Ptr("ASC score"),
		// 				Score: &armsecurity.ScoreDetails{
		// 					Current: to.Ptr[float64](23.53),
		// 					Max: to.Ptr[int32](39),
		// 					Percentage: to.Ptr[float64](0.6033),
		// 				},
		// 				Weight: to.Ptr[int64](67),
		// 			},
		// 	}},
		// }
	}
}
