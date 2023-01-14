/** Samples for DataLakeStoreAccounts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/DataLakeStoreAccounts_Get.json
     */
    /**
     * Sample code: Gets the specified Data Lake Store account details.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void getsTheSpecifiedDataLakeStoreAccountDetails(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .dataLakeStoreAccounts()
            .getWithResponse("contosorg", "contosoadla", "test_adls_account", com.azure.core.util.Context.NONE);
    }
}
