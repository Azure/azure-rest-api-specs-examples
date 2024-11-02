
/**
 * Samples for Outputs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Get_ServiceBusTopic_CSV.json
     */
    /**
     * Sample code: Get a Service Bus Topic output with CSV serialization.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAServiceBusTopicOutputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg6450", "sj7094", "output7886", com.azure.core.util.Context.NONE);
    }
}
