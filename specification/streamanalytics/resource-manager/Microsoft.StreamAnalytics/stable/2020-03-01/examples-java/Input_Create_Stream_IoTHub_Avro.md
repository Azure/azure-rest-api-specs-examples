Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.streamanalytics.models.AvroSerialization;
import com.azure.resourcemanager.streamanalytics.models.IoTHubStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;

/** Samples for Inputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Create_Stream_IoTHub_Avro.json
     */
    /**
     * Sample code: Create a stream IoT Hub input with Avro serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAStreamIoTHubInputWithAvroSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .inputs()
            .define("input7970")
            .withExistingStreamingjob("sjrg3467", "sj9742")
            .withProperties(
                new StreamInputProperties()
                    .withSerialization(new AvroSerialization())
                    .withDatasource(
                        new IoTHubStreamInputDataSource()
                            .withIotHubNamespace("iothub")
                            .withSharedAccessPolicyName("owner")
                            .withSharedAccessPolicyKey("sharedAccessPolicyKey=")
                            .withConsumerGroupName("sdkconsumergroup")
                            .withEndpoint("messages/events")))
            .create();
    }
}
```
