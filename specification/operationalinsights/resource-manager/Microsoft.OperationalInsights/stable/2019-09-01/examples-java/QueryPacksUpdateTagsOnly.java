
import com.azure.resourcemanager.loganalytics.models.LogAnalyticsQueryPack;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for QueryPacks UpdateTags.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/
     * QueryPacksUpdateTagsOnly.json
     */
    /**
     * Sample code: QueryPackUpdateTagsOnly.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryPackUpdateTagsOnly(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        LogAnalyticsQueryPack resource = manager.queryPacks()
            .getByResourceGroupWithResponse("my-resource-group", "my-querypack", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("Tag1", "Value1", "Tag2", "Value2")).apply();
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
