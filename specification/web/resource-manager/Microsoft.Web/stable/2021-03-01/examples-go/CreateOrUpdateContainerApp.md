Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappservice%2Farmappservice%2Fv0.2.0/sdk/resourcemanager/appservice/armappservice/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateContainerApp.json
func ExampleContainerAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewContainerAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<name>",
		armappservice.ContainerApp{
			Kind:     to.StringPtr("<kind>"),
			Location: to.StringPtr("<location>"),
			Properties: &armappservice.ContainerAppProperties{
				Configuration: &armappservice.Configuration{
					Ingress: &armappservice.Ingress{
						External:   to.BoolPtr(true),
						TargetPort: to.Int32Ptr(3000),
					},
				},
				KubeEnvironmentID: to.StringPtr("<kube-environment-id>"),
				Template: &armappservice.Template{
					Containers: []*armappservice.Container{
						{
							Name:  to.StringPtr("<name>"),
							Image: to.StringPtr("<image>"),
						}},
					Dapr: &armappservice.Dapr{
						AppPort: to.Int32Ptr(3000),
						Enabled: to.BoolPtr(true),
					},
					Scale: &armappservice.Scale{
						MaxReplicas: to.Int32Ptr(5),
						MinReplicas: to.Int32Ptr(1),
						Rules: []*armappservice.ScaleRule{
							{
								Name: to.StringPtr("<name>"),
								Custom: &armappservice.CustomScaleRule{
									Type: to.StringPtr("<type>"),
									Metadata: map[string]*string{
										"concurrentRequests": to.StringPtr("50"),
									},
								},
							}},
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ContainerAppsClientCreateOrUpdateResult)
}
```
