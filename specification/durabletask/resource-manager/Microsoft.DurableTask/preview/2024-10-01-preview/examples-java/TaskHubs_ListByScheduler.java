
/**
 * Samples for TaskHubs ListByScheduler.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01-preview/TaskHubs_ListByScheduler.json
     */
    /**
     * Sample code: TaskHubs_ListByScheduler.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskHubsListByScheduler(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskHubs().listByScheduler("rgopenapi", "testtaskhub", com.azure.core.util.Context.NONE);
    }
}
