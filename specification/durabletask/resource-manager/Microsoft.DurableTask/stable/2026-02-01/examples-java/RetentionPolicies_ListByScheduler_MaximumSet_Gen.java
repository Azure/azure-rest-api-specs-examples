
/**
 * Samples for RetentionPolicies ListByScheduler.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/RetentionPolicies_ListByScheduler_MaximumSet_Gen.json
     */
    /**
     * Sample code: RetentionPolicies_ListByScheduler_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        retentionPoliciesListBySchedulerMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.retentionPolicies().listByScheduler("rgdurabletask", "myscheduler", com.azure.core.util.Context.NONE);
    }
}
