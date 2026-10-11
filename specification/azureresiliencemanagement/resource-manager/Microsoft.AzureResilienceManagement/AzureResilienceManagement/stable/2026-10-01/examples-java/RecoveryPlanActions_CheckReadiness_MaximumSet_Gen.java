
/**
 * Samples for RecoveryPlanActions CheckReadiness.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlanActions_CheckReadiness_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlanActions_CheckReadiness_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlanActionsCheckReadinessMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlanActions().checkReadiness("sampleServiceGroupName", "qmn", "samplePlanName",
            com.azure.core.util.Context.NONE);
    }
}
