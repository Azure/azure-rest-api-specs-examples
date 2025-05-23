
/**
 * Samples for Outputs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Get_ServiceBusQueue_Avro.json
     */
    /**
     * Sample code: Get a Service Bus Queue output with Avro serialization.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAServiceBusQueueOutputWithAvroSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg3410", "sj5095", "output3456", com.azure.core.util.Context.NONE);
    }
}
