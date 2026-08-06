package armcontainerserviceaimanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerserviceaimanager/armcontainerserviceaimanager"
)

// Generated from example definition: 2026-05-02-preview/ModelDeployments_ListByAIManagerNamespace.json
func ExampleModelDeploymentsClient_NewListByAIManagerNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerserviceaimanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewModelDeploymentsClient().NewListByAIManagerNamespacePager("rgaimanagers", "aimanager1", "namespace-1", nil)
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
		// page = armcontainerserviceaimanager.ModelDeploymentsClientListByAIManagerNamespaceResponse{
		// 	ModelDeploymentListResult: armcontainerserviceaimanager.ModelDeploymentListResult{
		// 		Value: []*armcontainerserviceaimanager.ModelDeployment{
		// 			{
		// 				Properties: &armcontainerserviceaimanager.ModelDeploymentProperties{
		// 					ProvisioningState: to.Ptr(armcontainerserviceaimanager.ModelDeploymentProvisioningStateSucceeded),
		// 					ModelResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgaimanagers/providers/Microsoft.ContainerService/aiModels/9806f0c862fdd920"),
		// 					ModelSourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgaimanagers/providers/Microsoft.ContainerService/aiManagers/aimanager1/modelSources/huggingface"),
		// 					PerformanceMode: to.Ptr(armcontainerserviceaimanager.ModelDeploymentPerformanceModeBalanced),
		// 					VMSize: to.Ptr("Standard_NC96ads_A100_v4"),
		// 					Scale: &armcontainerserviceaimanager.ScalingProfile{
		// 						Autoscale: &armcontainerserviceaimanager.AutoscaleProfile{
		// 							MinReplicas: to.Ptr[int32](2),
		// 							MaxReplicas: to.Ptr[int32](8),
		// 						},
		// 					},
		// 					Status: &armcontainerserviceaimanager.ModelDeploymentStatus{
		// 						Endpoint: to.Ptr("https://team-alpha.aks-cluster.eastus.aksapp.io/v1"),
		// 						Engine: to.Ptr("vllm"),
		// 						EngineVersion: to.Ptr("0.20.0"),
		// 						MaxModelLen: to.Ptr[int32](32768),
		// 						Quantization: to.Ptr("fp16"),
		// 						DesiredReplicas: to.Ptr[int32](4),
		// 						CurrentReplicas: to.Ptr[int32](4),
		// 						PeakTokensPerMinute: to.Ptr[int32](12000),
		// 						EstimatedProvisionTimeSeconds: to.Ptr[int32](900),
		// 					},
		// 				},
		// 				ETag: to.Ptr("\"abc123def456\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgaimanagers/providers/Microsoft.ContainerService/aiManagers/aimanager1/namespaces/namespace-1/modelDeployments/deployment-1"),
		// 				Name: to.Ptr("deployment-1"),
		// 				Type: to.Ptr("Microsoft.ContainerService/aiManagers/namespaces/modelDeployments"),
		// 				SystemData: &armcontainerserviceaimanager.SystemData{
		// 					CreatedBy: to.Ptr("user@example.com"),
		// 					CreatedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@example.com"),
		// 					LastModifiedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgaimanagers/providers/Microsoft.ContainerService/aiManagers/aimanager1/namespaces/namespace-1/modelDeployments?api-version=2026-05-02-preview&$skiptoken=token"),
		// 	},
		// }
	}
}
