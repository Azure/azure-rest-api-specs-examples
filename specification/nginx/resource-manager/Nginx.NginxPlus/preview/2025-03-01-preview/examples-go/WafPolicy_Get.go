package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: 2025-03-01-preview/WafPolicy_Get.json
func ExampleWafPolicyClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWafPolicyClient().Get(ctx, "myResourceGroup", "myDeployment", "myWafPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnginx.WafPolicyClientGetResponse{
	// 	DeploymentWafPolicy: &armnginx.DeploymentWafPolicy{
	// 		Name: to.Ptr("myWafPolicy"),
	// 		Type: to.Ptr("Nginx.NginxPlus/nginxDeployments/wafPolicies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/wafPolicies/myWafPolicy"),
	// 		Properties: &armnginx.DeploymentWafPolicyProperties{
	// 			ApplyingState: &armnginx.DeploymentWafPolicyApplyingStatus{
	// 				Code: to.Ptr(armnginx.NginxDeploymentWafPolicyApplyingStatusCodeSucceeded),
	// 				DisplayStatus: to.Ptr("Policy Applied"),
	// 				Time: to.Ptr("2025-03-02T10:05:00.000Z"),
	// 			},
	// 			CompilingState: &armnginx.DeploymentWafPolicyCompilingStatus{
	// 				Code: to.Ptr(armnginx.NginxDeploymentWafPolicyCompilingStatusCodeSucceeded),
	// 				DisplayStatus: to.Ptr("Compiled Successfully"),
	// 				Time: to.Ptr("2025-03-02T10:00:00.000Z"),
	// 			},
	// 			Content: []byte("QUJDREVGR0g="),
	// 			Filepath: to.Ptr("/etc/nginx/waf/policy.conf"),
	// 			ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
