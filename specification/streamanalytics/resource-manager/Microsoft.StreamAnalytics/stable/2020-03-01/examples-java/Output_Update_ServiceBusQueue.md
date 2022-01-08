Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.JsonOutputSerializationFormat;
import com.azure.resourcemanager.streamanalytics.models.JsonSerialization;
import com.azure.resourcemanager.streamanalytics.models.Output;
import com.azure.resourcemanager.streamanalytics.models.ServiceBusQueueOutputDataSource;

/** Samples for Outputs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_ServiceBusQueue.json
     */
    /**
     * Sample code: Update a Service Bus Queue output with Avro serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAServiceBusQueueOutputWithAvroSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource =
            manager.outputs().getWithResponse("sjrg3410", "sj5095", "output3456", Context.NONE).getValue();
        resource
            .update()
            .withDatasource(new ServiceBusQueueOutputDataSource().withQueueName("differentQueueName"))
            .withSerialization(
                new JsonSerialization()
                    .withEncoding(Encoding.UTF8)
                    .withFormat(JsonOutputSerializationFormat.LINE_SEPARATED))
            .apply();
    }
}
```
