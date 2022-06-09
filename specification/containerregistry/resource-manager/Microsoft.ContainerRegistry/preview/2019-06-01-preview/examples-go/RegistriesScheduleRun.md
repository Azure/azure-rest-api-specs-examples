```go
package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RegistriesScheduleRun.json
func ExampleRegistriesClient_BeginScheduleRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewRegistriesClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginScheduleRun(ctx,
		"myResourceGroup",
		"myRegistry",
		&armcontainerregistry.DockerBuildRequest{
			Type:             to.Ptr("DockerBuildRequest"),
			IsArchiveEnabled: to.Ptr(true),
			AgentConfiguration: &armcontainerregistry.AgentProperties{
				CPU: to.Ptr[int32](2),
			},
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
			DockerFilePath: to.Ptr("DockerFile"),
			ImageNames: []*string{
				to.Ptr("azurerest:testtag")},
			IsPushEnabled: to.Ptr(true),
			NoCache:       to.Ptr(true),
			Platform: &armcontainerregistry.PlatformProperties{
				Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
				OS:           to.Ptr(armcontainerregistry.OSLinux),
			},
			SourceLocation: to.Ptr("https://myaccount.blob.core.windows.net/sascontainer/source.zip?sv=2015-04-05&st=2015-04-29T22%3A18%3A26Z&se=2015-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=Z%2FRHIX5Xcg0Mq2rqI3OlWTjEg2tYkboXr1P9ZUXDtkk%3D"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerregistry%2Farmcontainerregistry%2Fv0.6.0/sdk/resourcemanager/containerregistry/armcontainerregistry/README.md) on how to add the SDK to your project and authenticate.
