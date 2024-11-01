
import com.azure.resourcemanager.streamanalytics.models.AuthenticationMode;
import com.azure.resourcemanager.streamanalytics.models.AzureDataExplorerOutputDataSource;

/**
 * Samples for Outputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Create_AzureDataExplorer.json
     */
    /**
     * Sample code: Create an Azure Data Explorer output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAnAzureDataExplorerOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().define("adxOutput").withExistingStreamingjob("sjrg", "sjName")
            .withDatasource(
                new AzureDataExplorerOutputDataSource().withCluster("https://asakustotest.eastus.kusto.windows.net")
                    .withDatabase("dbname").withTable("mytable").withAuthenticationMode(AuthenticationMode.MSI))
            .create();
    }
}
