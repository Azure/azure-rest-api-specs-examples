import com.azure.resourcemanager.streamanalytics.models.AzureSqlDatabaseOutputDataSource;

/** Samples for Outputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_AzureSQL.json
     */
    /**
     * Sample code: Create an Azure SQL database output.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAnAzureSQLDatabaseOutput(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .outputs()
            .define("output1755")
            .withExistingStreamingjob("sjrg2157", "sj6458")
            .withDatasource(new AzureSqlDatabaseOutputDataSource())
            .create();
    }
}
