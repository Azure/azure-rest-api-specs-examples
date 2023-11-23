package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopedResourceList.json
func ExamplePrivateLinkScopedResourcesClient_NewListByPrivateLinkScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkScopedResourcesClient().NewListByPrivateLinkScopePager("MyResourceGroup", "MyPrivateLinkScope", nil)
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
		// page.ScopedResourceListResult = armmonitor.ScopedResourceListResult{
		// 	Value: []*armmonitor.ScopedResource{
		// 		{
		// 			Name: to.Ptr("scoped-resource-name"),
		// 			Type: to.Ptr("Microsoft.Insights/privateLinkScopes/scopedResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/privateLinkScopes/MyPrivateLinkScope/scopedResources/scoped-resource-name"),
		// 			Properties: &armmonitor.ScopedResourceProperties{
		// 				LinkedResourceID: to.Ptr("/subscriptions/00000000-0000-2222-3333-444444444444/resourceGroups/MyComponentResourceGroup/providers/Microsoft.Insights/components/my-component"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 			SystemData: &armmonitor.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051Z"); return t}()),
		// 				CreatedBy: to.Ptr("bobby@contoso.com"),
		// 				CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("bobby@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("other-scoped-resource-name"),
		// 			Type: to.Ptr("Microsoft.Insights/privateLinkScopes/scopedResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/privateLinkScopes/MyPrivateLinkScope/scopedResources/other-scoped-resource-name"),
		// 			Properties: &armmonitor.ScopedResourceProperties{
		// 				LinkedResourceID: to.Ptr("/subscriptions/00000000-3333-2222-5555-444444444444/resourceGroups/MyWorkspaceResourceGroup/providers/Microsoft.OperationalInsights/workspaces/my-workspace"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 			SystemData: &armmonitor.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051Z"); return t}()),
		// 				CreatedBy: to.Ptr("bobby@contoso.com"),
		// 				CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("bobby@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
