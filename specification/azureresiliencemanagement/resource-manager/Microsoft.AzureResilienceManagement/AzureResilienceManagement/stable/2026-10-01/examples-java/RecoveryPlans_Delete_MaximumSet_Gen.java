
/**
 * Samples for RecoveryPlans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlans_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlans_Delete_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlansDeleteMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlans().delete("sampleServiceGroupName", "samplePlanName", com.azure.core.util.Context.NONE);
    }
}
