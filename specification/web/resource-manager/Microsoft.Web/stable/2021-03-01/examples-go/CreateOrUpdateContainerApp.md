```go
package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateContainerApp.json
func ExampleContainerAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewContainerAppsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg",
		"testcontainerApp0",
		armappservice.ContainerApp{
			Kind:     to.Ptr("containerApp"),
			Location: to.Ptr("East US"),
			Properties: &armappservice.ContainerAppProperties{
				Configuration: &armappservice.Configuration{
					Ingress: &armappservice.Ingress{
						External:   to.Ptr(true),
						TargetPort: to.Ptr[int32](3000),
					},
				},
				KubeEnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/kubeEnvironments/demokube"),
				Template: &armappservice.Template{
					Containers: []*armappservice.Container{
						{
							Name:  to.Ptr("testcontainerApp0"),
							Image: to.Ptr("repo/testcontainerApp0:v1"),
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
								Name: to.Ptr("httpscalingrule"),
								Custom: &armappservice.CustomScaleRule{
									Type: to.Ptr("http"),
									Metadata: map[string]*string{
										"concurrentRequests": to.Ptr("50"),
									},
								},
							}},
					},
				},
			},
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappservice%2Farmappservice%2Fv1.0.0/sdk/resourcemanager/appservice/armappservice/README.md) on how to add the SDK to your project and authenticate.
