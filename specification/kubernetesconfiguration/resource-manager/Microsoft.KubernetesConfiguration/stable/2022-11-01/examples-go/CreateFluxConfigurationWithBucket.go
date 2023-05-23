package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/CreateFluxConfigurationWithBucket.json
func ExampleFluxConfigurationsClient_BeginCreateOrUpdate_createFluxConfigurationWithBucketSourceKind() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkubernetesconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFluxConfigurationsClient().BeginCreateOrUpdate(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "srs-fluxconfig", armkubernetesconfiguration.FluxConfiguration{
		Properties: &armkubernetesconfiguration.FluxConfigurationProperties{
			Bucket: &armkubernetesconfiguration.BucketDefinition{
				AccessKey:             to.Ptr("fluxminiotest"),
				BucketName:            to.Ptr("flux"),
				SyncIntervalInSeconds: to.Ptr[int64](1000),
				TimeoutInSeconds:      to.Ptr[int64](1000),
				URL:                   to.Ptr("https://fluxminiotest.az.minio.io"),
			},
			Kustomizations: map[string]*armkubernetesconfiguration.KustomizationDefinition{
				"srs-kustomization1": {
					Path:                  to.Ptr("./test/path"),
					DependsOn:             []*string{},
					SyncIntervalInSeconds: to.Ptr[int64](600),
					TimeoutInSeconds:      to.Ptr[int64](600),
				},
				"srs-kustomization2": {
					Path: to.Ptr("./other/test/path"),
					DependsOn: []*string{
						to.Ptr("srs-kustomization1")},
					Prune:                  to.Ptr(false),
					RetryIntervalInSeconds: to.Ptr[int64](600),
					SyncIntervalInSeconds:  to.Ptr[int64](600),
					TimeoutInSeconds:       to.Ptr[int64](600),
				},
			},
			Namespace:  to.Ptr("srs-namespace"),
			Scope:      to.Ptr(armkubernetesconfiguration.ScopeTypeCluster),
			SourceKind: to.Ptr(armkubernetesconfiguration.SourceKindTypeBucket),
			Suspend:    to.Ptr(false),
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
	// res.FluxConfiguration = armkubernetesconfiguration.FluxConfiguration{
	// 	Name: to.Ptr("srs-fluxconfig"),
	// 	Type: to.Ptr("Microsoft.KubernetesConfiguration/fluxConfigurations"),
	// 	ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/srs-fluxconfig"),
	// 	Properties: &armkubernetesconfiguration.FluxConfigurationProperties{
	// 		Bucket: &armkubernetesconfiguration.BucketDefinition{
	// 			AccessKey: to.Ptr("fluxminiotest"),
	// 			BucketName: to.Ptr("flux"),
	// 			SyncIntervalInSeconds: to.Ptr[int64](1000),
	// 			TimeoutInSeconds: to.Ptr[int64](1000),
	// 			URL: to.Ptr("https://fluxminiotest.az.minio.io"),
	// 		},
	// 		ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
	// 		ErrorMessage: to.Ptr(""),
	// 		Kustomizations: map[string]*armkubernetesconfiguration.KustomizationDefinition{
	// 			"srs-kustomization1": &armkubernetesconfiguration.KustomizationDefinition{
	// 				Name: to.Ptr("srs-kustomization1"),
	// 				Path: to.Ptr("./test/path"),
	// 				DependsOn: []*string{
	// 				},
	// 				SyncIntervalInSeconds: to.Ptr[int64](600),
	// 				TimeoutInSeconds: to.Ptr[int64](600),
	// 			},
	// 			"srs-kustomization2": &armkubernetesconfiguration.KustomizationDefinition{
	// 				Name: to.Ptr("srs-kustomization2"),
	// 				Path: to.Ptr("./other/test/path"),
	// 				DependsOn: []*string{
	// 					to.Ptr("srs-kustomization1")},
	// 					Prune: to.Ptr(false),
	// 					RetryIntervalInSeconds: to.Ptr[int64](600),
	// 					SyncIntervalInSeconds: to.Ptr[int64](600),
	// 					TimeoutInSeconds: to.Ptr[int64](600),
	// 				},
	// 			},
	// 			Namespace: to.Ptr("srs-namespace"),
	// 			ProvisioningState: to.Ptr(armkubernetesconfiguration.ProvisioningStateSucceeded),
	// 			RepositoryPublicKey: to.Ptr(""),
	// 			Scope: to.Ptr(armkubernetesconfiguration.ScopeTypeCluster),
	// 			SourceKind: to.Ptr(armkubernetesconfiguration.SourceKindTypeBucket),
	// 			SourceSyncedCommitID: to.Ptr("0ba6f0d30760d567de0bac86c8c4eec13ce1a590"),
	// 			SourceUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12Z"); return t}()),
	// 			StatusUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12Z"); return t}()),
	// 			Statuses: []*armkubernetesconfiguration.ObjectStatusDefinition{
	// 				{
	// 					Name: to.Ptr("srs-fluxconfig"),
	// 					ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
	// 					Kind: to.Ptr("Bucket"),
	// 					StatusConditions: []*armkubernetesconfiguration.ObjectStatusConditionDefinition{
	// 						{
	// 							Type: to.Ptr("Ready"),
	// 							LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-04T07:17:30+00:00"); return t}()),
	// 							Message: to.Ptr("stored artifact for revision '55396be14f18fa2b977c1f22becef26a94d1d9a7ccb1e19d12f9cac52d757a84'"),
	// 							Reason: to.Ptr("Succeeded"),
	// 							Status: to.Ptr("True"),
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("srs-fluxconfig-srs-kustomization1"),
	// 					ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
	// 					HelmReleaseProperties: &armkubernetesconfiguration.HelmReleasePropertiesDefinition{
	// 						HelmChartRef: &armkubernetesconfiguration.ObjectReferenceDefinition{
	// 							Name: to.Ptr("myname"),
	// 							Namespace: to.Ptr("mynamespace"),
	// 						},
	// 						LastRevisionApplied: to.Ptr[int64](1),
	// 					},
	// 					Kind: to.Ptr("Kustomization"),
	// 					StatusConditions: []*armkubernetesconfiguration.ObjectStatusConditionDefinition{
	// 						{
	// 							Type: to.Ptr("Ready"),
	// 							LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:12:40Z"); return t}()),
	// 							Message: to.Ptr("'Applied revision: 0ba6f0d30760d567de0bac86c8c4eec13ce1a590'"),
	// 							Reason: to.Ptr("ReconciliationSucceeded"),
	// 							Status: to.Ptr("True"),
	// 					}},
	// 				},
	// 				{
	// 					Name: to.Ptr("srs-fluxconfig-srs-kustomization2"),
	// 					ComplianceState: to.Ptr(armkubernetesconfiguration.FluxComplianceStateCompliant),
	// 					HelmReleaseProperties: &armkubernetesconfiguration.HelmReleasePropertiesDefinition{
	// 						HelmChartRef: &armkubernetesconfiguration.ObjectReferenceDefinition{
	// 							Name: to.Ptr("myname"),
	// 							Namespace: to.Ptr("mynamespace"),
	// 						},
	// 						LastRevisionApplied: to.Ptr[int64](1),
	// 					},
	// 					Kind: to.Ptr("Kustomization"),
	// 					StatusConditions: []*armkubernetesconfiguration.ObjectStatusConditionDefinition{
	// 						{
	// 							Type: to.Ptr("Ready"),
	// 							LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:12:40Z"); return t}()),
	// 							Message: to.Ptr("'Applied revision: 0ba6f0d30760d567de0bac86c8c4eec13ce1a590'"),
	// 							Reason: to.Ptr("ReconciliationSucceeded"),
	// 							Status: to.Ptr("True"),
	// 					}},
	// 			}},
	// 			Suspend: to.Ptr(false),
	// 		},
	// 		SystemData: &armkubernetesconfiguration.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T05:10:57.027Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-19T05:10:57.027Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
	// 		},
	// 	}
}
