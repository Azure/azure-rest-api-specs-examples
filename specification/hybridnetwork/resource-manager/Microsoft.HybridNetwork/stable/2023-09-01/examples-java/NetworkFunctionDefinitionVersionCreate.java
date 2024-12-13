
import com.azure.resourcemanager.hybridnetwork.models.ApplicationEnablement;
import com.azure.resourcemanager.hybridnetwork.models.AzureArcKubernetesArtifactProfile;
import com.azure.resourcemanager.hybridnetwork.models.AzureArcKubernetesDeployMappingRuleProfile;
import com.azure.resourcemanager.hybridnetwork.models.AzureArcKubernetesHelmApplication;
import com.azure.resourcemanager.hybridnetwork.models.AzureArcKubernetesNetworkFunctionTemplate;
import com.azure.resourcemanager.hybridnetwork.models.ContainerizedNetworkFunctionDefinitionVersion;
import com.azure.resourcemanager.hybridnetwork.models.DependsOnProfile;
import com.azure.resourcemanager.hybridnetwork.models.HelmArtifactProfile;
import com.azure.resourcemanager.hybridnetwork.models.HelmInstallOptions;
import com.azure.resourcemanager.hybridnetwork.models.HelmMappingRuleProfile;
import com.azure.resourcemanager.hybridnetwork.models.HelmMappingRuleProfileOptions;
import com.azure.resourcemanager.hybridnetwork.models.HelmUpgradeOptions;
import com.azure.resourcemanager.hybridnetwork.models.ReferencedResource;
import java.util.Arrays;

/**
 * Samples for NetworkFunctionDefinitionVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionDefinitionVersionCreate.json
     */
    /**
     * Sample code: Create or update a network function definition version resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdateANetworkFunctionDefinitionVersionResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionDefinitionVersions().define("1.0.0").withRegion("eastus")
            .withExistingNetworkFunctionDefinitionGroup("rg", "TestPublisher", "TestNetworkFunctionDefinitionGroupName")
            .withProperties(new ContainerizedNetworkFunctionDefinitionVersion().withDeployParameters(
                "{\"type\":\"object\",\"properties\":{\"releaseName\":{\"type\":\"string\"},\"namespace\":{\"type\":\"string\"}}}")
                .withNetworkFunctionTemplate(new AzureArcKubernetesNetworkFunctionTemplate()
                    .withNetworkFunctionApplications(Arrays.asList(new AzureArcKubernetesHelmApplication()
                        .withName("fedrbac")
                        .withDependsOnProfile(new DependsOnProfile().withInstallDependsOn(Arrays.asList())
                            .withUninstallDependsOn(Arrays.asList()).withUpdateDependsOn(Arrays.asList()))
                        .withArtifactProfile(new AzureArcKubernetesArtifactProfile()
                            .withArtifactStore(new ReferencedResource().withId(
                                "/subscriptions/subid/resourcegroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/testArtifactStore"))
                            .withHelmArtifactProfile(new HelmArtifactProfile().withHelmPackageName("fed-rbac")
                                .withHelmPackageVersionRange("~2.1.3")
                                .withRegistryValuesPaths(Arrays.asList("global.registry.docker.repoPath"))
                                .withImagePullSecretsValuesPaths(Arrays.asList("global.imagePullSecrets"))))
                        .withDeployParametersMappingRuleProfile(new AzureArcKubernetesDeployMappingRuleProfile()
                            .withApplicationEnablement(ApplicationEnablement.ENABLED).withHelmMappingRuleProfile(
                                new HelmMappingRuleProfile().withReleaseNamespace("{deployParameters.namesapce}")
                                    .withReleaseName("{deployParameters.releaseName}").withHelmPackageVersion("2.1.3")
                                    .withValues("")
                                    .withOptions(new HelmMappingRuleProfileOptions()
                                        .withInstallOptions(new HelmInstallOptions().withAtomic("true")
                                            .withWaitOption("true").withTimeout("30"))
                                        .withUpgradeOptions(new HelmUpgradeOptions().withAtomic("true")
                                            .withWaitOption("true").withTimeout("30")))))))))
            .create();
    }
}
