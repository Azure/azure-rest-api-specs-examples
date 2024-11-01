
/**
 * Samples for Outputs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Get_DataWarehouse.json
     */
    /**
     * Sample code: Get an Azure Data Warehouse output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        getAnAzureDataWarehouseOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg", "sjName", "dwOutput", com.azure.core.util.Context.NONE);
    }
}
