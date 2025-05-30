
import com.azure.resourcemanager.connectedcache.models.EnterpriseMccCustomerResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for EnterpriseMccCustomers Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/EnterpriseMccCustomers_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: EnterpriseMccCustomers_Update.
     * 
     * @param manager Entry point to ConnectedCacheManager.
     */
    public static void
        enterpriseMccCustomersUpdate(com.azure.resourcemanager.connectedcache.ConnectedCacheManager manager) {
        EnterpriseMccCustomerResource resource = manager.enterpriseMccCustomers()
            .getByResourceGroupWithResponse("rgConnectedCache", "MccRPTest1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1878", "fakeTokenPlaceholder")).apply();
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
