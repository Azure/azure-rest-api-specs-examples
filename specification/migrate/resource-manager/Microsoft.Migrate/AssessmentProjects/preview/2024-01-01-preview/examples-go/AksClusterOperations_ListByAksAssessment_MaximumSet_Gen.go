package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AksClusterOperations_ListByAksAssessment_MaximumSet_Gen.json
func ExampleAksClusterOperationsClient_NewListByAksAssessmentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAksClusterOperationsClient().NewListByAksAssessmentPager("rgaksswagger", "testproject", "testaksassessment", &armmigrationassessment.AksClusterOperationsClientListByAksAssessmentOptions{Filter: to.Ptr("sfmesuqkmjezolfjvalm")})
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
		// page.AKSClusterListResult = armmigrationassessment.AKSClusterListResult{
		// 	Value: []*armmigrationassessment.AKSCluster{
		// 		{
		// 			Name: to.Ptr("testaksassessment-cluster"),
		// 			Type: to.Ptr("AKSAssessmentCluster"),
		// 			ID: to.Ptr("/subscriptions/D6F60DF4-CE70-4E39-8217-B8FBE7CA85AA/resourceGroups/rgaksswagger/providers/Microsoft.Migrate/assessmentProjects/testproject/aksAssessments/testaksassessment/clusters/testaksassessment-cluster"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.108Z"); return t}()),
		// 				CreatedBy: to.Ptr("User"),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.108Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("User"),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			ETag: to.Ptr("00000000-0000-0000-a616-12d4724c01d9"),
		// 			Properties: &armmigrationassessment.AKSClusterProperties{
		// 				Name: to.Ptr("testaksassessment-cluster"),
		// 				MonthlyCost: to.Ptr[float32](14),
		// 				NodePools: []*armmigrationassessment.NodePool{
		// 					{
		// 						Name: to.Ptr("testaksassessmentCostDetail"),
		// 						ArmSKUName: to.Ptr("StandardDS_v2"),
		// 						ClusterName: to.Ptr("testaksassessment-cluster"),
		// 						ID: to.Ptr("testaksassessmentCostDetail"),
		// 						Mode: to.Ptr(armmigrationassessment.NodePoolModeUser),
		// 						MonthlyCost: to.Ptr[float32](10),
		// 						NodeCount: to.Ptr[int32](1),
		// 						OSType: to.Ptr(armmigrationassessment.OSTypeLinux),
		// 						PodApproxMonthlyCost: to.Ptr[float32](10),
		// 						PodCount: to.Ptr[int32](1),
		// 				}},
		// 				PodCount: to.Ptr[int32](26),
		// 				Region: to.Ptr("Unknown"),
		// 				SystemNodeCount: to.Ptr[int32](18),
		// 				SystemNodePoolCount: to.Ptr[int32](16),
		// 				UserNodeCount: to.Ptr[int32](6),
		// 				UserNodePoolCount: to.Ptr[int32](0),
		// 			},
		// 	}},
		// }
	}
}
