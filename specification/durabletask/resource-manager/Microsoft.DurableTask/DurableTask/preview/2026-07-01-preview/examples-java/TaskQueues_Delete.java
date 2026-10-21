
/**
 * Samples for TaskQueues Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TaskQueues_Delete.json
     */
    /**
     * Sample code: TaskQueues_Delete.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskQueuesDelete(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskQueues().delete("rgopenapi", "testscheduler", "testtaskhub", "QueueA",
            com.azure.core.util.Context.NONE);
    }
}
