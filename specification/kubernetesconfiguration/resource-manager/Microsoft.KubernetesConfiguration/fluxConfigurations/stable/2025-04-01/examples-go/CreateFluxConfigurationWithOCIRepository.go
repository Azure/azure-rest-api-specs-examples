package armfluxconfigurations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armfluxconfigurations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba0c086df0ebe03a61579485c1c10de0d17804b2/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/fluxConfigurations/stable/2025-04-01/examples/CreateFluxConfigurationWithOCIRepository.json
func ExampleClient_BeginCreateOrUpdate_createFluxConfigurationWithOciRepositorySourceKind() {
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
			Kustomizations: map[string]*armfluxconfigurations.KustomizationDefinition{
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
			Namespace: to.Ptr("srs-namespace"),
			OciRepository: &armfluxconfigurations.OCIRepositoryDefinition{
				ServiceAccountName:    to.Ptr("testserviceaccount"),
				SyncIntervalInSeconds: to.Ptr[int64](1000),
				TimeoutInSeconds:      to.Ptr[int64](1000),
				URL:                   to.Ptr("oci://ghcr.io/stefanprodan/manifests/podinfo"),
			},
			Scope:      to.Ptr(armfluxconfigurations.ScopeTypeCluster),
			SourceKind: to.Ptr(armfluxconfigurations.SourceKindTypeOCIRepository),
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
	// res.FluxConfiguration = armfluxconfigurations.FluxConfiguration{
	// 	Name: to.Ptr("srs-fluxconfig"),
	// 	Type: to.Ptr("Microsoft.KubernetesConfiguration/fluxConfigurations"),
	// 	ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/srs-fluxconfig"),
	// 	Properties: &armfluxconfigurations.FluxConfigurationProperties{
	// 		ComplianceState: to.Ptr(armfluxconfigurations.FluxComplianceStateCompliant),
	// 		ErrorMessage: to.Ptr(""),
	// 		Kustomizations: map[string]*armfluxconfigurations.KustomizationDefinition{
	// 			"srs-kustomization1": &armfluxconfigurations.KustomizationDefinition{
	// 				Name: to.Ptr("srs-kustomization1"),
	// 				Path: to.Ptr("./test/path"),
	// 				DependsOn: []*string{
	// 				},
	// 				SyncIntervalInSeconds: to.Ptr[int64](600),
	// 				TimeoutInSeconds: to.Ptr[int64](600),
	// 			},
	// 			"srs-kustomization2": &armfluxconfigurations.KustomizationDefinition{
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
	// 			OciRepository: &armfluxconfigurations.OCIRepositoryDefinition{
	// 				ServiceAccountName: to.Ptr("testserviceaccount"),
	// 				SyncIntervalInSeconds: to.Ptr[int64](1000),
	// 				TimeoutInSeconds: to.Ptr[int64](1000),
	// 				URL: to.Ptr("oci://ghcr.io/stefanprodan/manifests/podinfo"),
	// 			},
	// 			ProvisioningState: to.Ptr(armfluxconfigurations.ProvisioningStateSucceeded),
	// 			RepositoryPublicKey: to.Ptr(""),
	// 			Scope: to.Ptr(armfluxconfigurations.ScopeTypeCluster),
	// 			SourceKind: to.Ptr(armfluxconfigurations.SourceKindTypeOCIRepository),
	// 			SourceSyncedCommitID: to.Ptr("0ba6f0d30760d567de0bac86c8c4eec13ce1a590"),
	// 			SourceUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12.000Z"); return t}()),
	// 			StatusUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12.000Z"); return t}()),
	// 			Statuses: []*armfluxconfigurations.ObjectStatusDefinition{
	// 				{
	// 					Name: to.Ptr("srs-fluxconfig"),
	// 					ComplianceState: to.Ptr(armfluxconfigurations.FluxComplianceStateCompliant),
	// 					Kind: to.Ptr("OCIRepository"),
	// 					StatusConditions: []*armfluxconfigurations.ObjectStatusConditionDefinition{
	// 						{
	// 							Type: to.Ptr("Ready"),
	// 							LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-04T07:17:30.000Z"); return t}()),
	// 							Message: to.Ptr("stored artifact for revision '55396be14f18fa2b977c1f22becef26a94d1d9a7ccb1e19d12f9cac52d757a84'"),
	// 							Reason: to.Ptr("Succeeded"),
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
	// 							Message: to.Ptr("'Applied revision: 0ba6f0d30760d567de0bac86c8c4eec13ce1a590'"),
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
	// 							Message: to.Ptr("'Applied revision: 0ba6f0d30760d567de0bac86c8c4eec13ce1a590'"),
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
