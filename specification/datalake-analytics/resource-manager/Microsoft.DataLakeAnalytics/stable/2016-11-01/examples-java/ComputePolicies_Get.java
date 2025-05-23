
/**
 * Samples for ComputePolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/
     * ComputePolicies_Get.json
     */
    /**
     * Sample code: Gets the specified compute policy.
     * 
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void
        getsTheSpecifiedComputePolicy(com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager.computePolicies().getWithResponse("contosorg", "contosoadla", "test_policy",
            com.azure.core.util.Context.NONE);
    }
}
