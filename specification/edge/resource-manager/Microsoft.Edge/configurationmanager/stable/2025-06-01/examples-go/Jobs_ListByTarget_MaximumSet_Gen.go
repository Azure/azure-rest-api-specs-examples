package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/Jobs_ListByTarget_MaximumSet_Gen.json
func ExampleJobsClient_NewListByTargetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListByTargetPager("gt", nil)
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
		// page = armworkloadorchestration.JobsClientListByTargetResponse{
		// 	JobListResult: armworkloadorchestration.JobListResult{
		// 		Value: []*armworkloadorchestration.Job{
		// 			{
		// 				Properties: &armworkloadorchestration.JobProperties{
		// 					JobType: to.Ptr(armworkloadorchestration.JobTypeDeploy),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-28T15:10:05.470Z"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-28T15:10:05.470Z"); return t}()),
		// 					Status: to.Ptr(armworkloadorchestration.JobStatusInProgress),
		// 					JobParameter: &armworkloadorchestration.JobParameterBase{
		// 						JobType: to.Ptr(armworkloadorchestration.JobType("JobParameterBase")),
		// 					},
		// 					CorrelationID: to.Ptr("fsebqmeouxfucluqkaoyagtp"),
		// 					Steps: []*armworkloadorchestration.JobStep{
		// 						{
		// 							Name: to.Ptr("duezohitlpz"),
		// 							Status: to.Ptr(armworkloadorchestration.JobStatusInProgress),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-28T15:10:05.470Z"); return t}()),
		// 							EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-28T15:10:05.470Z"); return t}()),
		// 							Message: to.Ptr("lkodiekkbtqwuixxryezipwvvavfu"),
		// 							Statistics: &armworkloadorchestration.DeployJobStepStatistics{
		// 								TotalCount: to.Ptr[int32](18),
		// 								SuccessCount: to.Ptr[int32](25),
		// 								FailedCount: to.Ptr[int32](25),
		// 								StatisticsType: to.Ptr(armworkloadorchestration.JobTypeDeploy),
		// 							},
		// 							Steps: []*armworkloadorchestration.JobStep{
		// 							},
		// 							ErrorDetails: &armworkloadorchestration.ErrorDetail{
		// 								Code: to.Ptr("onimqg"),
		// 								Message: to.Ptr("mzrbqqcxcbajkv"),
		// 								Target: to.Ptr("cemnupikyh"),
		// 								Details: []*armworkloadorchestration.ErrorDetail{
		// 									{
		// 										Code: to.Ptr("rbclcfwozglwnucds"),
		// 										Message: to.Ptr("zxbxf"),
		// 										Target: to.Ptr("rt"),
		// 										Details: []*armworkloadorchestration.ErrorDetail{
		// 										},
		// 										AdditionalInfo: []*armworkloadorchestration.ErrorAdditionalInfo{
		// 											{
		// 												Type: to.Ptr("nuqfqjakzxpxdthqvxo"),
		// 												Info: map[string]any{
		// 												},
		// 											},
		// 										},
		// 									},
		// 								},
		// 								AdditionalInfo: []*armworkloadorchestration.ErrorAdditionalInfo{
		// 									{
		// 										Type: to.Ptr("pttb"),
		// 										Info: map[string]any{
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					TriggeredBy: to.Ptr("lquvgncryxca"),
		// 					ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
		// 					ErrorDetails: &armworkloadorchestration.ErrorDetail{
		// 						Code: to.Ptr("onimqg"),
		// 						Message: to.Ptr("mzrbqqcxcbajkv"),
		// 						Target: to.Ptr("cemnupikyh"),
		// 						Details: []*armworkloadorchestration.ErrorDetail{
		// 							{
		// 								Code: to.Ptr("rbclcfwozglwnucds"),
		// 								Message: to.Ptr("zxbxf"),
		// 								Target: to.Ptr("rt"),
		// 								Details: []*armworkloadorchestration.ErrorDetail{
		// 								},
		// 								AdditionalInfo: []*armworkloadorchestration.ErrorAdditionalInfo{
		// 									{
		// 										Type: to.Ptr("nuqfqjakzxpxdthqvxo"),
		// 										Info: map[string]any{
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						AdditionalInfo: []*armworkloadorchestration.ErrorAdditionalInfo{
		// 							{
		// 								Type: to.Ptr("pttb"),
		// 								Info: map[string]any{
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 				ETag: to.Ptr("edzixghgizgqmjccbgpvcewnd"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
		// 				Name: to.Ptr("lxwubwlp"),
		// 				Type: to.Ptr("enekwruxhkecxg"),
		// 				SystemData: &armworkloadorchestration.SystemData{
		// 					CreatedBy: to.Ptr("favedmahrbemfqzeuggazxzrvwugxw"),
		// 					CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-24T11:04:49.597Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("lywqfnyqrutroctdfbxzytel"),
		// 					LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-24T11:04:49.597Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
