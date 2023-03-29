package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateSuppression.json
func ExampleSuppressionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSuppressionsClient().Create(ctx, "resourceUri", "recommendationId", "suppressionName1", armadvisor.SuppressionContract{
		Properties: &armadvisor.SuppressionProperties{
			TTL: to.Ptr("07:00:00:00"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SuppressionContract = armadvisor.SuppressionContract{
	// 	Name: to.Ptr("suppressionName1"),
	// 	Type: to.Ptr("Microsoft.Advisor/suppressions"),
	// 	ID: to.Ptr("/resourceUri/providers/Microsoft.Advisor/recommendations/recommendationId/suppressions/suppressionName1"),
	// 	Properties: &armadvisor.SuppressionProperties{
	// 		SuppressionID: to.Ptr("suppressionId"),
	// 		TTL: to.Ptr("07:00:00:00"),
	// 	},
	// }
}
