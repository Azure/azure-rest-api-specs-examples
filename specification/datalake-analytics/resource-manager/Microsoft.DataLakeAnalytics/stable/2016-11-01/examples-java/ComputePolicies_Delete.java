/** Samples for ComputePolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/ComputePolicies_Delete.json
     */
    /**
     * Sample code: Deletes the specified compute policy from the adla account.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void deletesTheSpecifiedComputePolicyFromTheAdlaAccount(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager
            .computePolicies()
            .deleteWithResponse("contosorg", "contosoadla", "test_policy", com.azure.core.util.Context.NONE);
    }
}
