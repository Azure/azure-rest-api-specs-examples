
import com.azure.resourcemanager.streamanalytics.models.AzureSynapseOutputDataSource;

/**
 * Samples for Outputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/
     * Output_Create_DataWarehouse.json
     */
    /**
     * Sample code: Create an Azure Data Warehouse output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAnAzureDataWarehouseOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().define("dwOutput").withExistingStreamingjob("sjrg", "sjName")
            .withDatasource(new AzureSynapseOutputDataSource().withServer("asatestserver").withDatabase("zhayaSQLpool")
                .withTable("test2").withUser("tolladmin").withPassword("fakeTokenPlaceholder"))
            .create();
    }
}
