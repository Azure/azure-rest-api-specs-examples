import com.azure.core.util.Context;

/** Samples for Inputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Get_Stream_IoTHub_Avro.json
     */
    /**
     * Sample code: Get a stream IoT Hub input with Avro serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAStreamIoTHubInputWithAvroSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().getWithResponse("sjrg3467", "sj9742", "input7970", Context.NONE);
    }
}
