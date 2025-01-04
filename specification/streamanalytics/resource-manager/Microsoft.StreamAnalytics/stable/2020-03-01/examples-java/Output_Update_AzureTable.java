
import com.azure.resourcemanager.streamanalytics.models.AzureTableOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.Output;

/**
 * Samples for Outputs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/
     * Output_Update_AzureTable.json
     */
    /**
     * Sample code: Update an Azure Table output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        updateAnAzureTableOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource = manager.outputs()
            .getWithResponse("sjrg5176", "sj2790", "output958", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDatasource(new AzureTableOutputDataSource().withPartitionKey("fakeTokenPlaceholder"))
            .apply();
    }
}
