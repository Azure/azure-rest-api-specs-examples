package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/metadata/GetAllMetadataOData.json
func ExampleMetadataClient_NewListPager_getAllMetadataWithODataFilterOrderbySkipTop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetadataClient().NewListPager("myRg", "myWorkspace", &armsecurityinsights.MetadataClientListOptions{Filter: nil,
		Orderby: nil,
		Top:     nil,
		Skip:    nil,
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
		// page.MetadataList = armsecurityinsights.MetadataList{
		// 	Value: []*armsecurityinsights.MetadataModel{
		// 		{
		// 			Name: to.Ptr("metadataName1"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
		// 			ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName1"),
		// 			Properties: &armsecurityinsights.MetadataProperties{
		// 				ContentID: to.Ptr("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
		// 				Kind: to.Ptr(armsecurityinsights.KindAnalyticsRule),
		// 				ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName1"),
		// 				Source: &armsecurityinsights.MetadataSource{
		// 					Name: to.Ptr("Contoso Solution 1.0"),
		// 					Kind: to.Ptr(armsecurityinsights.SourceKindSolution),
		// 					SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
		// 				},
		// 				Version: to.Ptr("1.0.0.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("metadataName2"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
		// 			ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName2"),
		// 			Properties: &armsecurityinsights.MetadataProperties{
		// 				ContentID: to.Ptr("f5160682-0e10-4e23-8fcf-df3df49c5522"),
		// 				Kind: to.Ptr(armsecurityinsights.KindAnalyticsRule),
		// 				ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName2"),
		// 				Source: &armsecurityinsights.MetadataSource{
		// 					Name: to.Ptr("Contoso Solution 1.0"),
		// 					Kind: to.Ptr(armsecurityinsights.SourceKindSolution),
		// 					SourceID: to.Ptr("b688a130-76f4-4a07-bf57-762222a3cadf"),
		// 				},
		// 				Version: to.Ptr("1.0.0.0"),
		// 			},
		// 	}},
		// }
	}
}
