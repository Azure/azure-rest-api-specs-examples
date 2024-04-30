package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/AdvancedThreatProtectionSettingsPutEnabled.json
func ExampleAdvancedThreatProtectionSettingsClient_BeginUpdatePut_enableAServersAdvancedThreatProtectionSettingsUsingPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAdvancedThreatProtectionSettingsClient().BeginUpdatePut(ctx, "threatprotection-4799", "threatprotection-6440", armmysqlflexibleservers.AdvancedThreatProtectionNameDefault, armmysqlflexibleservers.AdvancedThreatProtection{
		Properties: &armmysqlflexibleservers.AdvancedThreatProtectionProperties{
			State: to.Ptr(armmysqlflexibleservers.AdvancedThreatProtectionStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AdvancedThreatProtection = armmysqlflexibleservers.AdvancedThreatProtection{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/advancedThreatProtectionSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/threatprotection-4799/providers/Microsoft.DBforMySQL/flexibleServers/threatprotection-6440/advancedThreatProtectionSettings/Default"),
	// 	SystemData: &armmysqlflexibleservers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmysqlflexibleservers.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmysqlflexibleservers.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmysqlflexibleservers.AdvancedThreatProtectionProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armmysqlflexibleservers.AdvancedThreatProtectionProvisioningStateSucceeded),
	// 		State: to.Ptr(armmysqlflexibleservers.AdvancedThreatProtectionStateEnabled),
	// 	},
	// }
}
