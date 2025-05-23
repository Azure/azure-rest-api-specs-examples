
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapApplicationServerInstance;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapApplicationServerInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapApplicationServerInstances_Update.json
     */
    /**
     * Sample code: SAPApplicationServerInstances_Update.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPApplicationServerInstancesUpdate(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        SapApplicationServerInstance resource = manager.sapApplicationServerInstances()
            .getWithResponse("test-rg", "X00", "app01", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1")).apply();
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
