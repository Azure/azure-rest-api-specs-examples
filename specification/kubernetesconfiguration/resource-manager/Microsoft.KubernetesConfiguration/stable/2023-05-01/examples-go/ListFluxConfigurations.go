package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/ListFluxConfigurations.json
func ExampleFluxConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkubernetesconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFluxConfigurationsClient().NewListPager("rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FluxConfigurationsList = armkubernetesconfiguration.FluxConfigurationsList{
		// 	Value: []*armkubernetesconfiguration.FluxConfiguration{
		// 		{
		// 			Name: to.Ptr("srs-fluxconfig"),
		// 			Type: to.Ptr("Microsoft.KubernetesConfiguration/fluxConfigurations"),
		// 			ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/srs-fluxconfig"),
		// 			Properties: &armkubernetesconfiguration.FluxConfigurationProperties{
		// 				ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
		// 				ErrorMessage: to.Ptr(""),
		// 				GitRepository: &armkubernetesconfiguration.GitRepositoryDefinition{
		// 					RepositoryRef: &armkubernetesconfiguration.RepositoryRefDefinition{
		// 						Branch: to.Ptr("master"),
		// 					},
		// 					SyncIntervalInSeconds: to.Ptr[int64](600),
		// 					TimeoutInSeconds: to.Ptr[int64](600),
		// 					URL: to.Ptr("https://github.com/Azure/arc-k8s-demo"),
		// 				},
		// 				Kustomizations: map[string]*armkubernetesconfiguration.KustomizationDefinition{
		// 					"srs-kustomization1": &armkubernetesconfiguration.KustomizationDefinition{
		// 						Name: to.Ptr("srs-kustomization1"),
		// 						Path: to.Ptr("./test/path"),
		// 						DependsOn: []*string{
		// 						},
		// 						PostBuild: &armkubernetesconfiguration.PostBuildDefinition{
		// 							Substitute: map[string]*string{
		// 								"cluster_env": to.Ptr("prod"),
		// 								"replica_count": to.Ptr("2"),
		// 							},
		// 							SubstituteFrom: []*armkubernetesconfiguration.SubstituteFromDefinition{
		// 								{
		// 									Name: to.Ptr("cluster-test"),
		// 									Kind: to.Ptr("ConfigMap"),
		// 									Optional: to.Ptr(true),
		// 							}},
		// 						},
		// 						SyncIntervalInSeconds: to.Ptr[int64](600),
		// 						TimeoutInSeconds: to.Ptr[int64](600),
		// 						Wait: to.Ptr(true),
		// 					},
		// 					"srs-kustomization2": &armkubernetesconfiguration.KustomizationDefinition{
		// 						Name: to.Ptr("srs-kustomization2"),
		// 						Path: to.Ptr("./other/test/path"),
		// 						DependsOn: []*string{
		// 							to.Ptr("srs-kustomization1")},
		// 							PostBuild: &armkubernetesconfiguration.PostBuildDefinition{
		// 								SubstituteFrom: []*armkubernetesconfiguration.SubstituteFromDefinition{
		// 									{
		// 										Name: to.Ptr("cluster-values"),
		// 										Kind: to.Ptr("ConfigMap"),
		// 										Optional: to.Ptr(true),
		// 									},
		// 									{
		// 										Name: to.Ptr("secret-name"),
		// 										Kind: to.Ptr("Secret"),
		// 										Optional: to.Ptr(false),
		// 								}},
		// 							},
		// 							Prune: to.Ptr(false),
		// 							RetryIntervalInSeconds: to.Ptr[int64](600),
		// 							SyncIntervalInSeconds: to.Ptr[int64](600),
		// 							TimeoutInSeconds: to.Ptr[int64](600),
		// 							Wait: to.Ptr(false),
		// 						},
		// 					},
		// 					Namespace: to.Ptr("srs-namespace"),
		// 					ProvisioningState: to.Ptr(armkubernetesconfiguration.ProvisioningStateSucceeded),
		// 					RepositoryPublicKey: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDiNkrANrhtRy+02Xc7b5bwvgOKvQMbqUQaXeB6FCDkbLaavw/zO/vIhIBEQu+vbBt6IlI/Pui0rMFr5JjA8Vwwd85oabzU07TPzbFvKSU9eCXqWRKWf0DHNQj/kxPJNtyPYFv3lGoiZZ6QzejOxlW/lPPokUePN0oI10daWwqznm2q0Cmh1EgPUYveq3J5KCWncZXCdwY36zWYulCWFaqazoaGy4kxcqlVy+mPjE/UJthaoLm3mq+23uhlmmfCc1j7W6+H6fcOwTyOtcbimxz2Ug8HlTzSTXBPtEe7qyllMyk62EPNUUq4bVoVsex9sKBK6/hW0Cn2P5i5jslUPCQF"),
		// 					Scope: to.Ptr(armkubernetesconfiguration.ScopeTypeCluster),
		// 					SourceKind: to.Ptr(armkubernetesconfiguration.SourceKindTypeGitRepository),
		// 					SourceSyncedCommitID: to.Ptr("master/0ba6f0d30760d567de0bac86c8c4eec13ce1a590"),
		// 					SourceUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12.000Z"); return t}()),
		// 					StatusUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12.000Z"); return t}()),
		// 					Statuses: []*armkubernetesconfiguration.ObjectStatusDefinition{
		// 						{
		// 							Name: to.Ptr("srs-fluxconfig"),
		// 							ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
		// 							Kind: to.Ptr("GitRepository"),
		// 						},
		// 						{
		// 							Name: to.Ptr("srs-fluxconfig-srs-kustomization1"),
		// 							ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
		// 							Kind: to.Ptr("Kustomization"),
		// 						},
		// 						{
		// 							Name: to.Ptr("srs-fluxconfig-srs-kustomization2"),
		// 							ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
		// 							Kind: to.Ptr("Kustomization"),
		// 					}},
		// 					Suspend: to.Ptr(false),
		// 				},
		// 				SystemData: &armkubernetesconfiguration.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T05:10:57.027Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-19T05:10:57.027Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("srs-fluxconfig"),
		// 				Type: to.Ptr("Microsoft.KubernetesConfiguration/fluxConfigurations"),
		// 				ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/srs-fluxconfig"),
		// 				Properties: &armkubernetesconfiguration.FluxConfigurationProperties{
		// 					ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
		// 					ErrorMessage: to.Ptr(""),
		// 					GitRepository: &armkubernetesconfiguration.GitRepositoryDefinition{
		// 						RepositoryRef: &armkubernetesconfiguration.RepositoryRefDefinition{
		// 							Branch: to.Ptr("master"),
		// 						},
		// 						SyncIntervalInSeconds: to.Ptr[int64](600),
		// 						TimeoutInSeconds: to.Ptr[int64](600),
		// 						URL: to.Ptr("https://github.com/Azure/arc-k8s-demo"),
		// 					},
		// 					Kustomizations: map[string]*armkubernetesconfiguration.KustomizationDefinition{
		// 						"srs-kustomization1": &armkubernetesconfiguration.KustomizationDefinition{
		// 							Name: to.Ptr("srs-kustomization1"),
		// 							Path: to.Ptr("./test/path"),
		// 							DependsOn: []*string{
		// 							},
		// 							SyncIntervalInSeconds: to.Ptr[int64](600),
		// 							TimeoutInSeconds: to.Ptr[int64](600),
		// 						},
		// 						"srs-kustomization2": &armkubernetesconfiguration.KustomizationDefinition{
		// 							Name: to.Ptr("srs-kustomization2"),
		// 							Path: to.Ptr("./other/test/path"),
		// 							DependsOn: []*string{
		// 								to.Ptr("srs-kustomization1")},
		// 								Prune: to.Ptr(false),
		// 								RetryIntervalInSeconds: to.Ptr[int64](600),
		// 								SyncIntervalInSeconds: to.Ptr[int64](600),
		// 								TimeoutInSeconds: to.Ptr[int64](600),
		// 							},
		// 						},
		// 						Namespace: to.Ptr("srs-namespace"),
		// 						ProvisioningState: to.Ptr(armkubernetesconfiguration.ProvisioningStateSucceeded),
		// 						RepositoryPublicKey: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDiNkrANrhtRy+02Xc7b5bwvgOKvQMbqUQaXeB6FCDkbLaavw/zO/vIhIBEQu+vbBt6IlI/Pui0rMFr5JjA8Vwwd85oabzU07TPzbFvKSU9eCXqWRKWf0DHNQj/kxPJNtyPYFv3lGoiZZ6QzejOxlW/lPPokUePN0oI10daWwqznm2q0Cmh1EgPUYveq3J5KCWncZXCdwY36zWYulCWFaqazoaGy4kxcqlVy+mPjE/UJthaoLm3mq+23uhlmmfCc1j7W6+H6fcOwTyOtcbimxz2Ug8HlTzSTXBPtEe7qyllMyk62EPNUUq4bVoVsex9sKBK6/hW0Cn2P5i5jslUPCQF"),
		// 						Scope: to.Ptr(armkubernetesconfiguration.ScopeTypeCluster),
		// 						SourceKind: to.Ptr(armkubernetesconfiguration.SourceKindTypeGitRepository),
		// 						SourceSyncedCommitID: to.Ptr("master/0ba6f0d30760d567de0bac86c8c4eec13ce1a590"),
		// 						SourceUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12.000Z"); return t}()),
		// 						StatusUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12.000Z"); return t}()),
		// 						Statuses: []*armkubernetesconfiguration.ObjectStatusDefinition{
		// 							{
		// 								Name: to.Ptr("srs-fluxconfig"),
		// 								ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
		// 								Kind: to.Ptr("GitRepository"),
		// 							},
		// 							{
		// 								Name: to.Ptr("srs-fluxconfig-srs-kustomization1"),
		// 								ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
		// 								Kind: to.Ptr("Kustomization"),
		// 							},
		// 							{
		// 								Name: to.Ptr("srs-fluxconfig-srs-kustomization2"),
		// 								ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
		// 								Kind: to.Ptr("Kustomization"),
		// 						}},
		// 						Suspend: to.Ptr(false),
		// 					},
		// 					SystemData: &armkubernetesconfiguration.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T05:10:57.027Z"); return t}()),
		// 						CreatedBy: to.Ptr("string"),
		// 						CreatedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-19T05:10:57.027Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("string"),
		// 						LastModifiedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 					},
		// 			}},
		// 		}
	}
}
