Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappservice%2Farmappservice%2Fv0.4.0/sdk/resourcemanager/appservice/armappservice/README.md) on how to add the SDK to your project and authenticate.

```go
package armappservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateContainerApp.json
func ExampleContainerAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappservice.NewContainerAppsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.ContainerApp{
			Kind:     to.Ptr("<kind>"),
			Location: to.Ptr("<location>"),
			Properties: &armappservice.ContainerAppProperties{
				Configuration: &armappservice.Configuration{
					Ingress: &armappservice.Ingress{
						External:   to.Ptr(true),
						TargetPort: to.Ptr[int32](3000),
					},
				},
				KubeEnvironmentID: to.Ptr("<kube-environment-id>"),
				Template: &armappservice.Template{
					Containers: []*armappservice.Container{
						{
							Name:  to.Ptr("<name>"),
							Image: to.Ptr("<image>"),
						}},
					Dapr: &armappservice.Dapr{
						AppPort: to.Ptr[int32](3000),
						Enabled: to.Ptr(true),
					},
					Scale: &armappservice.Scale{
						MaxReplicas: to.Ptr[int32](5),
						MinReplicas: to.Ptr[int32](1),
						Rules: []*armappservice.ScaleRule{
							{
								Name: to.Ptr("<name>"),
								Custom: &armappservice.CustomScaleRule{
									Type: to.Ptr("<type>"),
									Metadata: map[string]*string{
										"concurrentRequests": to.Ptr("50"),
									},
								},
							}},
					},
				},
			},
		},
		&armappservice.ContainerAppsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
