
/**
 * Samples for TaskQueues ListByTaskHub.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TaskQueues_ListByTaskHub.json
     */
    /**
     * Sample code: TaskQueues_ListByTaskHub.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskQueuesListByTaskHub(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskQueues().listByTaskHub("rgopenapi", "testscheduler", "testtaskhub",
            com.azure.core.util.Context.NONE);
    }
}
