
import com.azure.resourcemanager.streamanalytics.models.AuthenticationMode;
import com.azure.resourcemanager.streamanalytics.models.AzureSynapseOutputDataSource;

/**
 * Samples for Outputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
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
                .withTable("test2").withUser("tolladmin").withPassword("fakeTokenPlaceholder")
                .withAuthenticationMode(AuthenticationMode.MSI))
            .create();
    }
}
