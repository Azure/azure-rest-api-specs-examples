package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/RecoveryJobs_Get_MaximumSet_Gen.json
func ExampleRecoveryJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoveryJobsClient().Get(ctx, "sampleServiceGroupName", "samplePlanName", "c56888ef-9ced-4001-a6d4-7145a0309bdb", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresiliencemanagement.RecoveryJobsClientGetResponse{
	// 	RecoveryJob: armresiliencemanagement.RecoveryJob{
	// 		Properties: &armresiliencemanagement.RecoveryJobProperties{
	// 			ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
	// 			ResourceID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName"),
	// 			Operation: to.Ptr("Recovery"),
	// 			Status: to.Ptr(armresiliencemanagement.JobStatusNotStarted),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-09T10:14:25.588Z"); return t}()),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-09T10:14:25.588Z"); return t}()),
	// 			Duration: to.Ptr("PT44M"),
	// 			JobType: to.Ptr(armresiliencemanagement.JobTypeRecoveryPlan),
	// 			TriggeredBy: to.Ptr(armresiliencemanagement.JobTriggeredBySystem),
	// 			JobExtendedInfo: &armresiliencemanagement.JobExtendedInfo{
	// 				TasksList: []*armresiliencemanagement.JobTaskDetail{
	// 					{
	// 						TaskID: to.Ptr("task1"),
	// 						TaskName: to.Ptr("SampleTask"),
	// 						Duration: to.Ptr("PT36M"),
	// 						LinkedJobIDs: []*string{
	// 							to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryJobs/dc6998ef-80ed-4001-b6d4-5325a0309beh"),
	// 						},
	// 						Status: to.Ptr(armresiliencemanagement.JobStatusNotStarted),
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-15T12:20:34.443Z"); return t}()),
	// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-15T12:20:34.443Z"); return t}()),
	// 						ErrorDetails: &armresiliencemanagement.JobErrorInfo{
	// 							ErrorCode: to.Ptr("qeh"),
	// 							ErrorMessage: to.Ptr("xgtvlqabfcgszwkmqzlegdwtr"),
	// 							Recommendations: []*string{
	// 								to.Ptr("ryfmivugajhhyl"),
	// 							},
	// 						},
	// 						UserComments: []*armresiliencemanagement.JobUserComment{
	// 							{
	// 								CommentType: to.Ptr(armresiliencemanagement.CommentTypeDescription),
	// 								CommentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-15T12:20:34.406Z"); return t}()),
	// 								Comments: to.Ptr("lmmybouankjq"),
	// 							},
	// 						},
	// 						SubTasksList: []*armresiliencemanagement.JobTaskDetail{
	// 							{
	// 								Status: to.Ptr(armresiliencemanagement.JobStatusNotStarted),
	// 								StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-15T12:20:34.406Z"); return t}()),
	// 								EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-15T12:20:34.406Z"); return t}()),
	// 								Duration: to.Ptr("PT50M"),
	// 								ErrorDetails: &armresiliencemanagement.JobErrorInfo{
	// 									ErrorCode: to.Ptr("qeh"),
	// 									ErrorMessage: to.Ptr("xgtvlqabfcgszwkmqzlegdwtr"),
	// 									Recommendations: []*string{
	// 										to.Ptr("ryfmivugajhhyl"),
	// 									},
	// 								},
	// 								TaskID: to.Ptr("rjiflmtiwlti"),
	// 								TaskName: to.Ptr("grdozhipbnimwbcbyb"),
	// 								LinkedJobIDs: []*string{
	// 									to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryJobs/dc6998ef-80ed-4001-b6d4-5325a0309beh"),
	// 								},
	// 								UserComments: []*armresiliencemanagement.JobUserComment{
	// 									{
	// 										CommentType: to.Ptr(armresiliencemanagement.CommentTypeDescription),
	// 										CommentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-15T12:20:34.406Z"); return t}()),
	// 										Comments: to.Ptr("lmmybouankjq"),
	// 									},
	// 								},
	// 								SubTasksList: []*armresiliencemanagement.JobTaskDetail{
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				DynamicErrorMessage: to.Ptr("Sample error message"),
	// 			},
	// 			ErrorDetails: &armresiliencemanagement.JobErrorInfo{
	// 				ErrorCode: to.Ptr("SampleErrorCode"),
	// 				ErrorMessage: to.Ptr("Sample error message"),
	// 				Recommendations: []*string{
	// 					to.Ptr("Sample recommendation"),
	// 				},
	// 			},
	// 			RetryDetails: []*armresiliencemanagement.JobRetryDetails{
	// 				{
	// 					Status: to.Ptr(armresiliencemanagement.JobStatusNotStarted),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-15T12:20:34.443Z"); return t}()),
	// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-15T12:20:34.443Z"); return t}()),
	// 					Duration: to.Ptr("PT50M"),
	// 					ErrorDetails: &armresiliencemanagement.JobErrorInfo{
	// 						ErrorCode: to.Ptr("qeh"),
	// 						ErrorMessage: to.Ptr("xgtvlqabfcgszwkmqzlegdwtr"),
	// 						Recommendations: []*string{
	// 							to.Ptr("ryfmivugajhhyl"),
	// 						},
	// 					},
	// 					RetryAttempt: to.Ptr[int32](24),
	// 					UserComments: []*armresiliencemanagement.JobUserComment{
	// 						{
	// 							CommentType: to.Ptr(armresiliencemanagement.CommentTypeDescription),
	// 							CommentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-15T12:21:00.000Z"); return t}()),
	// 							Comments: to.Ptr("Resuming after verification"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			UserComments: []*armresiliencemanagement.JobUserComment{
	// 				{
	// 					CommentType: to.Ptr(armresiliencemanagement.CommentTypeDescription),
	// 					CommentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-02-16T11:28:56.176Z"); return t}()),
	// 					Comments: to.Ptr("User comments"),
	// 				},
	// 			},
	// 			ExecutionConfigurations: &armresiliencemanagement.ExecutionConfigurations{
	// 				UserConsent: to.Ptr(armresiliencemanagement.UserConsentUnspecified),
	// 			},
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/recoveryPlans/samplePlanName/recoveryJobs/c56888ef-9ced-4001-a6d4-7145a0309bdb"),
	// 		Name: to.Ptr("c56888ef-9ced-4001-a6d4-7145a0309bdb"),
	// 		Type: to.Ptr("Microsoft.AzureResilienceManagement/recoveryPlans/recoveryJobs"),
	// 		SystemData: &armresiliencemanagement.SystemData{
	// 			CreatedBy: to.Ptr("sampleUser"),
	// 			CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.796Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("sampleUser"),
	// 			LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.797Z"); return t}()),
	// 		},
	// 	},
	// }
}
