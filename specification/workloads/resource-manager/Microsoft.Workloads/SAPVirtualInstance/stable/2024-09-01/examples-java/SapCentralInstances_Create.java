
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapCentralServerProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapCentralServerInstances Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapCentralInstances_Create.json
     */
    /**
     * Sample code: SapCentralServerInstances_Create.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sapCentralServerInstancesCreate(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapCentralServerInstances().define("centralServer").withRegion("westcentralus")
            .withExistingSapVirtualInstance("test-rg", "X00").withTags(mapOf())
            .withProperties(new SapCentralServerProperties()).create();
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
