/** Samples for DataLakeStoreAccounts ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/DataLakeStoreAccounts_ListByAccount.json
     */
    /**
     * Sample code: Gets the first page of Data Lake Store accounts linked to the specified Data Lake Analytics account.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void getsTheFirstPageOfDataLakeStoreAccountsLinkedToTheSpecifiedDataLakeAnalyticsAccount(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .dataLakeStoreAccounts()
            .listByAccount(
                "contosorg",
                "contosoadla",
                "test_filter",
                1,
                1,
                "test_select",
                "test_orderby",
                false,
                com.azure.core.util.Context.NONE);
    }
}
