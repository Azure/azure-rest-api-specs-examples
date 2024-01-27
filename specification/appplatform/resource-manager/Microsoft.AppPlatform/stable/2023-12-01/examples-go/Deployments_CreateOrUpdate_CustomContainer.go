package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/685aad3f33d355c1d9c89d493ee9398865367bd8/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Deployments_CreateOrUpdate_CustomContainer.json
func ExampleDeploymentsClient_BeginCreateOrUpdate_deploymentsCreateOrUpdateCustomContainer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", "myapp", "mydeployment", armappplatform.DeploymentResource{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeploymentResource = armappplatform.DeploymentResource{
	// 	Name: to.Ptr("mydeployment"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/apps/deployments"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apps/myapp/deployments/mydeployment"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.DeploymentResourceProperties{
	// 		Active: to.Ptr(false),
	// 		DeploymentSettings: &armappplatform.DeploymentSettings{
	// 			Apms: []*armappplatform.ApmReference{
	// 				{
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/apms/myappinsights"),
	// 			}},
	// 			EnvironmentVariables: map[string]*string{
	// 				"env": to.Ptr("test"),
	// 			},
	// 			LivenessProbe: &armappplatform.Probe{
	// 				DisableProbe: to.Ptr(false),
	// 				FailureThreshold: to.Ptr[int32](3),
	// 				InitialDelaySeconds: to.Ptr[int32](30),
	// 				PeriodSeconds: to.Ptr[int32](10),
	// 				ProbeAction: &armappplatform.HTTPGetAction{
	// 					Type: to.Ptr(armappplatform.ProbeActionTypeHTTPGetAction),
	// 					Path: to.Ptr("/health"),
	// 					Scheme: to.Ptr(armappplatform.HTTPSchemeTypeHTTP),
	// 				},
	// 			},
	// 			ReadinessProbe: &armappplatform.Probe{
	// 				DisableProbe: to.Ptr(false),
	// 				FailureThreshold: to.Ptr[int32](3),
	// 				InitialDelaySeconds: to.Ptr[int32](30),
	// 				PeriodSeconds: to.Ptr[int32](10),
	// 				ProbeAction: &armappplatform.HTTPGetAction{
	// 					Type: to.Ptr(armappplatform.ProbeActionTypeHTTPGetAction),
	// 					Path: to.Ptr("/health"),
	// 					Scheme: to.Ptr(armappplatform.HTTPSchemeTypeHTTP),
	// 				},
	// 			},
	// 			ResourceRequests: &armappplatform.ResourceRequests{
	// 				CPU: to.Ptr("1000m"),
	// 				Memory: to.Ptr("3Gi"),
	// 			},
	// 			TerminationGracePeriodSeconds: to.Ptr[int32](30),
	// 		},
	// 		Instances: []*armappplatform.DeploymentInstance{
	// 			{
	// 				Name: to.Ptr("instance1"),
	// 				DiscoveryStatus: to.Ptr("N/A"),
	// 				StartTime: to.Ptr("2020-08-26T01:55:02Z"),
	// 				Status: to.Ptr("Running"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armappplatform.DeploymentResourceProvisioningStateSucceeded),
	// 		Source: &armappplatform.CustomContainerUserSourceInfo{
	// 			Type: to.Ptr("Container"),
	// 			CustomContainer: &armappplatform.CustomContainer{
	// 				ContainerImage: to.Ptr("myContainerImage:v1"),
	// 				ImageRegistryCredential: &armappplatform.ImageRegistryCredential{
	// 					Password: to.Ptr(""),
	// 					Username: to.Ptr("myUsername"),
	// 				},
	// 				LanguageFramework: to.Ptr("springboot"),
	// 				Server: to.Ptr("myacr.azurecr.io"),
	// 			},
	// 		},
	// 		Status: to.Ptr(armappplatform.DeploymentResourceStatusRunning),
	// 	},
	// 	SKU: &armappplatform.SKU{
	// 		Name: to.Ptr("S0"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}
