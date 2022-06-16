import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.EventHubOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.JsonOutputSerializationFormat;
import com.azure.resourcemanager.streamanalytics.models.JsonSerialization;
import com.azure.resourcemanager.streamanalytics.models.Output;

/** Samples for Outputs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_EventHub.json
     */
    /**
     * Sample code: Update an Event Hub output with JSON serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAnEventHubOutputWithJSONSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource =
            manager.outputs().getWithResponse("sjrg6912", "sj3310", "output5195", Context.NONE).getValue();
        resource
            .update()
            .withDatasource(new EventHubOutputDataSource().withPartitionKey("differentPartitionKey"))
            .withSerialization(
                new JsonSerialization()
                    .withEncoding(Encoding.UTF8)
                    .withFormat(JsonOutputSerializationFormat.LINE_SEPARATED))
            .apply();
    }
}
