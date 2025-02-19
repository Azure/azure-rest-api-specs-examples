
import com.azure.resourcemanager.delegatednetwork.models.Orchestrator;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for OrchestratorInstanceService Patch.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/patchOrchestrator.json
     */
    /**
     * Sample code: update Orchestrator Instance.
     * 
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void
        updateOrchestratorInstance(com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        Orchestrator resource = manager.orchestratorInstanceServices()
            .getByResourceGroupWithResponse("TestRG", "testk8s1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key", "fakeTokenPlaceholder")).apply();
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
