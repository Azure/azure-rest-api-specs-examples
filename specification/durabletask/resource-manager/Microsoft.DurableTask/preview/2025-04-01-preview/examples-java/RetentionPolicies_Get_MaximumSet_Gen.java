
/**
 * Samples for RetentionPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/RetentionPolicies_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: RetentionPolicies_Get_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        retentionPoliciesGetMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.retentionPolicies().getWithResponse("rgdurabletask", "testscheduler", com.azure.core.util.Context.NONE);
    }
}
