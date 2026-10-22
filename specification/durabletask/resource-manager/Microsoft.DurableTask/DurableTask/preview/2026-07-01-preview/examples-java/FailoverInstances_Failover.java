
/**
 * Samples for FailoverInstances Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/FailoverInstances_Failover.json
     */
    /**
     * Sample code: FailoverInstances_Failover.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void failoverInstancesFailover(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.failoverInstances().failover("rgdurabletask", "testscheduler", com.azure.core.util.Context.NONE);
    }
}
