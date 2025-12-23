
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for QueryPacks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/QueryPacksCreate.json
     */
    /**
     * Sample code: QueryPackCreate.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryPackCreate(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.queryPacks().define("my-querypack").withRegion("South Central US")
            .withExistingResourceGroup("my-resource-group").create();
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
