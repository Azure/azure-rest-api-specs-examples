package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Job/CommandJob/createOrUpdate.json
func ExampleJobsClient_CreateOrUpdate_createOrUpdateCommandJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().CreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearning.JobBase{
		Properties: &armmachinelearning.CommandJob{
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
			JobType: to.Ptr(armmachinelearning.JobTypeCommand),
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
			CodeID:  to.Ptr("string"),
			Command: to.Ptr("string"),
			Distribution: &armmachinelearning.TensorFlow{
				DistributionType:     to.Ptr(armmachinelearning.DistributionTypeTensorFlow),
				ParameterServerCount: to.Ptr[int32](1),
				WorkerCount:          to.Ptr[int32](1),
			},
			EnvironmentID: to.Ptr("string"),
			EnvironmentVariables: map[string]*string{
				"string": to.Ptr("string"),
			},
			Inputs: map[string]armmachinelearning.JobInputClassification{
				"string": &armmachinelearning.LiteralJobInput{
					Description:  to.Ptr("string"),
					JobInputType: to.Ptr(armmachinelearning.JobInputTypeLiteral),
					Value:        to.Ptr("string"),
				},
			},
			Limits: &armmachinelearning.CommandJobLimits{
				JobLimitsType: to.Ptr(armmachinelearning.JobLimitsTypeCommand),
				Timeout:       to.Ptr("PT5M"),
			},
			Outputs: map[string]armmachinelearning.JobOutputClassification{
				"string": &armmachinelearning.URIFileJobOutput{
					Mode:          to.Ptr(armmachinelearning.OutputDeliveryModeReadWriteMount),
					URI:           to.Ptr("string"),
					Description:   to.Ptr("string"),
					JobOutputType: to.Ptr(armmachinelearning.JobOutputTypeURIFile),
				},
			},
			Resources: &armmachinelearning.JobResourceConfiguration{
				InstanceCount: to.Ptr[int32](1),
				InstanceType:  to.Ptr("string"),
				Properties: map[string]any{
					"string": map[string]any{
						"e6b6493e-7d5e-4db3-be1e-306ec641327e": nil,
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobBase = armmachinelearning.JobBase{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearning.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmachinelearning.CommandJob{
	// 		Description: to.Ptr("string"),
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		ComputeID: to.Ptr("string"),
	// 		DisplayName: to.Ptr("string"),
	// 		ExperimentName: to.Ptr("string"),
	// 		Identity: &armmachinelearning.AmlToken{
	// 			IdentityType: to.Ptr(armmachinelearning.IdentityConfigurationTypeAMLToken),
	// 		},
	// 		JobType: to.Ptr(armmachinelearning.JobTypeCommand),
	// 		Services: map[string]*armmachinelearning.JobService{
	// 			"string": &armmachinelearning.JobService{
	// 				Endpoint: to.Ptr("string"),
	// 				ErrorMessage: to.Ptr("string"),
	// 				JobServiceType: to.Ptr("string"),
	// 				Port: to.Ptr[int32](1),
	// 				Properties: map[string]*string{
	// 					"string": to.Ptr("string"),
	// 				},
	// 				Status: to.Ptr("string"),
	// 			},
	// 		},
	// 		Status: to.Ptr(armmachinelearning.JobStatusNotStarted),
	// 		CodeID: to.Ptr("string"),
	// 		Command: to.Ptr("string"),
	// 		Distribution: &armmachinelearning.TensorFlow{
	// 			DistributionType: to.Ptr(armmachinelearning.DistributionTypeTensorFlow),
	// 			ParameterServerCount: to.Ptr[int32](1),
	// 			WorkerCount: to.Ptr[int32](1),
	// 		},
	// 		EnvironmentID: to.Ptr("string"),
	// 		EnvironmentVariables: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Inputs: map[string]armmachinelearning.JobInputClassification{
	// 			"string": &armmachinelearning.LiteralJobInput{
	// 				Description: to.Ptr("string"),
	// 				JobInputType: to.Ptr(armmachinelearning.JobInputTypeLiteral),
	// 				Value: to.Ptr("string"),
	// 			},
	// 		},
	// 		Limits: &armmachinelearning.CommandJobLimits{
	// 			JobLimitsType: to.Ptr(armmachinelearning.JobLimitsTypeCommand),
	// 			Timeout: to.Ptr("PT5M"),
	// 		},
	// 		Outputs: map[string]armmachinelearning.JobOutputClassification{
	// 			"string": &armmachinelearning.URIFileJobOutput{
	// 				Mode: to.Ptr(armmachinelearning.OutputDeliveryModeReadWriteMount),
	// 				URI: to.Ptr("string"),
	// 				Description: to.Ptr("string"),
	// 				JobOutputType: to.Ptr(armmachinelearning.JobOutputTypeURIFile),
	// 			},
	// 		},
	// 		Parameters: map[string]any{
	// 			"string": "string",
	// 		},
	// 		Resources: &armmachinelearning.JobResourceConfiguration{
	// 			InstanceCount: to.Ptr[int32](1),
	// 			InstanceType: to.Ptr("string"),
	// 			Properties: map[string]any{
	// 				"string": map[string]any{
	// 					"a0847709-f5aa-4561-8ba5-d915d403fdcf": nil,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
