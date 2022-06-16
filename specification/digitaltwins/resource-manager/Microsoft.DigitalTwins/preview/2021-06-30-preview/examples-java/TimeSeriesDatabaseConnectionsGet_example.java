import com.azure.core.util.Context;

/** Samples for TimeSeriesDatabaseConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/TimeSeriesDatabaseConnectionsGet_example.json
     */
    /**
     * Sample code: Get time series database connection for a DigitalTwins instance.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getTimeSeriesDatabaseConnectionForADigitalTwinsInstance(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .timeSeriesDatabaseConnections()
            .getWithResponse("resRg", "myDigitalTwinsService", "myConnection", Context.NONE);
    }
}
