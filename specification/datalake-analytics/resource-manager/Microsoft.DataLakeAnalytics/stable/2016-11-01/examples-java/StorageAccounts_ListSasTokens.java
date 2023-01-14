/** Samples for StorageAccounts ListSasTokens. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/StorageAccounts_ListSasTokens.json
     */
    /**
     * Sample code: Gets the SAS token.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void getsTheSASToken(com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .storageAccounts()
            .listSasTokens(
                "contosorg", "contosoadla", "test_storage", "test_container", com.azure.core.util.Context.NONE);
    }
}
