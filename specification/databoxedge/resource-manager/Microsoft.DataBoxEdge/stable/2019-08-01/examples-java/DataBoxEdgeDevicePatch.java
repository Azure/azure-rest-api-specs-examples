
import com.azure.resourcemanager.databoxedge.models.DataBoxEdgeDevice;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Devices Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/
     * DataBoxEdgeDevicePatch.json
     */
    /**
     * Sample code: DataBoxEdgeDevicePatch.
     * 
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void dataBoxEdgeDevicePatch(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        DataBoxEdgeDevice resource = manager.devices().getByResourceGroupWithResponse("GroupForEdgeAutomation",
            "testedgedevice", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("Key1", "fakeTokenPlaceholder", "Key2", "fakeTokenPlaceholder")).apply();
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
