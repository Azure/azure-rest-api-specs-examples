
/**
 * Samples for RecoveryJobs Retry.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryJobs_Retry_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryJobs_Retry_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryJobsRetryMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryJobs().retry("sampleServiceGroupName", "qmn", "samplePlanName",
            "c56888ef-9ced-4001-a6d4-7145a0309bdb", com.azure.core.util.Context.NONE);
    }
}
