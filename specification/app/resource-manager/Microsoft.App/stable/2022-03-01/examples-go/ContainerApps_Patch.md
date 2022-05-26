Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappcontainers%2Farmappcontainers%2Fv0.1.0/sdk/resourcemanager/appcontainers/armappcontainers/README.md) on how to add the SDK to your project and authenticate.

```go
package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ContainerApps_Patch.json
func ExampleContainerAppsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappcontainers.NewContainerAppsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"rg",
		"testcontainerApp0",
		armappcontainers.ContainerApp{
			Location: to.Ptr("East US"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Properties: &armappcontainers.ContainerAppProperties{
				Configuration: &armappcontainers.Configuration{
					Dapr: &armappcontainers.Dapr{
						AppPort:     to.Ptr[int32](3000),
						AppProtocol: to.Ptr(armappcontainers.AppProtocolHTTP),
						Enabled:     to.Ptr(true),
					},
					Ingress: &armappcontainers.Ingress{
						CustomDomains: []*armappcontainers.CustomDomain{
							{
								Name:          to.Ptr("www.my-name.com"),
								BindingType:   to.Ptr(armappcontainers.BindingTypeSniEnabled),
								CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
							},
							{
								Name:          to.Ptr("www.my-other-name.com"),
								BindingType:   to.Ptr(armappcontainers.BindingTypeSniEnabled),
								CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com"),
							}},
						External:   to.Ptr(true),
						TargetPort: to.Ptr[int32](3000),
						Traffic: []*armappcontainers.TrafficWeight{
							{
								Label:        to.Ptr("production"),
								RevisionName: to.Ptr("testcontainerApp0-ab1234"),
								Weight:       to.Ptr[int32](100),
							}},
					},
				},
				Template: &armappcontainers.Template{
					Containers: []*armappcontainers.Container{
						{
							Name:  to.Ptr("testcontainerApp0"),
							Image: to.Ptr("repo/testcontainerApp0:v1"),
							Probes: []*armappcontainers.ContainerAppProbe{
								{
									Type: to.Ptr(armappcontainers.TypeLiveness),
									HTTPGet: &armappcontainers.ContainerAppProbeHTTPGet{
										Path: to.Ptr("/health"),
										HTTPHeaders: []*armappcontainers.ContainerAppProbeHTTPGetHTTPHeadersItem{
											{
												Name:  to.Ptr("Custom-Header"),
												Value: to.Ptr("Awesome"),
											}},
										Port: to.Ptr[int32](8080),
									},
									InitialDelaySeconds: to.Ptr[int32](3),
									PeriodSeconds:       to.Ptr[int32](3),
								}},
						}},
					Scale: &armappcontainers.Scale{
						MaxReplicas: to.Ptr[int32](5),
						MinReplicas: to.Ptr[int32](1),
						Rules: []*armappcontainers.ScaleRule{
							{
								Name: to.Ptr("httpscalingrule"),
								Custom: &armappcontainers.CustomScaleRule{
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
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```
