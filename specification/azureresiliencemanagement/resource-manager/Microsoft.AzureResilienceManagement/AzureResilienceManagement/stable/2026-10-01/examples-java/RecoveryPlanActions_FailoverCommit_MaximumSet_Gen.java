
/**
 * Samples for RecoveryPlanActions FailoverCommit.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlanActions_FailoverCommit_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlanActions_FailoverCommit_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlanActionsFailoverCommitMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlanActions().failoverCommit("sampleServiceGroupName", "qmn", "samplePlanName",
            com.azure.core.util.Context.NONE);
    }
}
