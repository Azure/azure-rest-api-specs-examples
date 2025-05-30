
/**
 * Samples for Accounts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/
     * Accounts_ListByResourceGroup.json
     */
    /**
     * Sample code: Gets the first page of Data Lake Analytics accounts, if any, within a specific resource group. This
     * includes a link to the next page, if any.
     * 
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void
        getsTheFirstPageOfDataLakeAnalyticsAccountsIfAnyWithinASpecificResourceGroupThisIncludesALinkToTheNextPageIfAny(
            com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager.accounts().listByResourceGroup("contosorg", "test_filter", 1, 1, "test_select", "test_orderby", false,
            com.azure.core.util.Context.NONE);
    }
}
