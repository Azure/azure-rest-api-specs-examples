
import com.azure.resourcemanager.streamanalytics.models.GatewayMessageBusStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;

/**
 * Samples for Inputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Create_GatewayMessageBus.json
     */
    /**
     * Sample code: Create a Gateway Message Bus input.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAGatewayMessageBusInput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().define("input7970").withExistingStreamingjob("sjrg3467", "sj9742")
            .withProperties(new StreamInputProperties()
                .withDatasource(new GatewayMessageBusStreamInputDataSource().withTopic("EdgeTopic1")))
            .create();
    }
}
