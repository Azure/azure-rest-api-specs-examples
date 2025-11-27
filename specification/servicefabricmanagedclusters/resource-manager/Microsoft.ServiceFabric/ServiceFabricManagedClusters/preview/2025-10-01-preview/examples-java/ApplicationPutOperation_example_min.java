
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Applications CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ApplicationPutOperation_example_min.json
     */
    /**
     * Sample code: Put an application with minimum parameters.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putAnApplicationWithMinimumParameters(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applications().define("myApp").withExistingManagedCluster("resRg", "myCluster").withRegion("eastus")
            .withVersion(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType/versions/1.0")
            .create();
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
