package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Workspace/FeaturesetVersion/list.json
func ExampleFeaturesetVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFeaturesetVersionsClient().NewListPager("test-rg", "my-aml-workspace", "string", &armmachinelearning.FeaturesetVersionsClientListOptions{
		ListViewType: to.Ptr(armmachinelearning.ListViewTypeAll),
		Tags:         to.Ptr("string")})
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
		// page = armmachinelearning.FeaturesetVersionsClientListResponse{
		// 	FeaturesetVersionResourceArmPaginatedResult: armmachinelearning.FeaturesetVersionResourceArmPaginatedResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/my-aml-workspace/featuresets/string/versions?api-version=2025-07-01-preview&$skip=2"),
		// 		Value: []*armmachinelearning.FeaturesetVersion{
		// 			{
		// 				Name: to.Ptr("string"),
		// 				Type: to.Ptr("string"),
		// 				ID: to.Ptr("string"),
		// 				Properties: &armmachinelearning.FeaturesetVersionProperties{
		// 					Description: to.Ptr("string"),
		// 					Entities: []*string{
		// 						to.Ptr("string"),
		// 					},
		// 					IsAnonymous: to.Ptr(false),
		// 					IsArchived: to.Ptr(false),
		// 					MaterializationSettings: &armmachinelearning.MaterializationSettings{
		// 						Notification: &armmachinelearning.NotificationSetting{
		// 							EmailOn: []*armmachinelearning.EmailNotificationEnableType{
		// 								to.Ptr(armmachinelearning.EmailNotificationEnableTypeJobCompleted),
		// 							},
		// 							Emails: []*string{
		// 								to.Ptr("string"),
		// 							},
		// 						},
		// 						Resource: &armmachinelearning.MaterializationComputeResource{
		// 							InstanceType: to.Ptr("string"),
		// 						},
		// 						Schedule: &armmachinelearning.RecurrenceTrigger{
		// 							EndTime: to.Ptr("string"),
		// 							Frequency: to.Ptr(armmachinelearning.RecurrenceFrequencyMonth),
		// 							Interval: to.Ptr[int32](1),
		// 							Schedule: &armmachinelearning.RecurrenceSchedule{
		// 								Hours: []*int32{
		// 									to.Ptr[int32](1),
		// 								},
		// 								Minutes: []*int32{
		// 									to.Ptr[int32](1),
		// 								},
		// 								MonthDays: []*int32{
		// 									to.Ptr[int32](1),
		// 								},
		// 								WeekDays: []*armmachinelearning.WeekDay{
		// 									to.Ptr(armmachinelearning.WeekDaySaturday),
		// 								},
		// 							},
		// 							StartTime: to.Ptr("string"),
		// 							TimeZone: to.Ptr("string"),
		// 							TriggerType: to.Ptr(armmachinelearning.TriggerTypeRecurrence),
		// 						},
		// 						SparkConfiguration: map[string]*string{
		// 							"string": to.Ptr("string"),
		// 						},
		// 						StoreType: to.Ptr(armmachinelearning.MaterializationStoreTypeOffline),
		// 					},
		// 					Properties: map[string]*string{
		// 						"string": to.Ptr("string"),
		// 					},
		// 					ProvisioningState: to.Ptr(armmachinelearning.AssetProvisioningStateCanceled),
		// 					Specification: &armmachinelearning.FeaturesetSpecification{
		// 						Path: to.Ptr("string"),
		// 					},
		// 					Stage: to.Ptr("string"),
		// 					Tags: map[string]*string{
		// 						"string": to.Ptr("string"),
		// 					},
		// 				},
		// 				SystemData: &armmachinelearning.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:49"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:49"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeKey),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
