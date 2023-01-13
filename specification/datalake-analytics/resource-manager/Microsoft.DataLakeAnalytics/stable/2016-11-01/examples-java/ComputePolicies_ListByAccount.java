/** Samples for ComputePolicies ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/ComputePolicies_ListByAccount.json
     */
    /**
     * Sample code: Lists the compute policies within the adla account.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void listsTheComputePoliciesWithinTheAdlaAccount(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager.computePolicies().listByAccount("contosorg", "contosoadla", com.azure.core.util.Context.NONE);
    }
}
