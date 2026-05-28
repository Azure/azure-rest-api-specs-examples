package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2023-06-01-preview/PrivateLinkScopesCreate.json
func ExamplePrivateLinkScopesClient_CreateOrUpdate_privateLinkScopeCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("86dc51d3-92ed-4d7e-947a-775ea79b4919", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkScopesClient().CreateOrUpdate(ctx, "my-resource-group", "my-privatelinkscope", armmonitor.AzureMonitorPrivateLinkScope{
		Location: to.Ptr("Global"),
		Properties: &armmonitor.AzureMonitorPrivateLinkScopeProperties{
			AccessModeSettings: &armmonitor.AccessModeSettings{
				Exclusions:          []*armmonitor.AccessModeSettingsExclusion{},
				IngestionAccessMode: to.Ptr(armmonitor.AccessModeOpen),
				QueryAccessMode:     to.Ptr(armmonitor.AccessModeOpen),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.PrivateLinkScopesClientCreateOrUpdateResponse{
	// 	AzureMonitorPrivateLinkScope: armmonitor.AzureMonitorPrivateLinkScope{
	// 		Name: to.Ptr("my-privatelinkscope"),
	// 		Type: to.Ptr("Microsoft.Insights/privateLinkScopes"),
	// 		ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.insights/privateLinkScopes/my-privatelinkscope"),
	// 		Location: to.Ptr("Global"),
	// 		Properties: &armmonitor.AzureMonitorPrivateLinkScopeProperties{
	// 			AccessModeSettings: &armmonitor.AccessModeSettings{
	// 				Exclusions: []*armmonitor.AccessModeSettingsExclusion{
	// 				},
	// 				IngestionAccessMode: to.Ptr(armmonitor.AccessModeOpen),
	// 				QueryAccessMode: to.Ptr(armmonitor.AccessModeOpen),
	// 			},
	// 			ProvisioningState: to.Ptr(armmonitor.PrivateLinkScopeProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armmonitor.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
	// 			CreatedBy: to.Ptr("bobby@contoso.com"),
	// 			CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("bobby@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
