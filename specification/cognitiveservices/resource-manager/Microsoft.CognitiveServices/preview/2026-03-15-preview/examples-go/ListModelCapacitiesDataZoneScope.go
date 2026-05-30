package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-03-15-preview/ListModelCapacitiesDataZoneScope.json
func ExampleModelCapacitiesClient_NewListPager_listModelCapacitiesDataZoneScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewModelCapacitiesClient().NewListPager("OpenAI", "ada", "1", nil)
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
		// page = armcognitiveservices.ModelCapacitiesClientListResponse{
		// 	ModelCapacityListResult: armcognitiveservices.ModelCapacityListResult{
		// 		Value: []*armcognitiveservices.ModelCapacityListResultValueItem{
		// 			{
		// 				Name: to.Ptr("DataZoneStandard"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/locations/models/skuCapacities"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CognitiveServices/locations/WestUS/models/OpenAI.ada.1/skuCapacities/DataZoneStandard"),
		// 				Location: to.Ptr("WestUS"),
		// 				Properties: &armcognitiveservices.ModelSKUCapacityProperties{
		// 					AvailableCapacity: to.Ptr[float32](400),
		// 					AvailableFinetuneCapacity: to.Ptr[float32](25),
		// 					Model: &armcognitiveservices.DeploymentModel{
		// 						Name: to.Ptr("ada"),
		// 						Format: to.Ptr("OpenAI"),
		// 						Version: to.Ptr("1"),
		// 					},
		// 					SKUName: to.Ptr("DataZoneStandard"),
		// 					ScopeID: to.Ptr("US"),
		// 					ScopeType: to.Ptr(armcognitiveservices.QuotaScopeTypeDataZone),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
