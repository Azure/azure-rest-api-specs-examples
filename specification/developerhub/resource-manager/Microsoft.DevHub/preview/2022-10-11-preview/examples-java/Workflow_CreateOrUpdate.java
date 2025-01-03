
import com.azure.resourcemanager.devhub.models.Acr;
import com.azure.resourcemanager.devhub.models.DeploymentProperties;
import com.azure.resourcemanager.devhub.models.GitHubWorkflowProfileOidcCredentials;
import com.azure.resourcemanager.devhub.models.ManifestType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workflow CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/
     * Workflow_CreateOrUpdate.json
     */
    /**
     * Sample code: Create Workflow.
     * 
     * @param manager Entry point to DevHubManager.
     */
    public static void createWorkflow(com.azure.resourcemanager.devhub.DevHubManager manager) {
        manager.workflows().define("workflow1").withRegion("location1").withExistingResourceGroup("resourceGroup1")
            .withTags(mapOf("appname", "testApp")).withRepositoryOwner("owner1").withRepositoryName("repo1")
            .withBranchName("branch1").withDockerfile("repo1/images/Dockerfile").withDockerBuildContext("repo1/src/")
            .withDeploymentProperties(new DeploymentProperties().withManifestType(ManifestType.KUBE)
                .withKubeManifestLocations(Arrays.asList("/src/manifests/"))
                .withOverrides(mapOf("key1", "fakeTokenPlaceholder")))
            .withNamespace("namespace1")
            .withAcr(new Acr().withAcrSubscriptionId("subscriptionId1").withAcrResourceGroup("resourceGroup1")
                .withAcrRegistryName("registry1").withAcrRepositoryName("repo1"))
            .withOidcCredentials(
                new GitHubWorkflowProfileOidcCredentials().withAzureClientId("12345678-3456-7890-5678-012345678901")
                    .withAzureTenantId("66666666-3456-7890-5678-012345678901"))
            .withAksResourceId(
                "/subscriptions/subscriptionId1/resourcegroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1")
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
