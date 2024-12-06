
/**
 * Samples for TimeSeriesDatabaseConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/
     * TimeSeriesDatabaseConnectionsDelete_example.json
     */
    /**
     * Sample code: Delete a time series database connection for a DigitalTwins instance.
     * 
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void deleteATimeSeriesDatabaseConnectionForADigitalTwinsInstance(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.timeSeriesDatabaseConnections().delete("resRg", "myDigitalTwinsService", "myConnection", null,
            com.azure.core.util.Context.NONE);
    }
}
