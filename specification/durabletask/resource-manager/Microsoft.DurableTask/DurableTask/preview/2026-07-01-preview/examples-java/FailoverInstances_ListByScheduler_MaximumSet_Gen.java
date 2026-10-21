
/**
 * Samples for FailoverInstances ListByScheduler.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/FailoverInstances_ListByScheduler_MaximumSet_Gen.json
     */
    /**
     * Sample code: FailoverInstances_ListByScheduler_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        failoverInstancesListBySchedulerMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.failoverInstances().listByScheduler("rgdurabletask", "testscheduler", com.azure.core.util.Context.NONE);
    }
}
