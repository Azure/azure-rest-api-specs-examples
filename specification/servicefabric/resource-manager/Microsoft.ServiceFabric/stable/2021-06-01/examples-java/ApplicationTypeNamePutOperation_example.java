
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ApplicationTypes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/
     * ApplicationTypeNamePutOperation_example.json
     */
    /**
     * Sample code: Put an application type.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void putAnApplicationType(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applicationTypes().define("myAppType").withExistingCluster("resRg", "myCluster").withTags(mapOf())
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
