
/**
 * Samples for TaskHubs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/TaskHubs_Delete.json
     */
    /**
     * Sample code: TaskHubs_Delete.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskHubsDelete(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskHubs().delete("rgopenapi", "testscheduler", "testtaskhub", com.azure.core.util.Context.NONE);
    }
}
