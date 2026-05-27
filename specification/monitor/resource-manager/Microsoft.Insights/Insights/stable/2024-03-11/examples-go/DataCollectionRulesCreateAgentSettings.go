package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-03-11/DataCollectionRulesCreateAgentSettings.json
func ExampleDataCollectionRulesClient_Create_createOrUpdateAnAgentSettingsConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("703362b3-f278-4e4b-9179-c76eaf41ffc2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataCollectionRulesClient().Create(ctx, "myResourceGroup", "myCollectionRule", armmonitor.DataCollectionRuleResource{
		Kind:     to.Ptr(armmonitor.KnownDataCollectionRuleResourceKind("AgentSettings")),
		Location: to.Ptr("eastus"),
		Properties: &armmonitor.DataCollectionRuleResourceProperties{
			Description: to.Ptr("An agent settings configuration"),
			AgentSettings: &armmonitor.DataCollectionRuleAgentSettings{
				Logs: []*armmonitor.AgentSetting{
					{
						Name:  to.Ptr(armmonitor.KnownAgentSettingNameMaxDiskQuotaInMB),
						Value: to.Ptr("5000"),
					},
					{
						Name:  to.Ptr(armmonitor.KnownAgentSettingNameUseTimeReceivedForForwardedEvents),
						Value: to.Ptr("1"),
					},
					{
						Name:  to.Ptr(armmonitor.KnownAgentSettingNameTags),
						Value: to.Ptr("Azure, Monitoring"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.DataCollectionRulesClientCreateResponse{
	// 	DataCollectionRuleResource: armmonitor.DataCollectionRuleResource{
	// 		Name: to.Ptr("myCollectionRule"),
	// 		Type: to.Ptr("Microsoft.Insights/dataCollectionRules"),
	// 		Etag: to.Ptr("070057da-0000-0000-0000-5ba70d6c0000"),
	// 		ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionRules/myCollectionRule"),
	// 		Kind: to.Ptr(armmonitor.KnownDataCollectionRuleResourceKind("AgentSettings")),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armmonitor.DataCollectionRuleResourceProperties{
	// 			Description: to.Ptr("An agent settings configuration"),
	// 			AgentSettings: &armmonitor.DataCollectionRuleAgentSettings{
	// 				Logs: []*armmonitor.AgentSetting{
	// 					{
	// 						Name: to.Ptr(armmonitor.KnownAgentSettingNameMaxDiskQuotaInMB),
	// 						Value: to.Ptr("5000"),
	// 					},
	// 					{
	// 						Name: to.Ptr(armmonitor.KnownAgentSettingNameUseTimeReceivedForForwardedEvents),
	// 						Value: to.Ptr("1"),
	// 					},
	// 					{
	// 						Name: to.Ptr(armmonitor.KnownAgentSettingNameTags),
	// 						Value: to.Ptr("Azure, Monitoring"),
	// 					},
	// 				},
	// 			},
	// 			ImmutableID: to.Ptr("dcr-76ce901eee3a400b9945b1e263a70000"),
	// 			ProvisioningState: to.Ptr(armmonitor.KnownDataCollectionRuleProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armmonitor.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-26T05:41:40.7885407Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-26T05:41:40.7885407Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user1"),
	// 			LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
