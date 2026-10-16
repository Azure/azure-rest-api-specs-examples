
import com.azure.resourcemanager.edgeactions.models.EdgeAction;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for EdgeActions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeActions_Update.json
     */
    /**
     * Sample code: UpdateEdgeAction.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void updateEdgeAction(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        EdgeAction resource = manager.edgeActions()
            .getByResourceGroupWithResponse("testrg", "edgeAction1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("environment", "production")).apply();
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
