package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopesCreate.json
func ExamplePrivateLinkScopesClient_CreateOrUpdate_privateLinkScopeCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.AzureMonitorPrivateLinkScope = armmonitor.AzureMonitorPrivateLinkScope{
	// 	Name: to.Ptr("my-privatelinkscope"),
	// 	Type: to.Ptr("Microsoft.Insights/privateLinkScopes"),
	// 	ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/microsoft.insights/privateLinkScopes/my-privatelinkscope"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmonitor.AzureMonitorPrivateLinkScopeProperties{
	// 		AccessModeSettings: &armmonitor.AccessModeSettings{
	// 			Exclusions: []*armmonitor.AccessModeSettingsExclusion{
	// 			},
	// 			IngestionAccessMode: to.Ptr(armmonitor.AccessModeOpen),
	// 			QueryAccessMode: to.Ptr(armmonitor.AccessModeOpen),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// 	SystemData: &armmonitor.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051Z"); return t}()),
	// 		CreatedBy: to.Ptr("bobby@contoso.com"),
	// 		CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bobby@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 	},
	// }
}
