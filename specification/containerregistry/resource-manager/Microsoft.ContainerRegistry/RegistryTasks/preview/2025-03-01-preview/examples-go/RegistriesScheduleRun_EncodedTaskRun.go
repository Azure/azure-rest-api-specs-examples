package armcontainerregistrytasks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistrytasks"
)

// Generated from example definition: 2025-03-01-preview/RegistriesScheduleRun_EncodedTaskRun.json
func ExampleRegistriesClient_ScheduleRun_registriesScheduleRunEncodedTaskRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistrytasks.NewClientFactory("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegistriesClient().ScheduleRun(ctx, "myResourceGroup", "myRegistry", &armcontainerregistrytasks.EncodedTaskRunRequest{
		Type: to.Ptr("EncodedTaskRunRequest"),
		AgentConfiguration: &armcontainerregistrytasks.AgentProperties{
			CPU: to.Ptr[int32](2),
		},
		EncodedTaskContent:   to.Ptr("c3RlcHM6Cnt7IGlmIFZhbHVlcy5lbnZpcm9ubWVudCA9PSAncHJvZCcgfX0KICAtIHJ1bjogcHJvZCBzZXR1cAp7eyBlbHNlIGlmIFZhbHVlcy5lbnZpcm9ubWVudCA9PSAnc3RhZ2luZycgfX0KICAtIHJ1bjogc3RhZ2luZyBzZXR1cAp7eyBlbHNlIH19CiAgLSBydW46IGRlZmF1bHQgc2V0dXAKe3sgZW5kIH19CgogIC0gcnVuOiBidWlsZCAtdCBGYW5jeVRoaW5nOnt7LlZhbHVlcy5lbnZpcm9ubWVudH19LXt7LlZhbHVlcy52ZXJzaW9ufX0gLgoKcHVzaDogWydGYW5jeVRoaW5nOnt7LlZhbHVlcy5lbnZpcm9ubWVudH19LXt7LlZhbHVlcy52ZXJzaW9ufX0nXQ=="),
		EncodedValuesContent: to.Ptr("ZW52aXJvbm1lbnQ6IHByb2QKdmVyc2lvbjogMQ=="),
		Platform: &armcontainerregistrytasks.PlatformProperties{
			OS: to.Ptr(armcontainerregistrytasks.OSLinux),
		},
		Values: []*armcontainerregistrytasks.SetValue{
			{
				Name:     to.Ptr("mytestargument"),
				IsSecret: to.Ptr(false),
				Value:    to.Ptr("mytestvalue"),
			},
			{
				Name:     to.Ptr("mysecrettestargument"),
				IsSecret: to.Ptr(true),
				Value:    to.Ptr("mysecrettestvalue"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerregistrytasks.RegistriesClientScheduleRunResponse{
	// 	Run: &armcontainerregistrytasks.Run{
	// 		Name: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/run"),
	// 		ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 		Properties: &armcontainerregistrytasks.RunProperties{
	// 			LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.617Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armcontainerregistrytasks.ProvisioningStateSucceeded),
	// 			RunID: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 			Status: to.Ptr(armcontainerregistrytasks.RunStatusSucceeded),
	// 		},
	// 	},
	// }
}
