package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: 2025-03-01-preview/WafPolicy_List.json
func ExampleWafPolicyClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWafPolicyClient().NewListPager("myResourceGroup", "myDeployment", nil)
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
		// page = armnginx.WafPolicyClientListResponse{
		// 	DeploymentWafPolicyListResponse: armnginx.DeploymentWafPolicyListResponse{
		// 		NextLink: to.Ptr("https://management.azure.com/.../wafPolicies?api-version=2025-03-01-preview&$skiptoken=..."),
		// 		Value: []*armnginx.DeploymentWafPolicyMetadata{
		// 			{
		// 				Name: to.Ptr("wafPolicy1"),
		// 				Type: to.Ptr("Nginx.NginxPlus/nginxDeployments/wafPolicies"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/wafPolicies/wafPolicy1"),
		// 				Properties: &armnginx.DeploymentWafPolicyMetadataProperties{
		// 					ApplyingState: &armnginx.DeploymentWafPolicyApplyingStatus{
		// 						Code: to.Ptr(armnginx.NginxDeploymentWafPolicyApplyingStatusCodeSucceeded),
		// 						DisplayStatus: to.Ptr("Policy Applied"),
		// 						Time: to.Ptr("2025-03-02T10:05:00.000Z"),
		// 					},
		// 					CompilingState: &armnginx.DeploymentWafPolicyCompilingStatus{
		// 						Code: to.Ptr(armnginx.NginxDeploymentWafPolicyCompilingStatusCodeSucceeded),
		// 						DisplayStatus: to.Ptr("Compiled Successfully"),
		// 						Time: to.Ptr("2025-03-02T10:00:00.000Z"),
		// 					},
		// 					Filepath: to.Ptr("/etc/nginx/waf/policy.conf"),
		// 					ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("wafPolicy2"),
		// 				Type: to.Ptr("Nginx.NginxPlus/nginxDeployments/wafPolicies"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/wafPolicies/wafPolicy2"),
		// 				Properties: &armnginx.DeploymentWafPolicyMetadataProperties{
		// 					ApplyingState: &armnginx.DeploymentWafPolicyApplyingStatus{
		// 						Code: to.Ptr(armnginx.NginxDeploymentWafPolicyApplyingStatusCodeSucceeded),
		// 						DisplayStatus: to.Ptr("Policy Applied"),
		// 						Time: to.Ptr("2025-03-02T10:05:00.000Z"),
		// 					},
		// 					CompilingState: &armnginx.DeploymentWafPolicyCompilingStatus{
		// 						Code: to.Ptr(armnginx.NginxDeploymentWafPolicyCompilingStatusCodeSucceeded),
		// 						DisplayStatus: to.Ptr("Compiled Successfully"),
		// 						Time: to.Ptr("2025-03-02T10:00:00.000Z"),
		// 					},
		// 					Filepath: to.Ptr("/etc/nginx/waf/policy.conf"),
		// 					ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
