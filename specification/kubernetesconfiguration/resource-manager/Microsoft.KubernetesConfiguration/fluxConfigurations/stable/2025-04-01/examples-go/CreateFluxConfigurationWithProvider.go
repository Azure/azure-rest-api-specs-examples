package armfluxconfigurations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armfluxconfigurations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba0c086df0ebe03a61579485c1c10de0d17804b2/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/fluxConfigurations/stable/2025-04-01/examples/CreateFluxConfigurationWithProvider.json
func ExampleClient_BeginCreateOrUpdate_createFluxConfigurationWithGitRepositoryProvider() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfluxconfigurations.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreateOrUpdate(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "srs-fluxconfig", armfluxconfigurations.FluxConfiguration{
		Properties: &armfluxconfigurations.FluxConfigurationProperties{
			GitRepository: &armfluxconfigurations.GitRepositoryDefinition{
				HTTPSCACert: to.Ptr("ZXhhbXBsZWNlcnRpZmljYXRl"),
				Provider:    to.Ptr(armfluxconfigurations.ProviderTypeAzure),
				RepositoryRef: &armfluxconfigurations.RepositoryRefDefinition{
					Branch: to.Ptr("master"),
				},
				SyncIntervalInSeconds: to.Ptr[int64](600),
				TimeoutInSeconds:      to.Ptr[int64](600),
				URL:                   to.Ptr("https://dev.azure.com/org/proj/_git/arc-k8s-demo"),
			},
			Kustomizations: map[string]*armfluxconfigurations.KustomizationDefinition{
				"srs-kustomization1": {
					Path:      to.Ptr("./test/path"),
					DependsOn: []*string{},
					PostBuild: &armfluxconfigurations.PostBuildDefinition{
						Substitute: map[string]*string{
							"cluster_env":   to.Ptr("prod"),
							"replica_count": to.Ptr("2"),
						},
						SubstituteFrom: []*armfluxconfigurations.SubstituteFromDefinition{
							{
								Name:     to.Ptr("cluster-test"),
								Kind:     to.Ptr("ConfigMap"),
								Optional: to.Ptr(true),
							}},
					},
					SyncIntervalInSeconds: to.Ptr[int64](600),
					TimeoutInSeconds:      to.Ptr[int64](600),
					Wait:                  to.Ptr(true),
				},
				"srs-kustomization2": {
					Path: to.Ptr("./other/test/path"),
					DependsOn: []*string{
						to.Ptr("srs-kustomization1")},
					PostBuild: &armfluxconfigurations.PostBuildDefinition{
						SubstituteFrom: []*armfluxconfigurations.SubstituteFromDefinition{
							{
								Name:     to.Ptr("cluster-values"),
								Kind:     to.Ptr("ConfigMap"),
								Optional: to.Ptr(true),
							},
							{
								Name:     to.Ptr("secret-name"),
								Kind:     to.Ptr("Secret"),
								Optional: to.Ptr(false),
							}},
					},
					Prune:                  to.Ptr(false),
					RetryIntervalInSeconds: to.Ptr[int64](600),
					SyncIntervalInSeconds:  to.Ptr[int64](600),
					TimeoutInSeconds:       to.Ptr[int64](600),
					Wait:                   to.Ptr(false),
				},
			},
			Namespace:                  to.Ptr("srs-namespace"),
			ReconciliationWaitDuration: to.Ptr("PT30M"),
			Scope:                      to.Ptr(armfluxconfigurations.ScopeTypeCluster),
			SourceKind:                 to.Ptr(armfluxconfigurations.SourceKindTypeGitRepository),
			Suspend:                    to.Ptr(false),
			WaitForReconciliation:      to.Ptr(true),
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
	// res.FluxConfiguration = armfluxconfigurations.FluxConfiguration{
	// 	Name: to.Ptr("srs-fluxconfig"),
	// 	Type: to.Ptr("Microsoft.KubernetesConfiguration/fluxConfigurations"),
	// 	ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/srs-fluxconfig"),
	// 	Properties: &armfluxconfigurations.FluxConfigurationProperties{
	// 		ComplianceState: to.Ptr(armfluxconfigurations.FluxComplianceStateCompliant),
	// 		ErrorMessage: to.Ptr(""),
	// 		GitRepository: &armfluxconfigurations.GitRepositoryDefinition{
	// 			HTTPSCACert: to.Ptr("ZXhhbXBsZWNlcnRpZmljYXRl"),
	// 			RepositoryRef: &armfluxconfigurations.RepositoryRefDefinition{
	// 				Branch: to.Ptr("master"),
	// 			},
	// 			SyncIntervalInSeconds: to.Ptr[int64](600),
	// 			TimeoutInSeconds: to.Ptr[int64](600),
	// 			URL: to.Ptr("https://github.com/Azure/arc-k8s-demo"),
	// 		},
	// 		Kustomizations: map[string]*armfluxconfigurations.KustomizationDefinition{
	// 			"srs-kustomization1": &armfluxconfigurations.KustomizationDefinition{
	// 				Name: to.Ptr("srs-kustomization1"),
	// 				Path: to.Ptr("./test/path"),
	// 				DependsOn: []*string{
	// 				},
	// 				PostBuild: &armfluxconfigurations.PostBuildDefinition{
	// 					Substitute: map[string]*string{
	// 						"cluster_env": to.Ptr("prod"),
	// 						"replica_count": to.Ptr("2"),
	// 					},
	// 					SubstituteFrom: []*armfluxconfigurations.SubstituteFromDefinition{
	// 						{
	// 							Name: to.Ptr("cluster-test"),
	// 							Kind: to.Ptr("ConfigMap"),
	// 							Optional: to.Ptr(true),
	// 					}},
	// 				},
	// 				SyncIntervalInSeconds: to.Ptr[int64](600),
	// 				TimeoutInSeconds: to.Ptr[int64](600),
	// 				Wait: to.Ptr(true),
	// 			},
	// 			"srs-kustomization2": &armfluxconfigurations.KustomizationDefinition{
	// 				Name: to.Ptr("srs-kustomization2"),
	// 				Path: to.Ptr("./other/test/path"),
	// 				DependsOn: []*string{
	// 					to.Ptr("srs-kustomization1")},
	// 					PostBuild: &armfluxconfigurations.PostBuildDefinition{
	// 						SubstituteFrom: []*armfluxconfigurations.SubstituteFromDefinition{
	// 							{
	// 								Name: to.Ptr("cluster-values"),
	// 								Kind: to.Ptr("ConfigMap"),
	// 								Optional: to.Ptr(true),
	// 							},
	// 							{
	// 								Name: to.Ptr("secret-name"),
	// 								Kind: to.Ptr("Secret"),
	// 								Optional: to.Ptr(false),
	// 						}},
	// 					},
	// 					Prune: to.Ptr(false),
	// 					RetryIntervalInSeconds: to.Ptr[int64](600),
	// 					SyncIntervalInSeconds: to.Ptr[int64](600),
	// 					TimeoutInSeconds: to.Ptr[int64](600),
	// 					Wait: to.Ptr(false),
	// 				},
	// 			},
	// 			Namespace: to.Ptr("srs-namespace"),
	// 			ProvisioningState: to.Ptr(armfluxconfigurations.ProvisioningStateSucceeded),
	// 			RepositoryPublicKey: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDiNkrANrhtRy+02Xc7b5bwvgOKvQMbqUQaXeB6FCDkbLaavw/zO/vIhIBEQu+vbBt6IlI/Pui0rMFr5JjA8Vwwd85oabzU07TPzbFvKSU9eCXqWRKWf0DHNQj/kxPJNtyPYFv3lGoiZZ6QzejOxlW/lPPokUePN0oI10daWwqznm2q0Cmh1EgPUYveq3J5KCWncZXCdwY36zWYulCWFaqazoaGy4kxcqlVy+mPjE/UJthaoLm3mq+23uhlmmfCc1j7W6+H6fcOwTyOtcbimxz2Ug8HlTzSTXBPtEe7qyllMyk62EPNUUq4bVoVsex9sKBK6/hW0Cn2P5i5jslUPCQF"),
	// 			Scope: to.Ptr(armfluxconfigurations.ScopeTypeCluster),
	// 			SourceKind: to.Ptr(armfluxconfigurations.SourceKindTypeGitRepository),
	// 			SourceSyncedCommitID: to.Ptr("master/0ba6f0d30760d567de0bac86c8c4eec13ce1a590"),
	// 			SourceUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12.000Z"); return t}()),
	// 			StatusUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12.000Z"); return t}()),
	// 			Statuses: []*armfluxconfigurations.ObjectStatusDefinition{
	// 				{
	// 					Name: to.Ptr("srs-fluxconfig"),
	// 					ComplianceState: to.Ptr(armfluxconfigurations.FluxComplianceStateCompliant),
	// 					Kind: to.Ptr("GitRepository"),
	// 					StatusConditions: []*armfluxconfigurations.ObjectStatusConditionDefinition{
	// 						{
	// 							Type: to.Ptr("Ready"),
	// 							LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:12:40.000Z"); return t}()),
	// 							Message: to.Ptr("'Fetched revision: master/0ba6f0d30760d567de0bac86c8c4eec13ce1a590'"),
	// 							Reason: to.Ptr("GitOperationSucceed"),
	// 							Status: to.Ptr("True"),
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("srs-fluxconfig-srs-kustomization1"),
	// 					ComplianceState: to.Ptr(armfluxconfigurations.FluxComplianceStateCompliant),
	// 					HelmReleaseProperties: &armfluxconfigurations.HelmReleasePropertiesDefinition{
	// 						HelmChartRef: &armfluxconfigurations.ObjectReferenceDefinition{
	// 							Name: to.Ptr("myname"),
	// 							Namespace: to.Ptr("mynamespace"),
	// 						},
	// 						LastRevisionApplied: to.Ptr[int64](1),
	// 					},
	// 					Kind: to.Ptr("Kustomization"),
	// 					StatusConditions: []*armfluxconfigurations.ObjectStatusConditionDefinition{
	// 						{
	// 							Type: to.Ptr("Ready"),
	// 							LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:12:40.000Z"); return t}()),
	// 							Message: to.Ptr("'Applied revision: master/0ba6f0d30760d567de0bac86c8c4eec13ce1a590'"),
	// 							Reason: to.Ptr("ReconciliationSucceeded"),
	// 							Status: to.Ptr("True"),
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("srs-fluxconfig-srs-kustomization2"),
	// 					ComplianceState: to.Ptr(armfluxconfigurations.FluxComplianceStateCompliant),
	// 					HelmReleaseProperties: &armfluxconfigurations.HelmReleasePropertiesDefinition{
	// 						HelmChartRef: &armfluxconfigurations.ObjectReferenceDefinition{
	// 							Name: to.Ptr("myname"),
	// 							Namespace: to.Ptr("mynamespace"),
	// 						},
	// 						LastRevisionApplied: to.Ptr[int64](1),
	// 					},
	// 					Kind: to.Ptr("Kustomization"),
	// 					StatusConditions: []*armfluxconfigurations.ObjectStatusConditionDefinition{
	// 						{
	// 							Type: to.Ptr("Ready"),
	// 							LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:12:40.000Z"); return t}()),
	// 							Message: to.Ptr("'Applied revision: master/0ba6f0d30760d567de0bac86c8c4eec13ce1a590'"),
	// 							Reason: to.Ptr("ReconciliationSucceeded"),
	// 							Status: to.Ptr("True"),
	// 					}},
	// 			}},
	// 			Suspend: to.Ptr(false),
	// 		},
	// 		SystemData: &armfluxconfigurations.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T05:10:57.027Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armfluxconfigurations.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-19T05:10:57.027Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armfluxconfigurations.CreatedByTypeApplication),
	// 		},
	// 	}
}
