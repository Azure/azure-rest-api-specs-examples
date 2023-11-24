package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aafb0944f7ab936e8cfbad8969bd5eb32263fb4f/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Compute/get/ComputeInstance.json
func ExampleComputeClient_Get_getAnComputeInstance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputeClient().Get(ctx, "testrg123", "workspaces123", "compute123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComputeResource = armmachinelearning.ComputeResource{
	// 	Properties: &armmachinelearning.ComputeInstance{
	// 		Description: to.Ptr("some compute"),
	// 		ComputeType: to.Ptr(armmachinelearning.ComputeTypeComputeInstance),
	// 		CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T22:00:00.000Z"); return t}()),
	// 		ModifiedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T22:00:00.000Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armmachinelearning.ProvisioningStateSucceeded),
	// 		Properties: &armmachinelearning.ComputeInstanceProperties{
	// 			ApplicationSharingPolicy: to.Ptr(armmachinelearning.ApplicationSharingPolicyShared),
	// 			Applications: []*armmachinelearning.ComputeInstanceApplication{
	// 				{
	// 					DisplayName: to.Ptr("Jupyter"),
	// 					EndpointURI: to.Ptr("https://compute123.eastus2.azureml.net/jupyter"),
	// 			}},
	// 			ComputeInstanceAuthorizationType: to.Ptr(armmachinelearning.ComputeInstanceAuthorizationTypePersonal),
	// 			ConnectivityEndpoints: &armmachinelearning.ComputeInstanceConnectivityEndpoints{
	// 				PrivateIPAddress: to.Ptr("10.0.0.1"),
	// 				PublicIPAddress: to.Ptr("10.0.0.1"),
	// 			},
	// 			CreatedBy: &armmachinelearning.ComputeInstanceCreatedBy{
	// 				UserID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				UserName: to.Ptr("foobar@microsoft.com"),
	// 				UserOrgID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 			PersonalComputeInstanceSettings: &armmachinelearning.PersonalComputeInstanceSettings{
	// 				AssignedUser: &armmachinelearning.AssignedUser{
	// 					ObjectID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 			SSHSettings: &armmachinelearning.ComputeInstanceSSHSettings{
	// 				AdminUserName: to.Ptr("azureuser"),
	// 				SSHPort: to.Ptr[int32](22),
	// 				SSHPublicAccess: to.Ptr(armmachinelearning.SSHPublicAccessEnabled),
	// 			},
	// 			State: to.Ptr(armmachinelearning.ComputeInstanceStateRunning),
	// 			Subnet: &armmachinelearning.ResourceID{
	// 				ID: to.Ptr("test-subnet-resource-id"),
	// 			},
	// 			VMSize: to.Ptr("STANDARD_NC6"),
	// 		},
	// 	},
	// 	Name: to.Ptr("compute123"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/computes"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspaces123/computes/compute123"),
	// 	Location: to.Ptr("eastus2"),
	// }
}
