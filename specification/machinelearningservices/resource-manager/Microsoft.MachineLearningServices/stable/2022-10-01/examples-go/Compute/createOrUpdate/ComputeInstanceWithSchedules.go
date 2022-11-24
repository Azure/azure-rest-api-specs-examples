package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Compute/createOrUpdate/ComputeInstanceWithSchedules.json
func ExampleComputeClient_BeginCreateOrUpdate_createAnComputeInstanceComputeWithSchedules() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewComputeClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "testrg123", "workspaces123", "compute123", armmachinelearning.ComputeResource{
		Properties: &armmachinelearning.ComputeInstance{
			ComputeType: to.Ptr(armmachinelearning.ComputeTypeComputeInstance),
			Properties: &armmachinelearning.ComputeInstanceProperties{
				ApplicationSharingPolicy:         to.Ptr(armmachinelearning.ApplicationSharingPolicyPersonal),
				ComputeInstanceAuthorizationType: to.Ptr(armmachinelearning.ComputeInstanceAuthorizationTypePersonal),
				PersonalComputeInstanceSettings: &armmachinelearning.PersonalComputeInstanceSettings{
					AssignedUser: &armmachinelearning.AssignedUser{
						ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
						TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
					},
				},
				Schedules: &armmachinelearning.ComputeSchedules{
					ComputeStartStop: []*armmachinelearning.ComputeStartStopSchedule{
						{
							Action: to.Ptr(armmachinelearning.ComputePowerActionStop),
							Cron: &armmachinelearning.CronTrigger{
								StartTime:  to.Ptr("2021-04-23T01:30:00"),
								TimeZone:   to.Ptr("Pacific Standard Time"),
								Expression: to.Ptr("0 18 * * *"),
							},
							Status:      to.Ptr(armmachinelearning.ScheduleStatusEnabled),
							TriggerType: to.Ptr(armmachinelearning.TriggerTypeCron),
						}},
				},
				SSHSettings: &armmachinelearning.ComputeInstanceSSHSettings{
					SSHPublicAccess: to.Ptr(armmachinelearning.SSHPublicAccessDisabled),
				},
				VMSize: to.Ptr("STANDARD_NC6"),
			},
		},
		Location: to.Ptr("eastus"),
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
