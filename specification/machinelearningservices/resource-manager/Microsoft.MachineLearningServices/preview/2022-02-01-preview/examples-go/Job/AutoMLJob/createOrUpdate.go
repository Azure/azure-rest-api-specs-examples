package armmachinelearning_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/AutoMLJob/createOrUpdate.json
func ExampleJobsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"test-rg",
		"my-aml-workspace",
		"string",
		armmachinelearning.JobBaseData{
			Properties: &armmachinelearning.AutoMLJob{
				Description: to.Ptr("string"),
				Properties: map[string]*string{
					"string": to.Ptr("string"),
				},
				Tags: map[string]*string{
					"string": to.Ptr("string"),
				},
				ComputeID:      to.Ptr("string"),
				DisplayName:    to.Ptr("string"),
				ExperimentName: to.Ptr("string"),
				Identity: &armmachinelearning.AmlToken{
					IdentityType: to.Ptr(armmachinelearning.IdentityConfigurationTypeAMLToken),
				},
				IsArchived: to.Ptr(false),
				JobType:    to.Ptr(armmachinelearning.JobTypeAutoML),
				Schedule: &armmachinelearning.CronSchedule{
					EndTime:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t }()),
					ScheduleStatus: to.Ptr(armmachinelearning.ScheduleStatusDisabled),
					ScheduleType:   to.Ptr(armmachinelearning.ScheduleTypeCron),
					StartTime:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t }()),
					TimeZone:       to.Ptr("string"),
					Expression:     to.Ptr("string"),
				},
				Services: map[string]*armmachinelearning.JobService{
					"string": {
						Endpoint:       to.Ptr("string"),
						JobServiceType: to.Ptr("string"),
						Port:           to.Ptr[int32](1),
						Properties: map[string]*string{
							"string": to.Ptr("string"),
						},
					},
				},
				EnvironmentID: to.Ptr("string"),
				EnvironmentVariables: map[string]*string{
					"string": to.Ptr("string"),
				},
				Outputs: map[string]armmachinelearning.JobOutputClassification{
					"string": &armmachinelearning.URIFileJobOutput{
						Mode:          to.Ptr(armmachinelearning.OutputDeliveryModeReadWriteMount),
						URI:           to.Ptr("string"),
						Description:   to.Ptr("string"),
						JobOutputType: to.Ptr(armmachinelearning.JobOutputTypeURIFile),
					},
				},
				Resources: &armmachinelearning.ResourceConfiguration{
					InstanceCount: to.Ptr[int32](1),
					InstanceType:  to.Ptr("string"),
					Properties: map[string]interface{}{
						"string": map[string]interface{}{
							"9bec0ab0-c62f-4fa9-a97c-7b24bbcc90ad": nil,
						},
					},
				},
				TaskDetails: &armmachinelearning.ImageClassification{
					TaskType: to.Ptr(armmachinelearning.TaskTypeImageClassification),
					DataSettings: &armmachinelearning.ImageVerticalDataSettings{
						TargetColumnName: to.Ptr("string"),
						TrainingData: &armmachinelearning.TrainingDataSettings{
							Data: &armmachinelearning.MLTableJobInput{
								URI:          to.Ptr("string"),
								JobInputType: to.Ptr(armmachinelearning.JobInputTypeMLTable),
							},
						},
					},
					LimitSettings: &armmachinelearning.ImageLimitSettings{
						MaxTrials: to.Ptr[int32](2),
					},
					ModelSettings: &armmachinelearning.ImageModelSettingsClassification{
						ValidationCropSize: to.Ptr[int32](2),
					},
					SearchSpace: []*armmachinelearning.ImageModelDistributionSettingsClassification{
						{
							ValidationCropSize: to.Ptr("choice(2, 360)"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
