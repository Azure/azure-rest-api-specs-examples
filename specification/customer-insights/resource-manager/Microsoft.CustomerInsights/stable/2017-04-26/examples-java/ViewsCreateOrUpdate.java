
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Views CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/
     * ViewsCreateOrUpdate.json
     */
    /**
     * Sample code: Views_CreateOrUpdate.
     * 
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void viewsCreateOrUpdate(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.views().define("testView").withExistingHub("TestHubRG", "sdkTestHub").withUserId("testUser")
            .withDisplayName(mapOf("en", "some name"))
            .withDefinition("{\\\"isProfileType\\\":false,\\\"profileTypes\\\":[],\\\"widgets\\\":[],\\\"style\\\":[]}")
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
