package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Job/SweepJob/createOrUpdate.json
func ExampleJobsClient_CreateOrUpdate_createOrUpdateSweepJob() {
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
		Properties: &armmachinelearning.SweepJob{
			Description: to.Ptr("string"),
			ComputeID:   to.Ptr("string"),
			DisplayName: to.Ptr("string"),
			EarlyTermination: &armmachinelearning.MedianStoppingPolicy{
				DelayEvaluation:    to.Ptr[int32](1),
				EvaluationInterval: to.Ptr[int32](1),
				PolicyType:         to.Ptr(armmachinelearning.EarlyTerminationPolicyTypeMedianStopping),
			},
			ExperimentName: to.Ptr("string"),
			JobType:        to.Ptr(armmachinelearning.JobTypeSweep),
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
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			SamplingAlgorithm: &armmachinelearning.GridSamplingAlgorithm{
				SamplingAlgorithmType: to.Ptr(armmachinelearning.SamplingAlgorithmTypeGrid),
			},
			SearchSpace: map[string]any{
				"string": map[string]any{},
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
			Tags: map[string]*string{
				"string": to.Ptr("string"),
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
					Properties: map[string]any{
						"string": map[string]any{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.JobsClientCreateOrUpdateResponse{
	// 	JobBase: armmachinelearning.JobBase{
	// 		Name: to.Ptr("string"),
	// 		Type: to.Ptr("string"),
	// 		ID: to.Ptr("string"),
	// 		Properties: &armmachinelearning.SweepJob{
	// 			Description: to.Ptr("string"),
	// 			ComputeID: to.Ptr("string"),
	// 			DisplayName: to.Ptr("string"),
	// 			EarlyTermination: &armmachinelearning.MedianStoppingPolicy{
	// 				DelayEvaluation: to.Ptr[int32](1),
	// 				EvaluationInterval: to.Ptr[int32](1),
	// 				PolicyType: to.Ptr(armmachinelearning.EarlyTerminationPolicyTypeMedianStopping),
	// 			},
	// 			ExperimentName: to.Ptr("string"),
	// 			JobType: to.Ptr(armmachinelearning.JobTypeSweep),
	// 			Limits: &armmachinelearning.SweepJobLimits{
	// 				JobLimitsType: to.Ptr(armmachinelearning.JobLimitsTypeSweep),
	// 				MaxConcurrentTrials: to.Ptr[int32](1),
	// 				MaxTotalTrials: to.Ptr[int32](1),
	// 				TrialTimeout: to.Ptr("PT1S"),
	// 			},
	// 			Objective: &armmachinelearning.Objective{
	// 				Goal: to.Ptr(armmachinelearning.GoalMinimize),
	// 				PrimaryMetric: to.Ptr("string"),
	// 			},
	// 			Properties: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 			SamplingAlgorithm: &armmachinelearning.GridSamplingAlgorithm{
	// 				SamplingAlgorithmType: to.Ptr(armmachinelearning.SamplingAlgorithmTypeGrid),
	// 			},
	// 			SearchSpace: map[string]any{
	// 				"string": map[string]any{
	// 				},
	// 			},
	// 			Services: map[string]*armmachinelearning.JobService{
	// 				"string": &armmachinelearning.JobService{
	// 					Endpoint: to.Ptr("string"),
	// 					ErrorMessage: to.Ptr("string"),
	// 					JobServiceType: to.Ptr("string"),
	// 					Port: to.Ptr[int32](1),
	// 					Properties: map[string]*string{
	// 						"string": to.Ptr("string"),
	// 					},
	// 					Status: to.Ptr("string"),
	// 				},
	// 			},
	// 			Status: to.Ptr(armmachinelearning.JobStatusNotStarted),
	// 			Tags: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 			Trial: &armmachinelearning.TrialComponent{
	// 				CodeID: to.Ptr("string"),
	// 				Command: to.Ptr("string"),
	// 				Distribution: &armmachinelearning.Mpi{
	// 					DistributionType: to.Ptr(armmachinelearning.DistributionTypeMpi),
	// 					ProcessCountPerInstance: to.Ptr[int32](1),
	// 				},
	// 				EnvironmentID: to.Ptr("string"),
	// 				EnvironmentVariables: map[string]*string{
	// 					"string": to.Ptr("string"),
	// 				},
	// 				Resources: &armmachinelearning.JobResourceConfiguration{
	// 					InstanceCount: to.Ptr[int32](1),
	// 					InstanceType: to.Ptr("string"),
	// 					Properties: map[string]any{
	// 						"string": map[string]any{
	// 							"e6b6493e-7d5e-4db3-be1e-306ec641327e": nil,
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
