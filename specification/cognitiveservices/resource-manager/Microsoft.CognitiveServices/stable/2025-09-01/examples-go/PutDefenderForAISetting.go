package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6bbb6daca7175b2cab69b20b2dd01094e3c6a534/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/PutDefenderForAISetting.json
func ExampleDefenderForAISettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefenderForAISettingsClient().CreateOrUpdate(ctx, "resourceGroupName", "accountName", "Default", armcognitiveservices.DefenderForAISetting{
		Properties: &armcognitiveservices.DefenderForAISettingProperties{
			State: to.Ptr(armcognitiveservices.DefenderForAISettingStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DefenderForAISetting = armcognitiveservices.DefenderForAISetting{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/defenderForAISettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/defenderForAISettings/Default"),
	// 	SystemData: &armcognitiveservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		CreatedBy: to.Ptr("xxx@microsoft.com"),
	// 		CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("xxx@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 	},
	// 	Etag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armcognitiveservices.DefenderForAISettingProperties{
	// 		State: to.Ptr(armcognitiveservices.DefenderForAISettingStateEnabled),
	// 	},
	// }
}
