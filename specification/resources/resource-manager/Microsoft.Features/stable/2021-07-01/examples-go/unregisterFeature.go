package armfeatures_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/unregisterFeature.json
func ExampleClient_Unregister() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfeatures.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Unregister(ctx, "Resource Provider Namespace", "feature", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FeatureResult = armfeatures.FeatureResult{
	// 	Name: to.Ptr("Feature1"),
	// 	Type: to.Ptr("Microsoft.Features/providers/features"),
	// 	ID: to.Ptr("/subscriptions/ff23096b-f5a2-46ea-bd62-59c3e93fef9a/providers/Microsoft.Features/providers/Microsoft.Test/features/Feature1"),
	// 	Properties: &armfeatures.FeatureProperties{
	// 		State: to.Ptr("unregistered"),
	// 	},
	// }
}
