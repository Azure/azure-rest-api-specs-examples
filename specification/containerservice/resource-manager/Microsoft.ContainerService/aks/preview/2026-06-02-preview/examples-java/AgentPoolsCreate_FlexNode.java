
import com.azure.resourcemanager.containerservice.fluent.models.AgentPoolInner;
import com.azure.resourcemanager.containerservice.models.AgentPoolMode;
import com.azure.resourcemanager.containerservice.models.AgentPoolType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AgentPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/AgentPoolsCreate_FlexNode.json
     */
    /**
     * Sample code: Create FlexNode Agent Pool.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        createFlexNodeAgentPool(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().createOrUpdate(
            "rg1", "clustername1", "flexnode1", new AgentPoolInner().withTypePropertiesType(AgentPoolType.FLEX_NODES)
                .withMode(AgentPoolMode.USER).withOrchestratorVersion("1.32"),
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
