package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/Deployments_CreateOrUpdate_CustomContainer.json
func ExampleDeploymentsClient_BeginCreateOrUpdate_deploymentsCreateOrUpdateCustomContainer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewDeploymentsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", "myapp", "mydeployment", armappplatform.DeploymentResource{
		Properties: &armappplatform.DeploymentResourceProperties{
			DeploymentSettings: &armappplatform.DeploymentSettings{
				EnvironmentVariables: map[string]*string{
					"env": to.Ptr("test"),
				},
				LivenessProbe: &armappplatform.Probe{
					DisableProbe:        to.Ptr(false),
					FailureThreshold:    to.Ptr[int32](3),
					InitialDelaySeconds: to.Ptr[int32](30),
					PeriodSeconds:       to.Ptr[int32](10),
					ProbeAction: &armappplatform.HTTPGetAction{
						Type:   to.Ptr(armappplatform.ProbeActionTypeHTTPGetAction),
						Path:   to.Ptr("/health"),
						Scheme: to.Ptr(armappplatform.HTTPSchemeTypeHTTP),
					},
				},
				ReadinessProbe: &armappplatform.Probe{
					DisableProbe:        to.Ptr(false),
					FailureThreshold:    to.Ptr[int32](3),
					InitialDelaySeconds: to.Ptr[int32](30),
					PeriodSeconds:       to.Ptr[int32](10),
					ProbeAction: &armappplatform.HTTPGetAction{
						Type:   to.Ptr(armappplatform.ProbeActionTypeHTTPGetAction),
						Path:   to.Ptr("/health"),
						Scheme: to.Ptr(armappplatform.HTTPSchemeTypeHTTP),
					},
				},
				ResourceRequests: &armappplatform.ResourceRequests{
					CPU:    to.Ptr("1000m"),
					Memory: to.Ptr("3Gi"),
				},
				TerminationGracePeriodSeconds: to.Ptr[int32](30),
			},
			Source: &armappplatform.CustomContainerUserSourceInfo{
				Type: to.Ptr("Container"),
				CustomContainer: &armappplatform.CustomContainer{
					Args: []*string{
						to.Ptr("-c"),
						to.Ptr("while true; do echo hello; sleep 10;done")},
					Command: []*string{
						to.Ptr("/bin/sh")},
					ContainerImage: to.Ptr("myContainerImage:v1"),
					ImageRegistryCredential: &armappplatform.ImageRegistryCredential{
						Password: to.Ptr("myPassword"),
						Username: to.Ptr("myUsername"),
					},
					LanguageFramework: to.Ptr("springboot"),
					Server:            to.Ptr("myacr.azurecr.io"),
				},
			},
		},
	}, nil)
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
