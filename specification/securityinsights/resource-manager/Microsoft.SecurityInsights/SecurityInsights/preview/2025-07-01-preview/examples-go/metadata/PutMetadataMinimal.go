package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/metadata/PutMetadataMinimal.json
func ExampleMetadataClient_Create_createUpdateMinimalMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetadataClient().Create(ctx, "myRg", "myWorkspace", "metadataName", armsecurityinsights.MetadataModel{
		Properties: &armsecurityinsights.MetadataProperties{
			ContentID: to.Ptr("c00ee137-7475-47c8-9cce-ec6f0f1bedd0"),
			Kind:      to.Ptr("AnalyticsRule"),
			ParentID:  to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.MetadataClientCreateResponse{
	// 	MetadataModel: armsecurityinsights.MetadataModel{
	// 		Name: to.Ptr("metadataName"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/metadata"),
	// 		ID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/metadata/metadataName"),
	// 		Properties: &armsecurityinsights.MetadataProperties{
	// 			Kind: to.Ptr("AnalyticsRule"),
	// 			ParentID: to.Ptr("/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName"),
	// 		},
	// 	},
	// }
}
