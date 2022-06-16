import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.AvroSerialization;
import com.azure.resourcemanager.streamanalytics.models.EventHubStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.Input;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;

/** Samples for Inputs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Update_Stream_EventHub.json
     */
    /**
     * Sample code: Update a stream Event Hub input.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAStreamEventHubInput(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Input resource = manager.inputs().getWithResponse("sjrg3139", "sj197", "input7425", Context.NONE).getValue();
        resource
            .update()
            .withProperties(
                new StreamInputProperties()
                    .withSerialization(new AvroSerialization())
                    .withDatasource(
                        new EventHubStreamInputDataSource().withConsumerGroupName("differentConsumerGroupName")))
            .apply();
    }
}
