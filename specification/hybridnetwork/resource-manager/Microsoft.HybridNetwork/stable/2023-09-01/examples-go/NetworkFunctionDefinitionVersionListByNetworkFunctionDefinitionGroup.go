package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionVersionListByNetworkFunctionDefinitionGroup.json
func ExampleNetworkFunctionDefinitionVersionsClient_NewListByNetworkFunctionDefinitionGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkFunctionDefinitionVersionsClient().NewListByNetworkFunctionDefinitionGroupPager("rg", "TestPublisher", "TestNetworkFunctionDefinitionGroupNameName", nil)
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
		// page.NetworkFunctionDefinitionVersionListResult = armhybridnetwork.NetworkFunctionDefinitionVersionListResult{
		// 	Value: []*armhybridnetwork.NetworkFunctionDefinitionVersion{
		// 		{
		// 			Name: to.Ptr("networkFunctionDefinitionVersionName"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/publishers/publisherskugroups/publisherskuversions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/TestPublisher/publisherskuGroups/networkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/TestVersion"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armhybridnetwork.ContainerizedNetworkFunctionDefinitionVersion{
		// 				DeployParameters: to.Ptr("{\"releaseName\":{\"type\":\"string\"},\"namespace\":{\"type\":\"string\"}}"),
		// 				NetworkFunctionType: to.Ptr(armhybridnetwork.NetworkFunctionTypeContainerizedNetworkFunction),
		// 				VersionState: to.Ptr(armhybridnetwork.VersionStateActive),
		// 				NetworkFunctionTemplate: &armhybridnetwork.AzureArcKubernetesNetworkFunctionTemplate{
		// 					NfviType: to.Ptr(armhybridnetwork.ContainerizedNetworkFunctionNFVITypeAzureArcKubernetes),
		// 					NetworkFunctionApplications: []armhybridnetwork.AzureArcKubernetesNetworkFunctionApplicationClassification{
		// 						&armhybridnetwork.AzureArcKubernetesHelmApplication{
		// 							Name: to.Ptr("fedrbac"),
		// 							DependsOnProfile: &armhybridnetwork.DependsOnProfile{
		// 								InstallDependsOn: []*string{
		// 								},
		// 								UninstallDependsOn: []*string{
		// 								},
		// 								UpdateDependsOn: []*string{
		// 								},
		// 							},
		// 							ArtifactType: to.Ptr(armhybridnetwork.AzureArcKubernetesArtifactTypeHelmPackage),
		// 							ArtifactProfile: &armhybridnetwork.AzureArcKubernetesArtifactProfile{
		// 								ArtifactStore: &armhybridnetwork.ReferencedResource{
		// 									ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/testArtifactStore"),
		// 								},
		// 								HelmArtifactProfile: &armhybridnetwork.HelmArtifactProfile{
		// 									HelmPackageName: to.Ptr("fed-rbac"),
		// 									HelmPackageVersionRange: to.Ptr("~2.1.3"),
		// 									ImagePullSecretsValuesPaths: []*string{
		// 										to.Ptr("global.imagePullSecrets")},
		// 										RegistryValuesPaths: []*string{
		// 											to.Ptr("global.registry.docker.repoPath")},
		// 										},
		// 									},
		// 									DeployParametersMappingRuleProfile: &armhybridnetwork.AzureArcKubernetesDeployMappingRuleProfile{
		// 										ApplicationEnablement: to.Ptr(armhybridnetwork.ApplicationEnablementEnabled),
		// 										HelmMappingRuleProfile: &armhybridnetwork.HelmMappingRuleProfile{
		// 											HelmPackageVersion: to.Ptr("2.1.3"),
		// 											Options: &armhybridnetwork.HelmMappingRuleProfileOptions{
		// 												InstallOptions: &armhybridnetwork.HelmInstallOptions{
		// 													Atomic: to.Ptr("true"),
		// 													Timeout: to.Ptr("30"),
		// 													Wait: to.Ptr("waitValue"),
		// 												},
		// 												UpgradeOptions: &armhybridnetwork.HelmUpgradeOptions{
		// 													Atomic: to.Ptr("true"),
		// 													Timeout: to.Ptr("30"),
		// 													Wait: to.Ptr("waitValue"),
		// 												},
		// 											},
		// 											ReleaseName: to.Ptr("{deployParameters.releaseName}"),
		// 											ReleaseNamespace: to.Ptr("{deployParameters.namesapce}"),
		// 											Values: to.Ptr(""),
		// 										},
		// 									},
		// 							}},
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
