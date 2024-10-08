package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Quota/list.json
func ExampleQuotasClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQuotasClient().NewListPager("eastus", nil)
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
		// page.ListWorkspaceQuotas = armmachinelearning.ListWorkspaceQuotas{
		// 	Value: []*armmachinelearning.ResourceQuota{
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard D Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard D Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/quotas/Standard_D_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](48),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard D Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard D Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_D_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](12),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard D Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard D Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_D_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](12),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard D Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard D Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace3/quotas/Standard_D_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/quota"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](12),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard DSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace3/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](12),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/quotas/Standard_Dv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_Dv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_Dv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard Dv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace3/quotas/Standard_Dv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard FSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard FSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/quotas/Standard_FSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard FSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard FSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_FSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard FSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard FSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_FSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard FSv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard FSv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace3/quotas/Standard_FSv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](12),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/quotas/Standard_NC_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_NC_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_NC_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NC Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace3/quotas/Standard_NC_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NCv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NCv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/quotas/Standard_NCv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NCv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NCv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_NCv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NCv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NCv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_NCv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NCv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NCv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace3/quotas/Standard_NCv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NCv3 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NCv3 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/quotas/Standard_NCv3_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NCv3 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NCv3 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_NCv3_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NCv3 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NCv3 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_NCv3_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NCv3 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NCv3 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace3/quotas/Standard_NCv3_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard ND Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard ND Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/quotas/Standard_ND_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard ND Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard ND Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_ND_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard ND Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard ND Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_ND_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard ND Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard ND Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace3/quotas/Standard_ND_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NDv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NDv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/quotas/Standard_NDv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NDv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NDv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_NDv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NDv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NDv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_NDv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NDv2 Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NDv2 Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace3/quotas/Standard_NDv2_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](0),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NV Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NV Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/vmFamily/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/quotas/Standard_NV_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NV Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NV Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_NV_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NV Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NV Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_NV_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 		},
		// 		{
		// 			Name: &armmachinelearning.ResourceName{
		// 				LocalizedValue: to.Ptr("Standard NV Family Cluster Dedicated vCPUs"),
		// 				Value: to.Ptr("Standard NV Family Cluster Dedicated vCPUs"),
		// 			},
		// 			Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace3/quotas/Standard_NV_Family_Cluster_Dedicated_vCPUs"),
		// 			Limit: to.Ptr[int64](24),
		// 			Unit: to.Ptr(armmachinelearning.QuotaUnitCount),
		// 	}},
		// }
	}
}
