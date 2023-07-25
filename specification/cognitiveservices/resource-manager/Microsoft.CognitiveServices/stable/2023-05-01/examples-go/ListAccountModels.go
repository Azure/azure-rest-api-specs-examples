package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListAccountModels.json
func ExampleAccountsClient_NewListModelsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListModelsPager("resourceGroupName", "accountName", nil)
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
		// page.AccountModelListResult = armcognitiveservices.AccountModelListResult{
		// 	Value: []*armcognitiveservices.AccountModel{
		// 		{
		// 			Name: to.Ptr("ada.1"),
		// 			Format: to.Ptr("OpenAI"),
		// 			Version: to.Ptr("1"),
		// 			BaseModel: &armcognitiveservices.DeploymentModel{
		// 				Name: to.Ptr("ada"),
		// 				Format: to.Ptr("OpenAI"),
		// 				Version: to.Ptr("1"),
		// 			},
		// 			Capabilities: map[string]*string{
		// 				"fineTune": to.Ptr("true"),
		// 			},
		// 			Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 				FineTune: to.Ptr("2024-01-01T00:00:00Z"),
		// 				Inference: to.Ptr("2024-01-01T00:00:00Z"),
		// 			},
		// 			MaxCapacity: to.Ptr[int32](10),
		// 		},
		// 		{
		// 			Name: to.Ptr("davinci"),
		// 			Format: to.Ptr("OpenAI"),
		// 			Version: to.Ptr("1"),
		// 			Capabilities: map[string]*string{
		// 				"fineTune": to.Ptr("true"),
		// 			},
		// 			Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 				FineTune: to.Ptr("2024-01-01T00:00:00Z"),
		// 				Inference: to.Ptr("2024-01-01T00:00:00Z"),
		// 			},
		// 			MaxCapacity: to.Ptr[int32](10),
		// 	}},
		// }
	}
}
