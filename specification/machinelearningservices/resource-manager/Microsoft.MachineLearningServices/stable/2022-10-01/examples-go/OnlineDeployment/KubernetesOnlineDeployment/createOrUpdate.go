package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aafb0944f7ab936e8cfbad8969bd5eb32263fb4f/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineDeployment/KubernetesOnlineDeployment/createOrUpdate.json
func ExampleOnlineDeploymentsClient_BeginCreateOrUpdate_createOrUpdateKubernetesOnlineDeployment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOnlineDeploymentsClient().BeginCreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", armmachinelearning.OnlineDeployment{
		Location: to.Ptr("string"),
		Tags:     map[string]*string{},
		Identity: &armmachinelearning.ManagedServiceIdentity{
			Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
			UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
				"string": {},
			},
		},
		Kind: to.Ptr("string"),
		Properties: &armmachinelearning.KubernetesOnlineDeployment{
			Description: to.Ptr("string"),
			CodeConfiguration: &armmachinelearning.CodeConfiguration{
				CodeID:        to.Ptr("string"),
				ScoringScript: to.Ptr("string"),
			},
			EnvironmentID: to.Ptr("string"),
			EnvironmentVariables: map[string]*string{
				"string": to.Ptr("string"),
			},
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			AppInsightsEnabled:  to.Ptr(false),
			EndpointComputeType: to.Ptr(armmachinelearning.EndpointComputeTypeKubernetes),
			InstanceType:        to.Ptr("string"),
			LivenessProbe: &armmachinelearning.ProbeSettings{
				FailureThreshold: to.Ptr[int32](1),
				InitialDelay:     to.Ptr("PT5M"),
				Period:           to.Ptr("PT5M"),
				SuccessThreshold: to.Ptr[int32](1),
				Timeout:          to.Ptr("PT5M"),
			},
			Model:          to.Ptr("string"),
			ModelMountPath: to.Ptr("string"),
			RequestSettings: &armmachinelearning.OnlineRequestSettings{
				MaxConcurrentRequestsPerInstance: to.Ptr[int32](1),
				MaxQueueWait:                     to.Ptr("PT5M"),
				RequestTimeout:                   to.Ptr("PT5M"),
			},
			ScaleSettings: &armmachinelearning.DefaultScaleSettings{
				ScaleType: to.Ptr(armmachinelearning.ScaleTypeDefault),
			},
			ContainerResourceRequirements: &armmachinelearning.ContainerResourceRequirements{
				ContainerResourceLimits: &armmachinelearning.ContainerResourceSettings{
					CPU:    to.Ptr("\"1\""),
					Gpu:    to.Ptr("\"1\""),
					Memory: to.Ptr("\"2Gi\""),
				},
				ContainerResourceRequests: &armmachinelearning.ContainerResourceSettings{
					CPU:    to.Ptr("\"1\""),
					Gpu:    to.Ptr("\"1\""),
					Memory: to.Ptr("\"2Gi\""),
				},
			},
		},
		SKU: &armmachinelearning.SKU{
			Name:     to.Ptr("string"),
			Capacity: to.Ptr[int32](1),
			Family:   to.Ptr("string"),
			Size:     to.Ptr("string"),
			Tier:     to.Ptr(armmachinelearning.SKUTierFree),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OnlineDeployment = armmachinelearning.OnlineDeployment{
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
	// 	Location: to.Ptr("string"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmachinelearning.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
	// 			"string": &armmachinelearning.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr("string"),
	// 	Properties: &armmachinelearning.KubernetesOnlineDeployment{
	// 		Description: to.Ptr("string"),
	// 		CodeConfiguration: &armmachinelearning.CodeConfiguration{
	// 			CodeID: to.Ptr("string"),
	// 			ScoringScript: to.Ptr("string"),
	// 		},
	// 		EnvironmentID: to.Ptr("string"),
	// 		EnvironmentVariables: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		AppInsightsEnabled: to.Ptr(false),
	// 		EndpointComputeType: to.Ptr(armmachinelearning.EndpointComputeTypeKubernetes),
	// 		InstanceType: to.Ptr("string"),
	// 		LivenessProbe: &armmachinelearning.ProbeSettings{
	// 			FailureThreshold: to.Ptr[int32](1),
	// 			InitialDelay: to.Ptr("PT5M"),
	// 			Period: to.Ptr("PT5M"),
	// 			SuccessThreshold: to.Ptr[int32](1),
	// 			Timeout: to.Ptr("PT5M"),
	// 		},
	// 		Model: to.Ptr("string"),
	// 		ModelMountPath: to.Ptr("string"),
	// 		ProvisioningState: to.Ptr(armmachinelearning.DeploymentProvisioningStateSucceeded),
	// 		RequestSettings: &armmachinelearning.OnlineRequestSettings{
	// 			MaxConcurrentRequestsPerInstance: to.Ptr[int32](1),
	// 			MaxQueueWait: to.Ptr("PT5M"),
	// 			RequestTimeout: to.Ptr("PT5M"),
	// 		},
	// 		ScaleSettings: &armmachinelearning.DefaultScaleSettings{
	// 			ScaleType: to.Ptr(armmachinelearning.ScaleTypeDefault),
	// 		},
	// 		ContainerResourceRequirements: &armmachinelearning.ContainerResourceRequirements{
	// 			ContainerResourceLimits: &armmachinelearning.ContainerResourceSettings{
	// 				CPU: to.Ptr("\"1\""),
	// 				Gpu: to.Ptr("\"1\""),
	// 				Memory: to.Ptr("\"2Gi\""),
	// 			},
	// 			ContainerResourceRequests: &armmachinelearning.ContainerResourceSettings{
	// 				CPU: to.Ptr("\"1\""),
	// 				Gpu: to.Ptr("\"1\""),
	// 				Memory: to.Ptr("\"2Gi\""),
	// 			},
	// 		},
	// 	},
	// 	SKU: &armmachinelearning.SKU{
	// 		Name: to.Ptr("string"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("string"),
	// 		Size: to.Ptr("string"),
	// 		Tier: to.Ptr(armmachinelearning.SKUTierFree),
	// 	},
	// }
}
