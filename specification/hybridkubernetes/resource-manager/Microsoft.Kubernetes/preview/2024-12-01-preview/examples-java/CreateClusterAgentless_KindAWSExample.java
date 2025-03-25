
import com.azure.resourcemanager.hybridkubernetes.models.ConnectedClusterIdentity;
import com.azure.resourcemanager.hybridkubernetes.models.ConnectedClusterKind;
import com.azure.resourcemanager.hybridkubernetes.models.ResourceIdentityType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConnectedCluster CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/
     * CreateClusterAgentless_KindAWSExample.json
     */
    /**
     * Sample code: CreateClusterAgentless_KindAWSExample.
     * 
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void createClusterAgentlessKindAWSExample(
        com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        manager.connectedClusters().define("testCluster").withRegion("East US").withExistingResourceGroup("k8sc-rg")
            .withIdentity(new ConnectedClusterIdentity().withType(ResourceIdentityType.NONE))
            .withAgentPublicKeyCertificate("").withTags(mapOf()).withKind(ConnectedClusterKind.AWS)
            .withDistribution("eks").withInfrastructure("aws").create();
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
