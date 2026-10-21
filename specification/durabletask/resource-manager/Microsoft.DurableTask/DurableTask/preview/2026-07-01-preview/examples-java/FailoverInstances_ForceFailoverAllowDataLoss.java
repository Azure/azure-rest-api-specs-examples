
/**
 * Samples for FailoverInstances ForceFailoverAllowDataLoss.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/FailoverInstances_ForceFailoverAllowDataLoss.json
     */
    /**
     * Sample code: FailoverInstances_ForceFailoverAllowDataLoss.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        failoverInstancesForceFailoverAllowDataLoss(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.failoverInstances().forceFailoverAllowDataLoss("rgdurabletask", "testscheduler",
            com.azure.core.util.Context.NONE);
    }
}
