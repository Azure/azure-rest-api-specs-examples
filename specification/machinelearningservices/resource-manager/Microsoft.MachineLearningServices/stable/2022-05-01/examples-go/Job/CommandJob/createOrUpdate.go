package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/Job/CommandJob/createOrUpdate.json
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
		armmachinelearning.JobBase{
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
				Resources: &armmachinelearning.ResourceConfiguration{
					InstanceCount: to.Ptr[int32](1),
					InstanceType:  to.Ptr("string"),
					Properties: map[string]interface{}{
						"string": map[string]interface{}{
							"e6b6493e-7d5e-4db3-be1e-306ec641327e": nil,
						},
					},
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
