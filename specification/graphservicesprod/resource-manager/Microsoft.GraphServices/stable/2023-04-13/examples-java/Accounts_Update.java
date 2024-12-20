
import com.azure.resourcemanager.graphservices.models.AccountResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Accounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/
     * Accounts_Update.json
     */
    /**
     * Sample code: Update account resource.
     * 
     * @param manager Entry point to GraphServicesManager.
     */
    public static void updateAccountResource(com.azure.resourcemanager.graphservices.GraphServicesManager manager) {
        AccountResource resource = manager.accounts().getByResourceGroupWithResponse("testResourceGroupGRAM",
            "11111111-aaaa-1111-bbbb-111111111111", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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
