import com.azure.resourcemanager.timeseriesinsights.models.EnvironmentUpdateParameters;
import java.util.HashMap;
import java.util.Map;

/** Samples for Environments Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EnvironmentsPatchTags.json
     */
    /**
     * Sample code: EnvironmentsUpdate.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void environmentsUpdate(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager
            .environments()
            .update(
                "rg1",
                "env1",
                new EnvironmentUpdateParameters().withTags(mapOf("someTag", "someTagValue")),
                com.azure.core.util.Context.NONE);
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
