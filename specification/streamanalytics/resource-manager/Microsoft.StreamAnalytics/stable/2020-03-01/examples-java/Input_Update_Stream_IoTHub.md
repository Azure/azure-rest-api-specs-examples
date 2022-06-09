```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.Input;
import com.azure.resourcemanager.streamanalytics.models.IoTHubStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;

/** Samples for Inputs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Update_Stream_IoTHub.json
     */
    /**
     * Sample code: Update a stream IoT Hub input.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAStreamIoTHubInput(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Input resource = manager.inputs().getWithResponse("sjrg3467", "sj9742", "input7970", Context.NONE).getValue();
        resource
            .update()
            .withProperties(
                new StreamInputProperties()
                    .withSerialization(new CsvSerialization().withFieldDelimiter("|").withEncoding(Encoding.UTF8))
                    .withDatasource(
                        new IoTHubStreamInputDataSource().withEndpoint("messages/operationsMonitoringEvents")))
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
