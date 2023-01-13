/** Samples for StorageAccounts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/StorageAccounts_Get.json
     */
    /**
     * Sample code: Gets the specified Azure Storage account.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void getsTheSpecifiedAzureStorageAccount(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .storageAccounts()
            .getWithResponse("contosorg", "contosoadla", "test_storage", com.azure.core.util.Context.NONE);
    }
}
