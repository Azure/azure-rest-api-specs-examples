
/**
 * Samples for RecoveryPlans Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlans_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlans_Get_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        recoveryPlansGetMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlans().getWithResponse("sampleServiceGroupName", "samplePlanName",
            com.azure.core.util.Context.NONE);
    }
}
