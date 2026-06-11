package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Schedule/get.json
func ExampleSchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSchedulesClient().Get(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.SchedulesClientGetResponse{
	// 	Schedule: armmachinelearning.Schedule{
	// 		Name: to.Ptr("string"),
	// 		Type: to.Ptr("string"),
	// 		ID: to.Ptr("string"),
	// 		Properties: &armmachinelearning.ScheduleProperties{
	// 			Description: to.Ptr("string"),
	// 			Action: &armmachinelearning.EndpointScheduleAction{
	// 				ActionType: to.Ptr(armmachinelearning.ScheduleActionTypeInvokeBatchEndpoint),
	// 				EndpointInvocationDefinition: map[string]any{
	// 					"a108545b-def1-4c86-8e53-dbcb1de3a8bc": nil,
	// 				},
	// 			},
	// 			DisplayName: to.Ptr("string"),
	// 			IsEnabled: to.Ptr(false),
	// 			Properties: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 			ProvisioningState: to.Ptr(armmachinelearning.ScheduleProvisioningStatusCreating),
	// 			Tags: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 			Trigger: &armmachinelearning.CronTrigger{
	// 				EndTime: to.Ptr("string"),
	// 				Expression: to.Ptr("string"),
	// 				StartTime: to.Ptr("string"),
	// 				TimeZone: to.Ptr("string"),
	// 				TriggerType: to.Ptr(armmachinelearning.TriggerTypeCron),
	// 			},
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeKey),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeKey),
	// 		},
	// 	},
	// }
}
