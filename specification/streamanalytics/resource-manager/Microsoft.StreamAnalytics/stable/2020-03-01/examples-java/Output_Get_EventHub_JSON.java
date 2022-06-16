import com.azure.core.util.Context;

/** Samples for Outputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Get_EventHub_JSON.json
     */
    /**
     * Sample code: Get an Event Hub output with JSON serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAnEventHubOutputWithJSONSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg6912", "sj3310", "output5195", Context.NONE);
    }
}
