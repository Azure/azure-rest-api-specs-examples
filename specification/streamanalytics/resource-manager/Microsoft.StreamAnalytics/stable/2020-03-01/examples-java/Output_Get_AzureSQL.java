import com.azure.core.util.Context;

/** Samples for Outputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Get_AzureSQL.json
     */
    /**
     * Sample code: Get an Azure SQL database output.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAnAzureSQLDatabaseOutput(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg2157", "sj6458", "output1755", Context.NONE);
    }
}
