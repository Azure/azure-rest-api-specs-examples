
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.Input;
import com.azure.resourcemanager.streamanalytics.models.IoTHubStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;

/**
 * Samples for Inputs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Update_Stream_IoTHub.json
     */
    /**
     * Sample code: Update a stream IoT Hub input.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        updateAStreamIoTHubInput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Input resource = manager.inputs()
            .getWithResponse("sjrg3467", "sj9742", "input7970", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new StreamInputProperties()
                .withSerialization(new CsvSerialization().withFieldDelimiter("|").withEncoding(Encoding.UTF8))
                .withDatasource(new IoTHubStreamInputDataSource().withEndpoint("messages/operationsMonitoringEvents")))
            .apply();
    }
}
