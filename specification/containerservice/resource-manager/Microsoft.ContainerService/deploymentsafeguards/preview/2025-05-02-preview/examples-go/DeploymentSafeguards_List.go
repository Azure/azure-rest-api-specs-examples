package armdeploymentsafeguards_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armdeploymentsafeguards"
)

// Generated from example definition: 2025-05-02-preview/DeploymentSafeguards_List.json
func ExampleDeploymentSafeguardsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentsafeguards.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeploymentSafeguardsClient().NewListPager("subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1", nil)
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
		// page = armdeploymentsafeguards.DeploymentSafeguardsClientListResponse{
		// 	DeploymentSafeguardListResult: armdeploymentsafeguards.DeploymentSafeguardListResult{
		// 		Value: []*armdeploymentsafeguards.DeploymentSafeguard{
		// 			{
		// 				ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1/providers/Microsoft.ContainerService/deploymentSafeguards/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.ContainerService/deploymentSafeguards"),
		// 				SystemData: &armdeploymentsafeguards.SystemData{
		// 					CreatedBy: to.Ptr("someUser"),
		// 					CreatedByType: to.Ptr(armdeploymentsafeguards.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("someOtherUser"),
		// 					LastModifiedByType: to.Ptr(armdeploymentsafeguards.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 				},
		// 				ETag: to.Ptr("23ujdflewrj3="),
		// 				Properties: &armdeploymentsafeguards.Properties{
		// 					Level: to.Ptr(armdeploymentsafeguards.LevelWarn),
		// 					SystemExcludedNamespaces: []*string{
		// 						to.Ptr("kube-system"),
		// 						to.Ptr("gatekeeper-system"),
		// 					},
		// 					PodSecurityStandardsLevel: to.Ptr(armdeploymentsafeguards.PodSecurityStandardsLevelPodSecurityStandardsBaseline),
		// 					ProvisioningState: to.Ptr(armdeploymentsafeguards.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
