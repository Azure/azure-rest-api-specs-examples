
import com.azure.resourcemanager.deviceupdate.models.Instance;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Instances Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/Instances/
     * Instances_Update.json
     */
    /**
     * Sample code: Updates Instance.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void updatesInstance(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        Instance resource = manager.instances()
            .getWithResponse("test-rg", "contoso", "blue", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tagKey", "fakeTokenPlaceholder")).apply();
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
