package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: 2025-01-02/AlertRuleResources_Get_MaximumSet_Gen.json
func ExampleAlertRuleResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("A76F9850-996B-40B3-94D4-C98110A0EEC9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertRuleResourcesClient().Get(ctx, "rgWatcher", "testWatcher", "testAlert", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.AlertRuleResourcesClientGetResponse{
	// 	AlertRuleResource: &armdatabasewatcher.AlertRuleResource{
	// 		Properties: &armdatabasewatcher.AlertRuleResourceProperties{
	// 			AlertRuleResourceID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878Be/resourceGroups/rgWatcher/providers/microsoft.insights/scheduledqueryrules/alerts-demo"),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.798Z"); return t}()),
	// 			AlertRuleTemplateVersion: to.Ptr("1.0"),
	// 			AlertRuleTemplateID: to.Ptr("someTemplateId"),
	// 			CreatedWithProperties: to.Ptr(armdatabasewatcher.AlertRuleCreationPropertiesCreatedWithActionGroup),
	// 			ProvisioningState: to.Ptr(armdatabasewatcher.ResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878Be/resourceGroups/rgWatcher/providers/microsoft.databasewatcher/watchers/testWatcher/alertRuleResources/testAlert"),
	// 		Name: to.Ptr("testAlert"),
	// 		Type: to.Ptr("microsoft.databasewatcher/watchers/alertRuleResources"),
	// 		SystemData: &armdatabasewatcher.SystemData{
	// 			CreatedBy: to.Ptr("enbpvlpqbwd"),
	// 			CreatedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("mxp"),
	// 			LastModifiedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 		},
	// 	},
	// }
}
