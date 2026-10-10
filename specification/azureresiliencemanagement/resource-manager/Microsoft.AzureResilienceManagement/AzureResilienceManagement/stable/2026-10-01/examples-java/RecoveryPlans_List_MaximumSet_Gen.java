
/**
 * Samples for RecoveryPlans List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RecoveryPlans_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: RecoveryPlans_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void recoveryPlansListMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.recoveryPlans().list("sampleServiceGroupName", "jfpmvvhtt", 44, com.azure.core.util.Context.NONE);
    }
}
