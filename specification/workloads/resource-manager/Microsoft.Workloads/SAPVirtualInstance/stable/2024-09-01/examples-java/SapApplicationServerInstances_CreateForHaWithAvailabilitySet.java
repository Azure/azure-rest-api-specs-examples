
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapApplicationServerProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapApplicationServerInstances Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapApplicationServerInstances_CreateForHaWithAvailabilitySet.json
     */
    /**
     * Sample code: Create SAP Application Server Instances for HA System with Availability Set.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void createSAPApplicationServerInstancesForHASystemWithAvailabilitySet(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapApplicationServerInstances().define("app01").withRegion("westcentralus")
            .withExistingSapVirtualInstance("test-rg", "X00").withTags(mapOf())
            .withProperties(new SapApplicationServerProperties()).create();
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
