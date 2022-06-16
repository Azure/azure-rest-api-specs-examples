import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.Output;
import com.azure.resourcemanager.streamanalytics.models.ServiceBusTopicOutputDataSource;

/** Samples for Outputs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_ServiceBusTopic.json
     */
    /**
     * Sample code: Update a Service Bus Topic output with CSV serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAServiceBusTopicOutputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource =
            manager.outputs().getWithResponse("sjrg6450", "sj7094", "output7886", Context.NONE).getValue();
        resource
            .update()
            .withDatasource(new ServiceBusTopicOutputDataSource().withTopicName("differentTopicName"))
            .withSerialization(new CsvSerialization().withFieldDelimiter("|").withEncoding(Encoding.UTF8))
            .apply();
    }
}
