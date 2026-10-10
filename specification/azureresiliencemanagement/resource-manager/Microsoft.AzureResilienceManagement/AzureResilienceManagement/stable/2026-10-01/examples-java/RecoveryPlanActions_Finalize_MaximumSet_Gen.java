
/**
 * Samples for RecoveryPlanActions Finalize.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlanActions_Finalize_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlanActions_Finalize_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlanActionsFinalizeMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlanActions().finalize("sampleServiceGroupName", "qmn", "samplePlanName",
            com.azure.core.util.Context.NONE);
    }
}
