package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Schedule/createOrUpdate.json
func ExampleSchedulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewSchedulesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearning.Schedule{
		Properties: &armmachinelearning.ScheduleProperties{
			Description: to.Ptr("string"),
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			Tags: map[string]*string{
				"string": to.Ptr("string"),
			},
			Action: &armmachinelearning.EndpointScheduleAction{
				ActionType: to.Ptr(armmachinelearning.ScheduleActionTypeInvokeBatchEndpoint),
				EndpointInvocationDefinition: map[string]interface{}{
					"9965593e-526f-4b89-bb36-761138cf2794": nil,
				},
			},
			DisplayName: to.Ptr("string"),
			IsEnabled:   to.Ptr(false),
			Trigger: &armmachinelearning.CronTrigger{
				EndTime:     to.Ptr("string"),
				StartTime:   to.Ptr("string"),
				TimeZone:    to.Ptr("string"),
				TriggerType: to.Ptr(armmachinelearning.TriggerTypeCron),
				Expression:  to.Ptr("string"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
