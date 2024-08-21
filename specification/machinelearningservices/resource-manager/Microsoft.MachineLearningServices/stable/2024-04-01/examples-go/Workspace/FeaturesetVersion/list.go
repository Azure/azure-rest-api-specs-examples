package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/FeaturesetVersion/list.json
func ExampleFeaturesetVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFeaturesetVersionsClient().NewListPager("test-rg", "my-aml-workspace", "string", &armmachinelearning.FeaturesetVersionsClientListOptions{Skip: nil,
		Tags:         to.Ptr("string"),
		ListViewType: to.Ptr(armmachinelearning.ListViewTypeAll),
		PageSize:     nil,
		VersionName:  nil,
		Version:      nil,
		Description:  nil,
		CreatedBy:    nil,
		Stage:        nil,
	})
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
		// page.FeaturesetVersionResourceArmPaginatedResult = armmachinelearning.FeaturesetVersionResourceArmPaginatedResult{
		// 	Value: []*armmachinelearning.FeaturesetVersion{
		// 		{
		// 			Name: to.Ptr("string"),
		// 			Type: to.Ptr("string"),
		// 			ID: to.Ptr("string"),
		// 			SystemData: &armmachinelearning.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T11:45:56.999Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T11:45:56.999Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeKey),
		// 			},
		// 			Properties: &armmachinelearning.FeaturesetVersionProperties{
		// 				Description: to.Ptr("string"),
		// 				Properties: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				IsAnonymous: to.Ptr(false),
		// 				IsArchived: to.Ptr(false),
		// 				Entities: []*string{
		// 					to.Ptr("string")},
		// 					MaterializationSettings: &armmachinelearning.MaterializationSettings{
		// 						Notification: &armmachinelearning.NotificationSetting{
		// 							EmailOn: []*armmachinelearning.EmailNotificationEnableType{
		// 								to.Ptr(armmachinelearning.EmailNotificationEnableTypeJobCompleted)},
		// 								Emails: []*string{
		// 									to.Ptr("string")},
		// 								},
		// 								Resource: &armmachinelearning.MaterializationComputeResource{
		// 									InstanceType: to.Ptr("string"),
		// 								},
		// 								Schedule: &armmachinelearning.RecurrenceTrigger{
		// 									EndTime: to.Ptr("string"),
		// 									StartTime: to.Ptr("string"),
		// 									TimeZone: to.Ptr("string"),
		// 									TriggerType: to.Ptr(armmachinelearning.TriggerTypeRecurrence),
		// 									Frequency: to.Ptr(armmachinelearning.RecurrenceFrequencyMonth),
		// 									Interval: to.Ptr[int32](1),
		// 									Schedule: &armmachinelearning.RecurrenceSchedule{
		// 										Hours: []*int32{
		// 											to.Ptr[int32](1)},
		// 											Minutes: []*int32{
		// 												to.Ptr[int32](1)},
		// 												MonthDays: []*int32{
		// 													to.Ptr[int32](1)},
		// 													WeekDays: []*armmachinelearning.WeekDay{
		// 														to.Ptr(armmachinelearning.WeekDaySaturday)},
		// 													},
		// 												},
		// 												SparkConfiguration: map[string]*string{
		// 													"string": to.Ptr("string"),
		// 												},
		// 												StoreType: to.Ptr(armmachinelearning.MaterializationStoreTypeOffline),
		// 											},
		// 											ProvisioningState: to.Ptr(armmachinelearning.AssetProvisioningStateSucceeded),
		// 											Specification: &armmachinelearning.FeaturesetSpecification{
		// 												Path: to.Ptr("string"),
		// 											},
		// 											Stage: to.Ptr("string"),
		// 										},
		// 								}},
		// 							}
	}
}
