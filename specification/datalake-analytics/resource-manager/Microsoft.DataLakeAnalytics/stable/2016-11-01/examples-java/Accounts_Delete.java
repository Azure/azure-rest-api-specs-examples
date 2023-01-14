/** Samples for Accounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/Accounts_Delete.json
     */
    /**
     * Sample code: Begins the delete process for the Data Lake Analytics account object specified by the account name.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void beginsTheDeleteProcessForTheDataLakeAnalyticsAccountObjectSpecifiedByTheAccountName(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager.accounts().delete("contosorg", "contosoadla", com.azure.core.util.Context.NONE);
    }
}
