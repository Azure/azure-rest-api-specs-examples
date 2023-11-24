package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RegistriesScheduleRun_Task.json
func ExampleRegistriesClient_BeginScheduleRun_registriesScheduleRunTask() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegistriesClient().BeginScheduleRun(ctx, "myResourceGroup", "myRegistry", &armcontainerregistry.TaskRunRequest{
		Type: to.Ptr("TaskRunRequest"),
		OverrideTaskStepProperties: &armcontainerregistry.OverrideTaskStepProperties{
			Arguments: []*armcontainerregistry.Argument{
				{
					Name:     to.Ptr("mytestargument"),
					IsSecret: to.Ptr(false),
					Value:    to.Ptr("mytestvalue"),
				},
				{
					Name:     to.Ptr("mysecrettestargument"),
					IsSecret: to.Ptr(true),
					Value:    to.Ptr("mysecrettestvalue"),
				}},
			File:               to.Ptr("overriddenDockerfile"),
			Target:             to.Ptr("build"),
			UpdateTriggerToken: to.Ptr("aGVsbG8gd29ybGQ="),
			Values: []*armcontainerregistry.SetValue{
				{
					Name:     to.Ptr("mytestname"),
					IsSecret: to.Ptr(false),
					Value:    to.Ptr("mytestvalue"),
				},
				{
					Name:     to.Ptr("mysecrettestname"),
					IsSecret: to.Ptr(true),
					Value:    to.Ptr("mysecrettestvalue"),
				}},
		},
		TaskID: to.Ptr("myTask"),
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
	// res.Run = armcontainerregistry.Run{
	// 	Name: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/run"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 	Properties: &armcontainerregistry.RunProperties{
	// 		LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.617Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RunID: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 		Status: to.Ptr(armcontainerregistry.RunStatusSucceeded),
	// 	},
	// }
}
