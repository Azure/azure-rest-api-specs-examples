
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapDatabaseInstances Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/
     * SAPDatabaseInstances_Create.json
     */
    /**
     * Sample code: SAPDatabaseInstances_Create.
     * 
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPDatabaseInstancesCreate(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapDatabaseInstances().define("databaseServer").withRegion("westcentralus")
            .withExistingSapVirtualInstance("test-rg", "X00").withTags(mapOf()).create();
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
