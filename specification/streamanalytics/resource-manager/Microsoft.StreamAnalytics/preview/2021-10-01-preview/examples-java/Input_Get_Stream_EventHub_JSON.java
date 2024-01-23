
/**
 * Samples for Inputs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Get_Stream_EventHub_JSON.json
     */
    /**
     * Sample code: Get a stream Event Hub input with JSON serialization.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAStreamEventHubInputWithJSONSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().getWithResponse("sjrg3139", "sj197", "input7425", com.azure.core.util.Context.NONE);
    }
}
