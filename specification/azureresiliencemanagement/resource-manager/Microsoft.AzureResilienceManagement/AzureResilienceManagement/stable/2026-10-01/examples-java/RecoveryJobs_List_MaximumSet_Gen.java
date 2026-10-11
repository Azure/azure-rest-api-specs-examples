
/**
 * Samples for RecoveryJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryJobs_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryJobs_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        recoveryJobsListMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryJobs().list("sampleServiceGroupName", "samplePlanName", com.azure.core.util.Context.NONE);
    }
}
