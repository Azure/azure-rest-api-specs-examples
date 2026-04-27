package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-01-15-preview/ListRaiToolLabels.json
func ExampleRaiToolLabelsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRaiToolLabelsClient().NewListPager("resourceGroupName", "accountName", nil)
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
		// page = armcognitiveservices.RaiToolLabelsClientListResponse{
		// 	RaiToolLabelResult: armcognitiveservices.RaiToolLabelResult{
		// 		Value: []*armcognitiveservices.RaiToolLabel{
		// 			{
		// 				Name: to.Ptr("raiIfcToolLabelName"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/raiIfcToolLabels/raiIfcToolLabelName"),
		// 				Properties: &armcognitiveservices.RaiToolLabelProperties{
		// 					AccountScope: &armcognitiveservices.RaiToolLabelPropertiesAccountScope{
		// 						LabelValues: map[string]*string{
		// 							"confidentiality": to.Ptr("low"),
		// 						},
		// 					},
		// 					ProjectScopes: []*armcognitiveservices.RaiToolLabelPropertiesProjectScopesItem{
		// 						{
		// 							LabelValues: map[string]*string{
		// 								"confidentiality": to.Ptr("low"),
		// 							},
		// 							Project: to.Ptr("hualiang-project"),
		// 						},
		// 						{
		// 							LabelValues: map[string]*string{
		// 								"confidentiality": to.Ptr("low"),
		// 							},
		// 							Project: to.Ptr("sahana-project"),
		// 						},
		// 					},
		// 					ToolConnectionName: to.Ptr("Web_Search"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
