
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.OSType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-02-01/examples/
     * AgentPoolsCreate_PPG.json
     */
    /**
     * Sample code: Create Agent Pool with PPG.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAgentPoolWithPPG(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1",
            "agentpool1",
            new AgentPoolInner().withCount(3).withVmSize("Standard_DS2_v2").withOsType(OSType.LINUX)
                .withOrchestratorVersion("").withProximityPlacementGroupId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Compute/proximityPlacementGroups/ppg1"),
            null, null, com.azure.core.util.Context.NONE);
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
