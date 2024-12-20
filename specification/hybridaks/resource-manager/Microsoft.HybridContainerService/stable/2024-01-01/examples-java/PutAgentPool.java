
import com.azure.resourcemanager.hybridcontainerservice.models.AgentPoolProperties;
import com.azure.resourcemanager.hybridcontainerservice.models.OsType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPool CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/PutAgentPool
     * .json
     */
    /**
     * Sample code: PutAgentPool.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void
        putAgentPool(com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.agentPools().define("testnodepool").withExistingConnectedClusterResourceUri(
            "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster")
            .withProperties(
                new AgentPoolProperties().withOsType(OsType.LINUX).withNodeLabels(mapOf("env", "dev", "goal", "test"))
                    .withNodeTaints(Arrays.asList("env=prod:NoSchedule", "sku=gpu:NoSchedule")).withCount(1)
                    .withVmSize("Standard_A4_v2"))
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
