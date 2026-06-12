package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Job/AutoMLJob/list.json
func ExampleJobsClient_NewListPager_listAutoMlJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListPager("test-rg", "my-aml-workspace", nil)
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
		// page = armmachinelearning.JobsClientListResponse{
		// 	JobBaseResourceArmPaginatedResult: armmachinelearning.JobBaseResourceArmPaginatedResult{
		// 		Value: []*armmachinelearning.JobBase{
		// 			{
		// 				Name: to.Ptr("string"),
		// 				Type: to.Ptr("string"),
		// 				ID: to.Ptr("string"),
		// 				Properties: &armmachinelearning.AutoMLJob{
		// 					Description: to.Ptr("string"),
		// 					ComputeID: to.Ptr("string"),
		// 					DisplayName: to.Ptr("string"),
		// 					EnvironmentID: to.Ptr("string"),
		// 					EnvironmentVariables: map[string]*string{
		// 						"string": to.Ptr("string"),
		// 					},
		// 					ExperimentName: to.Ptr("string"),
		// 					Identity: &armmachinelearning.AmlToken{
		// 						IdentityType: to.Ptr(armmachinelearning.IdentityConfigurationTypeAMLToken),
		// 					},
		// 					IsArchived: to.Ptr(false),
		// 					JobType: to.Ptr(armmachinelearning.JobTypeAutoML),
		// 					Outputs: map[string]armmachinelearning.JobOutputClassification{
		// 						"string": &armmachinelearning.URIFileJobOutput{
		// 							Description: to.Ptr("string"),
		// 							JobOutputType: to.Ptr(armmachinelearning.JobOutputTypeURIFile),
		// 							Mode: to.Ptr(armmachinelearning.OutputDeliveryModeReadWriteMount),
		// 							URI: to.Ptr("string"),
		// 						},
		// 					},
		// 					Properties: map[string]*string{
		// 						"string": to.Ptr("string"),
		// 					},
		// 					Resources: &armmachinelearning.JobResourceConfiguration{
		// 						InstanceCount: to.Ptr[int32](1),
		// 						InstanceType: to.Ptr("string"),
		// 						Properties: map[string]any{
		// 							"string": map[string]any{
		// 								"9bec0ab0-c62f-4fa9-a97c-7b24bbcc90ad": nil,
		// 							},
		// 						},
		// 					},
		// 					Services: map[string]*armmachinelearning.JobService{
		// 						"string": &armmachinelearning.JobService{
		// 							Endpoint: to.Ptr("string"),
		// 							ErrorMessage: to.Ptr("string"),
		// 							JobServiceType: to.Ptr("string"),
		// 							Port: to.Ptr[int32](1),
		// 							Properties: map[string]*string{
		// 								"string": to.Ptr("string"),
		// 							},
		// 							Status: to.Ptr("string"),
		// 						},
		// 					},
		// 					Status: to.Ptr(armmachinelearning.JobStatus("Scheduled")),
		// 					Tags: map[string]*string{
		// 						"string": to.Ptr("string"),
		// 					},
		// 					TaskDetails: &armmachinelearning.ImageClassification{
		// 						LimitSettings: &armmachinelearning.ImageLimitSettings{
		// 							MaxTrials: to.Ptr[int32](2),
		// 						},
		// 						ModelSettings: &armmachinelearning.ImageModelSettingsClassification{
		// 							ValidationCropSize: to.Ptr[int32](2),
		// 						},
		// 						SearchSpace: []*armmachinelearning.ImageModelDistributionSettingsClassification{
		// 							{
		// 								ValidationCropSize: to.Ptr("choice(2, 360)"),
		// 							},
		// 						},
		// 						TargetColumnName: to.Ptr("string"),
		// 						TaskType: to.Ptr(armmachinelearning.TaskTypeImageClassification),
		// 						TrainingData: &armmachinelearning.MLTableJobInput{
		// 							JobInputType: to.Ptr(armmachinelearning.JobInputTypeMltable),
		// 							URI: to.Ptr("string"),
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armmachinelearning.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
