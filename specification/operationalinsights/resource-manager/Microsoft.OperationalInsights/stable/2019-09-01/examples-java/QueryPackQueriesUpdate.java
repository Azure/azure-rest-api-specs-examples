
import com.azure.resourcemanager.loganalytics.models.LogAnalyticsQueryPackQuery;
import com.azure.resourcemanager.loganalytics.models.LogAnalyticsQueryPackQueryPropertiesRelated;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Queries Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/
     * QueryPackQueriesUpdate.json
     */
    /**
     * Sample code: QueryPatch.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryPatch(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        LogAnalyticsQueryPackQuery resource = manager.queries().getWithResponse("my-resource-group", "my-querypack",
            "a449f8af-8e64-4b3a-9b16-5a7165ff98c4", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withTags(mapOf("my-label", Arrays.asList("label1"), "my-other-label", Arrays.asList("label2")))
            .withDisplayName("Exceptions - New in the last 24 hours").withDescription("my description")
            .withBody(
                "let newExceptionsTimeRange = 1d;\nlet timeRangeToCheckBefore = 7d;\nexceptions\n| where timestamp < ago(timeRangeToCheckBefore)\n| summarize count() by problemId\n| join kind= rightanti (\nexceptions\n| where timestamp >= ago(newExceptionsTimeRange)\n| extend stack = tostring(details[0].rawStack)\n| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  \n) on problemId \n| order by  count_ desc\n")
            .withRelated(new LogAnalyticsQueryPackQueryPropertiesRelated().withCategories(Arrays.asList("analytics")))
            .apply();
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
