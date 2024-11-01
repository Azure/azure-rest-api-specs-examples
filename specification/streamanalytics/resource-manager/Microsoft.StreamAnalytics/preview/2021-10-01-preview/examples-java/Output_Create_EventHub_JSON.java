
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.EventHubOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.JsonOutputSerializationFormat;
import com.azure.resourcemanager.streamanalytics.models.JsonSerialization;
import com.azure.resourcemanager.streamanalytics.models.OutputWatermarkMode;
import com.azure.resourcemanager.streamanalytics.models.OutputWatermarkProperties;

/**
 * Samples for Outputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Create_EventHub_JSON.json
     */
    /**
     * Sample code: Create an Event Hub output with JSON serialization.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAnEventHubOutputWithJSONSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().define("output5195").withExistingStreamingjob("sjrg6912", "sj3310")
            .withDatasource(
                new EventHubOutputDataSource().withPartitionKey("fakeTokenPlaceholder").withEventHubName("sdkeventhub")
                    .withServiceBusNamespace("sdktest").withSharedAccessPolicyName("RootManageSharedAccessKey")
                    .withSharedAccessPolicyKey("fakeTokenPlaceholder"))
            .withSerialization(
                new JsonSerialization().withEncoding(Encoding.UTF8).withFormat(JsonOutputSerializationFormat.ARRAY))
            .withWatermarkSettings(
                new OutputWatermarkProperties().withWatermarkMode(OutputWatermarkMode.SEND_CURRENT_PARTITION_WATERMARK)
                    .withMaxWatermarkDifferenceAcrossPartitions("16:14:30"))
            .create();
    }
}
