
/**
 * Samples for TaskHubs ListByScheduler.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/TaskHubs_ListByScheduler.json
     */
    /**
     * Sample code: TaskHubs_ListByScheduler.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskHubsListByScheduler(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskHubs().listByScheduler("rgopenapi", "testscheduler", com.azure.core.util.Context.NONE);
    }
}
