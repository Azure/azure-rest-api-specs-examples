package armcontainerregistrytasks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistrytasks"
)

// Generated from example definition: 2025-03-01-preview/RegistriesScheduleRun_WithCustomCredentials.json
func ExampleRegistriesClient_ScheduleRun_registriesScheduleRunWithCustomCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistrytasks.NewClientFactory("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegistriesClient().ScheduleRun(ctx, "myResourceGroup", "myRegistry", &armcontainerregistrytasks.DockerBuildRequest{
		Type: to.Ptr("DockerBuildRequest"),
		AgentConfiguration: &armcontainerregistrytasks.AgentProperties{
			CPU: to.Ptr[int32](2),
		},
		Arguments: []*armcontainerregistrytasks.Argument{
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
		Credentials: &armcontainerregistrytasks.Credentials{
			CustomRegistries: map[string]*armcontainerregistrytasks.CustomRegistryCredentials{
				"myregistry.azurecr.io": {
					Password: &armcontainerregistrytasks.SecretObject{
						Type:  to.Ptr(armcontainerregistrytasks.SecretObjectTypeOpaque),
						Value: to.Ptr("***"),
					},
					UserName: &armcontainerregistrytasks.SecretObject{
						Type:  to.Ptr(armcontainerregistrytasks.SecretObjectTypeOpaque),
						Value: to.Ptr("reg1"),
					},
				},
				"myregistry2.azurecr.io": {
					Password: &armcontainerregistrytasks.SecretObject{
						Type:  to.Ptr(armcontainerregistrytasks.SecretObjectTypeOpaque),
						Value: to.Ptr("***"),
					},
					UserName: &armcontainerregistrytasks.SecretObject{
						Type:  to.Ptr(armcontainerregistrytasks.SecretObjectTypeOpaque),
						Value: to.Ptr("reg2"),
					},
				},
			},
			SourceRegistry: &armcontainerregistrytasks.SourceRegistryCredentials{
				LoginMode: to.Ptr(armcontainerregistrytasks.SourceRegistryLoginModeDefault),
			},
		},
		DockerFilePath: to.Ptr("DockerFile"),
		ImageNames: []*string{
			to.Ptr("azurerest:testtag"),
		},
		IsArchiveEnabled: to.Ptr(true),
		IsPushEnabled:    to.Ptr(true),
		NoCache:          to.Ptr(true),
		Platform: &armcontainerregistrytasks.PlatformProperties{
			Architecture: to.Ptr(armcontainerregistrytasks.ArchitectureAmd64),
			OS:           to.Ptr(armcontainerregistrytasks.OSLinux),
		},
		SourceLocation: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/source.zip?sv=2015-04-05&st=2015-04-29T22%3A18%3A26Z&se=2015-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=Z%2FRHIX5Xcg0Mq2rqI3OlWTjEg2tYkboXr1P9ZUXDtkk%3D"),
		Target:         to.Ptr("stage1"),
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
