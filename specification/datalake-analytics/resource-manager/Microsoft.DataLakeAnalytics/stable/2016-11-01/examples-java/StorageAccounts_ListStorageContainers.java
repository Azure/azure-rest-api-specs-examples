/** Samples for StorageAccounts ListStorageContainers. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/StorageAccounts_ListStorageContainers.json
     */
    /**
     * Sample code: Lists the Azure Storage containers.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void listsTheAzureStorageContainers(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .storageAccounts()
            .listStorageContainers("contosorg", "contosoadla", "test_storage", com.azure.core.util.Context.NONE);
    }
}
