
/**
 * Samples for RetentionPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/RetentionPolicies_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RetentionPolicies_Delete_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        retentionPoliciesDeleteMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.retentionPolicies().delete("rgdurabletask", "testcheduler", com.azure.core.util.Context.NONE);
    }
}
