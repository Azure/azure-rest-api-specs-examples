package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-01-15-preview/GetRaiToolLabel.json
func ExampleRaiToolLabelsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiToolLabelsClient().Get(ctx, "resourceGroupName", "accountName", "ToolLabelName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.RaiToolLabelsClientGetResponse{
	// 	RaiToolLabel: &armcognitiveservices.RaiToolLabel{
	// 		Name: to.Ptr("raiToolLabelName"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/raiIfcToolLabels/raiIfcToolLabelName"),
	// 		Properties: &armcognitiveservices.RaiToolLabelProperties{
	// 			AccountScope: &armcognitiveservices.RaiToolLabelPropertiesAccountScope{
	// 				LabelValues: map[string]*string{
	// 					"confidentiality": to.Ptr("low"),
	// 					"dataDirection": to.Ptr("egress"),
	// 				},
	// 			},
	// 			ProjectScopes: []*armcognitiveservices.RaiToolLabelPropertiesProjectScopesItem{
	// 				{
	// 					LabelValues: map[string]*string{
	// 						"confidentiality": to.Ptr("low"),
	// 						"dataDirection": to.Ptr("egress"),
	// 					},
	// 					Project: to.Ptr("ProjectA"),
	// 				},
	// 				{
	// 					LabelValues: map[string]*string{
	// 						"confidentiality": to.Ptr("high"),
	// 						"dataDirection": to.Ptr("egress"),
	// 					},
	// 					Project: to.Ptr("ProjectB"),
	// 				},
	// 			},
	// 			ToolConnectionName: to.Ptr("SampleToolLabel"),
	// 		},
	// 	},
	// }
}
