package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/StartStopManagedInstanceScheduleCreateOrUpdateMax.json
func ExampleStartStopManagedInstanceSchedulesClient_CreateOrUpdate_createsOrUpdatesTheManagedInstancesStartStopScheduleWithAllOptionalParametersSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStartStopManagedInstanceSchedulesClient().CreateOrUpdate(ctx, "schedulerg", "schedulemi", armsql.StartStopScheduleNameDefault, armsql.StartStopManagedInstanceSchedule{
		Properties: &armsql.StartStopManagedInstanceScheduleProperties{
			Description: to.Ptr("This is a schedule for our Dev/Test environment."),
			ScheduleList: []*armsql.ScheduleItem{
				{
					StartDay:  to.Ptr(armsql.DayOfWeekThursday),
					StartTime: to.Ptr("18:00"),
					StopDay:   to.Ptr(armsql.DayOfWeekThursday),
					StopTime:  to.Ptr("17:00"),
				},
				{
					StartDay:  to.Ptr(armsql.DayOfWeekThursday),
					StartTime: to.Ptr("15:00"),
					StopDay:   to.Ptr(armsql.DayOfWeekThursday),
					StopTime:  to.Ptr("14:00"),
				}},
			TimeZoneID: to.Ptr("Central European Standard Time"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StartStopManagedInstanceSchedule = armsql.StartStopManagedInstanceSchedule{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/startStopSchedules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/schedulerg/providers/Microsoft.Sql/managedInstances/schedulemi/startStopSchedules/default"),
	// 	Properties: &armsql.StartStopManagedInstanceScheduleProperties{
	// 		Description: to.Ptr("This is a schedule for our Dev/Test environment."),
	// 		NextExecutionTime: to.Ptr("2021-08-26T14:00:00"),
	// 		NextRunAction: to.Ptr("Stop"),
	// 		ScheduleList: []*armsql.ScheduleItem{
	// 			{
	// 				StartDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StartTime: to.Ptr("06:00 PM"),
	// 				StopDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StopTime: to.Ptr("05:00 PM"),
	// 			},
	// 			{
	// 				StartDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StartTime: to.Ptr("03:00 PM"),
	// 				StopDay: to.Ptr(armsql.DayOfWeekThursday),
	// 				StopTime: to.Ptr("02:00 PM"),
	// 		}},
	// 		TimeZoneID: to.Ptr("Central European Standard Time"),
	// 	},
	// 	SystemData: &armsql.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-26T04:41:33.937Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T04:41:33.937Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 	},
	// }
}
