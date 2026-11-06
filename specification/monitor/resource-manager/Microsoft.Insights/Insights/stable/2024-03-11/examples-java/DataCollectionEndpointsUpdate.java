
import com.azure.resourcemanager.monitor.fluent.models.DataCollectionEndpointResourceInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DataCollectionEndpoints Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionEndpointsUpdate.json
     */
    /**
     * Sample code: Update a data collection endpoint.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void updateADataCollectionEndpoint(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionEndpoints().createWithResponse(
            "myResourceGroup", "myDataCollectionEndpoint", new DataCollectionEndpointResourceInner()
                .withLocation("eastus").withTags(mapOf("tag1", "A", "tag2", "B", "tag3", "C")),
            com.azure.core.util.Context.NONE);
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
