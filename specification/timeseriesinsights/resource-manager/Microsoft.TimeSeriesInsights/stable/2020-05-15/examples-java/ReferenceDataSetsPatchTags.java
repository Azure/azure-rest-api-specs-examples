import com.azure.resourcemanager.timeseriesinsights.models.ReferenceDataSetResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for ReferenceDataSets Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsPatchTags.json
     */
    /**
     * Sample code: ReferenceDataSetsUpdate.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void referenceDataSetsUpdate(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        ReferenceDataSetResource resource =
            manager
                .referenceDataSets()
                .getWithResponse("rg1", "env1", "rds1", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("someKey", "someValue")).apply();
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
