package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Job/CommandJob/createOrUpdateRayDistribution.json
func ExampleJobsClient_CreateOrUpdate_createOrUpdateCommandJobWithRayDistribution() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().CreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearning.JobBase{
		Properties: &armmachinelearning.CommandJob{
			Description: to.Ptr("Ray distributed command job"),
			CodeID:      to.Ptr("string"),
			Command:     to.Ptr("python train.py"),
			ComputeID:   to.Ptr("string"),
			DisplayName: to.Ptr("ray-multi-node-job"),
			Distribution: &armmachinelearning.Ray{
				DistributionType: to.Ptr(armmachinelearning.DistributionTypeRay),
				Port:             to.Ptr[int32](6379),
				IncludeDashboard: to.Ptr(true),
				DashboardPort:    to.Ptr[int32](8265),
			},
			EnvironmentID: to.Ptr("string"),
			EnvironmentVariables: map[string]*string{
				"string": to.Ptr("string"),
			},
			ExperimentName: to.Ptr("ray-experiment"),
			Identity: &armmachinelearning.AmlToken{
				IdentityType: to.Ptr(armmachinelearning.IdentityConfigurationTypeAMLToken),
			},
			JobType: to.Ptr(armmachinelearning.JobTypeCommand),
			Resources: &armmachinelearning.JobResourceConfiguration{
				InstanceCount: to.Ptr[int32](2),
				InstanceType:  to.Ptr("string"),
				Properties:    map[string]any{},
			},
			Tags: map[string]*string{
				"string": to.Ptr("string"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.JobsClientCreateOrUpdateResponse{
	// 	JobBase: armmachinelearning.JobBase{
	// 		Name: to.Ptr("string"),
	// 		ID: to.Ptr("string"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/jobs"),
	// 		Properties: &armmachinelearning.CommandJob{
	// 			Description: to.Ptr("Ray distributed command job"),
	// 			CodeID: to.Ptr("string"),
	// 			Command: to.Ptr("python train.py"),
	// 			ComputeID: to.Ptr("string"),
	// 			DisplayName: to.Ptr("ray-multi-node-job"),
	// 			Distribution: &armmachinelearning.Ray{
	// 				DistributionType: to.Ptr(armmachinelearning.DistributionTypeRay),
	// 				Port: to.Ptr[int32](6379),
	// 				Address: nil,
	// 				IncludeDashboard: to.Ptr(true),
	// 				DashboardPort: to.Ptr[int32](8265),
	// 				HeadNodeAdditionalArgs: nil,
	// 				WorkerNodeAdditionalArgs: nil,
	// 			},
	// 			EnvironmentID: to.Ptr("string"),
	// 			EnvironmentVariables: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 			ExperimentName: to.Ptr("ray-experiment"),
	// 			Identity: &armmachinelearning.AmlToken{
	// 				IdentityType: to.Ptr(armmachinelearning.IdentityConfigurationTypeAMLToken),
	// 			},
	// 			JobType: to.Ptr(armmachinelearning.JobTypeCommand),
	// 			Status: to.Ptr(armmachinelearning.JobStatusStarting),
	// 			Resources: &armmachinelearning.JobResourceConfiguration{
	// 				InstanceCount: to.Ptr[int32](2),
	// 				InstanceType: to.Ptr("string"),
	// 				Properties: map[string]any{
	// 				},
	// 			},
	// 			Tags: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-15T00:00:00.000Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
