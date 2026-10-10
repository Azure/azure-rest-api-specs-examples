
/**
 * Samples for RecoveryPlanActions ValidateForFailoverCommit.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlanActions_ValidateForFailoverCommit_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlanActions_ValidateForFailoverCommit_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlanActionsValidateForFailoverCommitMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlanActions().validateForFailoverCommit("sampleServiceGroupName", "qmn", "samplePlanName",
            com.azure.core.util.Context.NONE);
    }
}
