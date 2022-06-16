import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.EventHubOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.JsonOutputSerializationFormat;
import com.azure.resourcemanager.streamanalytics.models.JsonSerialization;

/** Samples for Outputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_EventHub_JSON.json
     */
    /**
     * Sample code: Create an Event Hub output with JSON serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAnEventHubOutputWithJSONSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .outputs()
            .define("output5195")
            .withExistingStreamingjob("sjrg6912", "sj3310")
            .withDatasource(
                new EventHubOutputDataSource()
                    .withPartitionKey("partitionKey")
                    .withEventHubName("sdkeventhub")
                    .withServiceBusNamespace("sdktest")
                    .withSharedAccessPolicyName("RootManageSharedAccessKey")
                    .withSharedAccessPolicyKey("sharedAccessPolicyKey="))
            .withSerialization(
                new JsonSerialization().withEncoding(Encoding.UTF8).withFormat(JsonOutputSerializationFormat.ARRAY))
            .create();
    }
}
