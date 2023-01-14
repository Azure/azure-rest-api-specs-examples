/** Samples for DataLakeStoreAccounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/DataLakeStoreAccounts_Delete.json
     */
    /**
     * Sample code: Removes the specified Data Lake Store account.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void removesTheSpecifiedDataLakeStoreAccount(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .dataLakeStoreAccounts()
            .deleteWithResponse("contosorg", "contosoadla", "test_adls_account", com.azure.core.util.Context.NONE);
    }
}
