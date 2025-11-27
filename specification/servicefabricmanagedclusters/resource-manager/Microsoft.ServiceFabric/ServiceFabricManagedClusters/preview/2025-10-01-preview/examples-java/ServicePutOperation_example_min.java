
import com.azure.resourcemanager.servicefabricmanagedclusters.models.SingletonPartitionScheme;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.StatelessServiceProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Services CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ServicePutOperation_example_min.json
     */
    /**
     * Sample code: Put a service with minimum parameters.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putAServiceWithMinimumParameters(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.services().define("myService").withExistingApplication("resRg", "myCluster", "myApp")
            .withRegion("eastus").withProperties(new StatelessServiceProperties().withServiceTypeName("myServiceType")
                .withPartitionDescription(new SingletonPartitionScheme()).withInstanceCount(1))
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
