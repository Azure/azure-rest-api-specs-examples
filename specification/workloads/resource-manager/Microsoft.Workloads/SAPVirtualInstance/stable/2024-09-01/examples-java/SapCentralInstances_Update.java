
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapCentralServerInstance;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapCentralServerInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapCentralInstances_Update.json
     */
    /**
     * Sample code: SapCentralServerInstances_Update.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sapCentralServerInstancesUpdate(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        SapCentralServerInstance resource = manager.sapCentralServerInstances()
            .getWithResponse("test-rg", "X00", "centralServer", com.azure.core.util.Context.NONE).getValue();
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
