import com.azure.core.util.Context;

/** Samples for Inputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Get_Stream_EventHub_JSON.json
     */
    /**
     * Sample code: Get a stream Event Hub input with JSON serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAStreamEventHubInputWithJSONSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().getWithResponse("sjrg3139", "sj197", "input7425", Context.NONE);
    }
}
