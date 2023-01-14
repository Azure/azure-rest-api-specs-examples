import com.azure.resourcemanager.datalakeanalytics.models.ComputePolicy;

/** Samples for ComputePolicies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/ComputePolicies_Update.json
     */
    /**
     * Sample code: Updates the specified compute policy.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void updatesTheSpecifiedComputePolicy(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        ComputePolicy resource =
            manager
                .computePolicies()
                .getWithResponse("contosorg", "contosoadla", "test_policy", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withMaxDegreeOfParallelismPerJob(11).withMinPriorityPerJob(31).apply();
    }
}
