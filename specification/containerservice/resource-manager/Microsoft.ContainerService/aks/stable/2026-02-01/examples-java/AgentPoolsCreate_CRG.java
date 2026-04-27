
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.OSType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/AgentPoolsCreate_CRG.json
     */
    /**
     * Sample code: Create Agent Pool with Capacity Reservation Group.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void createAgentPoolWithCapacityReservationGroup(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().createOrUpdate("rg1", "clustername1", "agentpool1", new AgentPoolInner()
            .withCount(3).withVmSize("Standard_DS2_v2").withOsType(OSType.LINUX).withOrchestratorVersion("")
            .withCapacityReservationGroupId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/CapacityReservationGroups/crg1"),
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
