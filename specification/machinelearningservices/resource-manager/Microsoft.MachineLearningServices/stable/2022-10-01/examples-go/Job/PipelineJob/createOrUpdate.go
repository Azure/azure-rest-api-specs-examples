package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Job/PipelineJob/createOrUpdate.json
func ExampleJobsClient_CreateOrUpdate_createOrUpdatePipelineJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearning.JobBase{
		Properties: &armmachinelearning.PipelineJob{
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
			JobType:        to.Ptr(armmachinelearning.JobTypePipeline),
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
			Inputs: map[string]armmachinelearning.JobInputClassification{
				"string": &armmachinelearning.LiteralJobInput{
					Description:  to.Ptr("string"),
					JobInputType: to.Ptr(armmachinelearning.JobInputTypeLiteral),
					Value:        to.Ptr("string"),
				},
			},
			Outputs: map[string]armmachinelearning.JobOutputClassification{
				"string": &armmachinelearning.URIFileJobOutput{
					Mode:          to.Ptr(armmachinelearning.OutputDeliveryModeUpload),
					URI:           to.Ptr("string"),
					Description:   to.Ptr("string"),
					JobOutputType: to.Ptr(armmachinelearning.JobOutputTypeURIFile),
				},
			},
			Settings: map[string]interface{}{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
