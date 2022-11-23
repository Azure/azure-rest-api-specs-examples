package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Job/SweepJob/createOrUpdate.json
func ExampleJobsClient_CreateOrUpdate_createOrUpdateSweepJob() {
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
		Properties: &armmachinelearning.SweepJob{
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
			JobType:        to.Ptr(armmachinelearning.JobTypeSweep),
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
			EarlyTermination: &armmachinelearning.MedianStoppingPolicy{
				DelayEvaluation:    to.Ptr[int32](1),
				EvaluationInterval: to.Ptr[int32](1),
				PolicyType:         to.Ptr(armmachinelearning.EarlyTerminationPolicyTypeMedianStopping),
			},
			Limits: &armmachinelearning.SweepJobLimits{
				JobLimitsType:       to.Ptr(armmachinelearning.JobLimitsTypeSweep),
				MaxConcurrentTrials: to.Ptr[int32](1),
				MaxTotalTrials:      to.Ptr[int32](1),
				TrialTimeout:        to.Ptr("PT1S"),
			},
			Objective: &armmachinelearning.Objective{
				Goal:          to.Ptr(armmachinelearning.GoalMinimize),
				PrimaryMetric: to.Ptr("string"),
			},
			SamplingAlgorithm: &armmachinelearning.GridSamplingAlgorithm{
				SamplingAlgorithmType: to.Ptr(armmachinelearning.SamplingAlgorithmTypeGrid),
			},
			SearchSpace: map[string]interface{}{
				"string": map[string]interface{}{},
			},
			Trial: &armmachinelearning.TrialComponent{
				CodeID:  to.Ptr("string"),
				Command: to.Ptr("string"),
				Distribution: &armmachinelearning.Mpi{
					DistributionType:        to.Ptr(armmachinelearning.DistributionTypeMpi),
					ProcessCountPerInstance: to.Ptr[int32](1),
				},
				EnvironmentID: to.Ptr("string"),
				EnvironmentVariables: map[string]*string{
					"string": to.Ptr("string"),
				},
				Resources: &armmachinelearning.JobResourceConfiguration{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
