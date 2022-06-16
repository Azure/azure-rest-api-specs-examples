import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.EventHubStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.JsonSerialization;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;

/** Samples for Inputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Create_Stream_EventHub_JSON.json
     */
    /**
     * Sample code: Create a stream Event Hub input with JSON serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAStreamEventHubInputWithJSONSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .inputs()
            .define("input7425")
            .withExistingStreamingjob("sjrg3139", "sj197")
            .withProperties(
                new StreamInputProperties()
                    .withSerialization(new JsonSerialization().withEncoding(Encoding.UTF8))
                    .withDatasource(
                        new EventHubStreamInputDataSource()
                            .withConsumerGroupName("sdkconsumergroup")
                            .withEventHubName("sdkeventhub")
                            .withServiceBusNamespace("sdktest")
                            .withSharedAccessPolicyName("RootManageSharedAccessKey")
                            .withSharedAccessPolicyKey("someSharedAccessPolicyKey==")))
            .create();
    }
}
