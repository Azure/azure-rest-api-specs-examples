
/**
 * Samples for TimeSeriesDatabaseConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/
     * TimeSeriesDatabaseConnectionsGet_example.json
     */
    /**
     * Sample code: Get time series database connection for a DigitalTwins instance.
     * 
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getTimeSeriesDatabaseConnectionForADigitalTwinsInstance(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.timeSeriesDatabaseConnections().getWithResponse("resRg", "myDigitalTwinsService", "myConnection",
            com.azure.core.util.Context.NONE);
    }
}
