
import com.azure.resourcemanager.streamanalytics.models.AzureSqlDatabaseOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.Output;

/**
 * Samples for Outputs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Update_AzureSQL.json
     */
    /**
     * Sample code: Update an Azure SQL database output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        updateAnAzureSQLDatabaseOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource = manager.outputs()
            .getWithResponse("sjrg2157", "sj6458", "output1755", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDatasource(new AzureSqlDatabaseOutputDataSource().withTable("differentTable")).apply();
    }
}
