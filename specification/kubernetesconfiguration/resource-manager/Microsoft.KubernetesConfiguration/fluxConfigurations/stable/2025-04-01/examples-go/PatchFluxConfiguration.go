package armfluxconfigurations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armfluxconfigurations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba0c086df0ebe03a61579485c1c10de0d17804b2/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/fluxConfigurations/stable/2025-04-01/examples/PatchFluxConfiguration.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfluxconfigurations.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginUpdate(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "srs-fluxconfig", armfluxconfigurations.FluxConfigurationPatch{
		Properties: &armfluxconfigurations.FluxConfigurationPatchProperties{
			GitRepository: &armfluxconfigurations.GitRepositoryPatchDefinition{
				URL: to.Ptr("https://github.com/jonathan-innis/flux2-kustomize-helm-example.git"),
			},
			Kustomizations: map[string]*armfluxconfigurations.KustomizationPatchDefinition{
				"srs-kustomization1": nil,
				"srs-kustomization2": {
					Path:                  to.Ptr("./test/alt-path"),
					SyncIntervalInSeconds: to.Ptr[int64](300),
				},
				"srs-kustomization3": {
					Path:                  to.Ptr("./test/another-path"),
					SyncIntervalInSeconds: to.Ptr[int64](300),
				},
			},
			Suspend: to.Ptr(true),
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
	// 	Properties: &armfluxconfigurations.FluxConfigurationProperties{
	// 		ComplianceState: to.Ptr(armfluxconfigurations.FluxComplianceStateCompliant),
	// 		ErrorMessage: to.Ptr(""),
	// 		GitRepository: &armfluxconfigurations.GitRepositoryDefinition{
	// 			RepositoryRef: &armfluxconfigurations.RepositoryRefDefinition{
	// 				Branch: to.Ptr("master"),
	// 			},
	// 			SyncIntervalInSeconds: to.Ptr[int64](600),
	// 			TimeoutInSeconds: to.Ptr[int64](600),
	// 			URL: to.Ptr("https://github.com/jonathan-innis/flux2-kustomize-helm-example.git"),
	// 		},
	// 		Kustomizations: map[string]*armfluxconfigurations.KustomizationDefinition{
	// 			"srs-kustomization2": &armfluxconfigurations.KustomizationDefinition{
	// 				Path: to.Ptr("./test/alt-path"),
	// 				Prune: to.Ptr(false),
	// 				RetryIntervalInSeconds: to.Ptr[int64](600),
	// 				SyncIntervalInSeconds: to.Ptr[int64](300),
	// 				TimeoutInSeconds: to.Ptr[int64](600),
	// 			},
	// 			"srs-kustomization3": &armfluxconfigurations.KustomizationDefinition{
	// 				Path: to.Ptr("./test/another-path"),
	// 				Prune: to.Ptr(false),
	// 				SyncIntervalInSeconds: to.Ptr[int64](300),
	// 				TimeoutInSeconds: to.Ptr[int64](600),
	// 			},
	// 		},
	// 		Namespace: to.Ptr("srs-namespace"),
	// 		ProvisioningState: to.Ptr(armfluxconfigurations.ProvisioningStateSucceeded),
	// 		RepositoryPublicKey: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDiNkrANrhtRy+02Xc7b5bwvgOKvQMbqUQaXeB6FCDkbLaavw/zO/vIhIBEQu+vbBt6IlI/Pui0rMFr5JjA8Vwwd85oabzU07TPzbFvKSU9eCXqWRKWf0DHNQj/kxPJNtyPYFv3lGoiZZ6QzejOxlW/lPPokUePN0oI10daWwqznm2q0Cmh1EgPUYveq3J5KCWncZXCdwY36zWYulCWFaqazoaGy4kxcqlVy+mPjE/UJthaoLm3mq+23uhlmmfCc1j7W6+H6fcOwTyOtcbimxz2Ug8HlTzSTXBPtEe7qyllMyk62EPNUUq4bVoVsex9sKBK6/hW0Cn2P5i5jslUPCQF"),
	// 		Scope: to.Ptr(armfluxconfigurations.ScopeTypeCluster),
	// 		SourceKind: to.Ptr(armfluxconfigurations.SourceKindTypeGitRepository),
	// 		SourceSyncedCommitID: to.Ptr("master/0ba6f0d30760d567de0bac86c8c4eec13ce1a590"),
	// 		SourceUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12.000Z"); return t}()),
	// 		StatusUpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:17:12.000Z"); return t}()),
	// 		Statuses: []*armfluxconfigurations.ObjectStatusDefinition{
	// 			{
	// 				Name: to.Ptr("srs-fluxconfig"),
	// 				ComplianceState: to.Ptr(armfluxconfigurations.FluxComplianceStateCompliant),
	// 				Kind: to.Ptr("GitRepository"),
	// 				StatusConditions: []*armfluxconfigurations.ObjectStatusConditionDefinition{
	// 					{
	// 						Type: to.Ptr("Ready"),
	// 						LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:12:40.000Z"); return t}()),
	// 						Message: to.Ptr("'Fetched revision: master/0ba6f0d30760d567de0bac86c8c4eec13ce1a590'"),
	// 						Reason: to.Ptr("GitOperationSucceed"),
	// 						Status: to.Ptr("True"),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("srs-fluxconfig-srs-kustomization1"),
	// 				ComplianceState: to.Ptr(armfluxconfigurations.FluxComplianceStateCompliant),
	// 				HelmReleaseProperties: &armfluxconfigurations.HelmReleasePropertiesDefinition{
	// 					HelmChartRef: &armfluxconfigurations.ObjectReferenceDefinition{
	// 						Name: to.Ptr("myname"),
	// 						Namespace: to.Ptr("mynamespace"),
	// 					},
	// 					LastRevisionApplied: to.Ptr[int64](1),
	// 				},
	// 				Kind: to.Ptr("Kustomization"),
	// 				StatusConditions: []*armfluxconfigurations.ObjectStatusConditionDefinition{
	// 					{
	// 						Type: to.Ptr("Ready"),
	// 						LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:12:40.000Z"); return t}()),
	// 						Message: to.Ptr("'Applied revision: master/0ba6f0d30760d567de0bac86c8c4eec13ce1a590'"),
	// 						Reason: to.Ptr("ReconciliationSucceeded"),
	// 						Status: to.Ptr("True"),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("srs-fluxconfig-srs-kustomization2"),
	// 				ComplianceState: to.Ptr(armfluxconfigurations.FluxComplianceStateCompliant),
	// 				HelmReleaseProperties: &armfluxconfigurations.HelmReleasePropertiesDefinition{
	// 					HelmChartRef: &armfluxconfigurations.ObjectReferenceDefinition{
	// 						Name: to.Ptr("myname"),
	// 						Namespace: to.Ptr("mynamespace"),
	// 					},
	// 					LastRevisionApplied: to.Ptr[int64](1),
	// 				},
	// 				Kind: to.Ptr("Kustomization"),
	// 				StatusConditions: []*armfluxconfigurations.ObjectStatusConditionDefinition{
	// 					{
	// 						Type: to.Ptr("Ready"),
	// 						LastTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-19T18:12:40.000Z"); return t}()),
	// 						Message: to.Ptr("'Applied revision: master/0ba6f0d30760d567de0bac86c8c4eec13ce1a590'"),
	// 						Reason: to.Ptr("ReconciliationSucceeded"),
	// 						Status: to.Ptr("True"),
	// 				}},
	// 		}},
	// 		Suspend: to.Ptr(false),
	// 	},
	// }
}
