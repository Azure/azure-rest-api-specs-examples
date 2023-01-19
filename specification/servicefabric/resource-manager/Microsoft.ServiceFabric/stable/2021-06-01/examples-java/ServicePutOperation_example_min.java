import com.azure.resourcemanager.servicefabric.models.SingletonPartitionSchemeDescription;
import com.azure.resourcemanager.servicefabric.models.StatelessServiceProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for Services CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePutOperation_example_min.json
     */
    /**
     * Sample code: Put a service with minimum parameters.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void putAServiceWithMinimumParameters(
        com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager
            .services()
            .define("myService")
            .withExistingApplication("resRg", "myCluster", "myApp")
            .withTags(mapOf())
            .withProperties(
                new StatelessServiceProperties()
                    .withServiceTypeName("myServiceType")
                    .withPartitionDescription(new SingletonPartitionSchemeDescription())
                    .withInstanceCount(1))
            .create();
    }

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
