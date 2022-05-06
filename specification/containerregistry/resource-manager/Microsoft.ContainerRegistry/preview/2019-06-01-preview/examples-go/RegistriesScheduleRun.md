Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerregistry%2Farmcontainerregistry%2Fv0.5.0/sdk/resourcemanager/containerregistry/armcontainerregistry/README.md) on how to add the SDK to your project and authenticate.

```go
package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RegistriesScheduleRun.json
func ExampleRegistriesClient_BeginScheduleRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewRegistriesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginScheduleRun(ctx,
		"<resource-group-name>",
		"<registry-name>",
		&armcontainerregistry.DockerBuildRequest{
			Type:             to.Ptr("<type>"),
			IsArchiveEnabled: to.Ptr(true),
			AgentConfiguration: &armcontainerregistry.AgentProperties{
				CPU: to.Ptr[int32](2),
			},
			Arguments: []*armcontainerregistry.Argument{
				{
					Name:     to.Ptr("<name>"),
					IsSecret: to.Ptr(false),
					Value:    to.Ptr("<value>"),
				},
				{
					Name:     to.Ptr("<name>"),
					IsSecret: to.Ptr(true),
					Value:    to.Ptr("<value>"),
				}},
			DockerFilePath: to.Ptr("<docker-file-path>"),
			ImageNames: []*string{
				to.Ptr("azurerest:testtag")},
			IsPushEnabled: to.Ptr(true),
			NoCache:       to.Ptr(true),
			Platform: &armcontainerregistry.PlatformProperties{
				Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
				OS:           to.Ptr(armcontainerregistry.OSLinux),
			},
			SourceLocation: to.Ptr("<source-location>"),
		},
		&armcontainerregistry.RegistriesClientBeginScheduleRunOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
