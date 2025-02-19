
/**
 * Samples for StorageAccounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/
     * StorageAccounts_Delete.json
     */
    /**
     * Sample code: Removes an Azure Storage account.
     * 
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void
        removesAnAzureStorageAccount(com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager.storageAccounts().deleteWithResponse("contosorg", "contosoadla", "test_storage",
            com.azure.core.util.Context.NONE);
    }
}
