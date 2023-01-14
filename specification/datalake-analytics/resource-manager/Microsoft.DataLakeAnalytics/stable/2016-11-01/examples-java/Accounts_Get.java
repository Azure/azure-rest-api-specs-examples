/** Samples for Accounts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/Accounts_Get.json
     */
    /**
     * Sample code: Gets details of the specified Data Lake Analytics account.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void getsDetailsOfTheSpecifiedDataLakeAnalyticsAccount(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager.accounts().getByResourceGroupWithResponse("contosorg", "contosoadla", com.azure.core.util.Context.NONE);
    }
}
