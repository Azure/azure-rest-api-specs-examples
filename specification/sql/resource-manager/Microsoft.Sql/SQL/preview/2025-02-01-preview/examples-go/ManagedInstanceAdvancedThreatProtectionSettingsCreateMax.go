package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/ManagedInstanceAdvancedThreatProtectionSettingsCreateMax.json
func ExampleManagedInstanceAdvancedThreatProtectionSettingsClient_BeginCreateOrUpdate_updateAManagedInstanceSAdvancedThreatProtectionSettingsWithAllParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstanceAdvancedThreatProtectionSettingsClient().BeginCreateOrUpdate(ctx, "threatprotection-4799", "threatprotection-6440", armsql.AdvancedThreatProtectionNameDefault, armsql.ManagedInstanceAdvancedThreatProtection{
		Properties: &armsql.AdvancedThreatProtectionProperties{
			State: to.Ptr(armsql.AdvancedThreatProtectionStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.ManagedInstanceAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse{
	// 	ManagedInstanceAdvancedThreatProtection: armsql.ManagedInstanceAdvancedThreatProtection{
	// 		Name: to.Ptr("Default"),
	// 		Type: to.Ptr("Microsoft.Sql/managedInstances/advancedThreatProtectionSettings"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/threatprotection-4799/providers/Microsoft.Sql/managedInstances/threatprotection-6440/advancedThreatProtectionSettings/Default"),
	// 		Properties: &armsql.AdvancedThreatProtectionProperties{
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 			State: to.Ptr(armsql.AdvancedThreatProtectionStateEnabled),
	// 		},
	// 		SystemData: &armsql.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
