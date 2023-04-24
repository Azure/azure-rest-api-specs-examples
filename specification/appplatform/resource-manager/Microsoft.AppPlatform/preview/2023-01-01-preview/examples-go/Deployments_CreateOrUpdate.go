package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-01-01-preview/examples/Deployments_CreateOrUpdate.json
func ExampleDeploymentsClient_BeginCreateOrUpdate_deploymentsCreateOrUpdate() {
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
				AddonConfigs: map[string]any{
					"ApplicationConfigurationService": map[string]any{
						"patterns": []any{
							"mypattern",
						},
					},
				},
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
			Source: &armappplatform.SourceUploadedUserSourceInfo{
				Type:             to.Ptr("Source"),
				Version:          to.Ptr("1.0"),
				RelativePath:     to.Ptr("resources/a172cedcae47474b615c54d510a5d84a8dea3032e958587430b413538be3f333-2019082605-e3095339-1723-44b7-8b5e-31b1003978bc"),
				ArtifactSelector: to.Ptr("sub-module-1"),
			},
		},
		SKU: &armappplatform.SKU{
			Name:     to.Ptr("S0"),
			Capacity: to.Ptr[int32](1),
			Tier:     to.Ptr("Standard"),
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
	// 				DiscoveryStatus: to.Ptr("pending"),
	// 				StartTime: to.Ptr("2020-08-26T01:55:02Z"),
	// 				Status: to.Ptr("Running"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armappplatform.DeploymentResourceProvisioningStateSucceeded),
	// 		Source: &armappplatform.SourceUploadedUserSourceInfo{
	// 			Type: to.Ptr("Source"),
	// 			Version: to.Ptr("1.0"),
	// 			RelativePath: to.Ptr("resources/a172cedcae47474b615c54d510a5d84a8dea3032e958587430b413538be3f333-2019082605-e3095339-1723-44b7-8b5e-31b1003978bc"),
	// 			ArtifactSelector: to.Ptr("sub-module-1"),
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
