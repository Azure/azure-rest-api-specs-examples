
import com.azure.resourcemanager.hybridkubernetes.models.AzureHybridBenefit;
import com.azure.resourcemanager.hybridkubernetes.models.ConnectedCluster;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConnectedCluster Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2024-12-01-preview/examples/
     * UpdateClusterByPatchExample.json
     */
    /**
     * Sample code: UpdateClusterExample.
     * 
     * @param manager Entry point to HybridKubernetesManager.
     */
    public static void
        updateClusterExample(com.azure.resourcemanager.hybridkubernetes.HybridKubernetesManager manager) {
        ConnectedCluster resource = manager.connectedClusters()
            .getByResourceGroupWithResponse("k8sc-rg", "testCluster", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).withDistribution("AKS")
            .withDistributionVersion("1.0").withAzureHybridBenefit(AzureHybridBenefit.NOT_APPLICABLE).apply();
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
