
import com.azure.resourcemanager.loganalytics.models.LogAnalyticsQueryPackQuerySearchProperties;
import com.azure.resourcemanager.loganalytics.models.LogAnalyticsQueryPackQuerySearchPropertiesRelated;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Queries Search.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/
     * QueryPackQueriesSearch.json
     */
    /**
     * Sample code: QuerySearch.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void querySearch(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.queries().search("my-resource-group", "my-querypack", new LogAnalyticsQueryPackQuerySearchProperties()
            .withRelated(new LogAnalyticsQueryPackQuerySearchPropertiesRelated()
                .withCategories(Arrays.asList("other", "analytics")))
            .withTags(mapOf("my-label", Arrays.asList("label1"))), 3L, true, null, com.azure.core.util.Context.NONE);
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
