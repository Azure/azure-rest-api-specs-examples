package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Compute/listKeys.json
func ExampleComputeClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputeClient().ListKeys(ctx, "testrg123", "workspaces123", "compute123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.ComputeClientListKeysResponse{
	// 	                            ComputeSecretsClassification: &armmachinelearning.AksComputeSecrets{
	// 		AdminKubeConfig: to.Ptr("admin kube config..."),
	// 		ImagePullSecretName: to.Ptr("the image pull secret name"),
	// 		UserKubeConfig: to.Ptr("user kube config..."),
	// 		ComputeType: to.Ptr(armmachinelearning.ComputeTypeAKS),
	// 	},
	// 	                        }
}
