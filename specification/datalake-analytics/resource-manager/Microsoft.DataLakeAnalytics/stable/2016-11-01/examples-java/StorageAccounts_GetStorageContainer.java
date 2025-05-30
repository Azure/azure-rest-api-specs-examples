
/**
 * Samples for StorageAccounts GetStorageContainer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/
     * StorageAccounts_GetStorageContainer.json
     */
    /**
     * Sample code: Gets the specified Azure Storage container.
     * 
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void getsTheSpecifiedAzureStorageContainer(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager.storageAccounts().getStorageContainerWithResponse("contosorg", "contosoadla", "test_storage",
            "test_container", com.azure.core.util.Context.NONE);
    }
}
