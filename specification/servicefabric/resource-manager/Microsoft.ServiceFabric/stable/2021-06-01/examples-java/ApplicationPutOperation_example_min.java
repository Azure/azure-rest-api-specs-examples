
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Applications CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/
     * ApplicationPutOperation_example_min.json
     */
    /**
     * Sample code: Put an application with minimum parameters.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void
        putAnApplicationWithMinimumParameters(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applications().define("myApp").withExistingCluster("resRg", "myCluster").withRegion("eastus")
            .withTags(mapOf()).withTypeName("myAppType").withTypeVersion("1.0").withRemoveApplicationCapacity(false)
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
