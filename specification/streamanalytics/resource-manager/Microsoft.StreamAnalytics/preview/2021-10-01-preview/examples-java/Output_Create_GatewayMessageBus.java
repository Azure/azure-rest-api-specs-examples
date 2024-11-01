
import com.azure.resourcemanager.streamanalytics.models.GatewayMessageBusOutputDataSource;

/**
 * Samples for Outputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Create_GatewayMessageBus.json
     */
    /**
     * Sample code: Create a Gateway Message Bus output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAGatewayMessageBusOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().define("output3022").withExistingStreamingjob("sjrg7983", "sj2331")
            .withDatasource(new GatewayMessageBusOutputDataSource().withTopic("EdgeTopic1")).create();
    }
}
