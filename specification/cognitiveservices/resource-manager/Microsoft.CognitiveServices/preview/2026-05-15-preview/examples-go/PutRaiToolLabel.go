package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/PutRaiToolLabel.json
func ExampleRaiToolLabelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiToolLabelsClient().CreateOrUpdate(ctx, "resourceGroupName", "accountName", "Web_Search", armcognitiveservices.RaiToolLabel{
		Properties: &armcognitiveservices.RaiToolLabelProperties{
			AccountScope: &armcognitiveservices.RaiToolLabelPropertiesAccountScope{
				LabelValues: map[string]*string{
					"confidentiality": to.Ptr("low"),
				},
			},
			ProjectScopes: []*armcognitiveservices.RaiToolLabelPropertiesProjectScopesItem{
				{
					LabelValues: map[string]*string{
						"confidentiality": to.Ptr("low"),
					},
					Project: to.Ptr("test-project"),
				},
				{
					LabelValues: map[string]*string{
						"confidentiality": to.Ptr("low"),
					},
					Project: to.Ptr("sample-project"),
				},
			},
			ToolConnectionName: to.Ptr("Web_Search"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.RaiToolLabelsClientCreateOrUpdateResponse{
	// 	RaiToolLabel: armcognitiveservices.RaiToolLabel{
	// 		Name: to.Ptr("Web_Search"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiToolLabels"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/raiToolLabels/Web_Search"),
	// 		Properties: &armcognitiveservices.RaiToolLabelProperties{
	// 			AccountScope: &armcognitiveservices.RaiToolLabelPropertiesAccountScope{
	// 				LabelValues: map[string]*string{
	// 					"confidentiality": to.Ptr("low"),
	// 				},
	// 			},
	// 			ProjectScopes: []*armcognitiveservices.RaiToolLabelPropertiesProjectScopesItem{
	// 				{
	// 					LabelValues: map[string]*string{
	// 						"confidentiality": to.Ptr("low"),
	// 					},
	// 					Project: to.Ptr("test-project"),
	// 				},
	// 				{
	// 					LabelValues: map[string]*string{
	// 						"confidentiality": to.Ptr("low"),
	// 					},
	// 					Project: to.Ptr("sample-project"),
	// 				},
	// 			},
	// 			ToolConnectionName: to.Ptr("Web_Search"),
	// 		},
	// 	},
	// }
}
