import com.azure.core.util.Context;

/** Samples for Outputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Get_AzureTable.json
     */
    /**
     * Sample code: Get an Azure Table output.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAnAzureTableOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg5176", "sj2790", "output958", Context.NONE);
    }
}
