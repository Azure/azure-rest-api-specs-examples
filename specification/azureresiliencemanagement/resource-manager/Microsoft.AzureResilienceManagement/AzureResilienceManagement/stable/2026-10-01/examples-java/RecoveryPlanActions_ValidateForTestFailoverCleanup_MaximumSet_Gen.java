
/**
 * Samples for RecoveryPlanActions ValidateForTestFailoverCleanup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlanActions_ValidateForTestFailoverCleanup_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlanActions_ValidateForTestFailoverCleanup_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlanActionsValidateForTestFailoverCleanupMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlanActions().validateForTestFailoverCleanup("sampleServiceGroupName", "qmn", "samplePlanName",
            com.azure.core.util.Context.NONE);
    }
}
