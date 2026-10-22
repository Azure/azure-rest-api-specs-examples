
/**
 * Samples for FailoverInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/FailoverInstances_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: FailoverInstances_Delete_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        failoverInstancesDeleteMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.failoverInstances().delete("rgdurabletask", "testscheduler", com.azure.core.util.Context.NONE);
    }
}
