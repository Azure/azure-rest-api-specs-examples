
/**
 * Samples for TaskQueues Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TaskQueues_Get.json
     */
    /**
     * Sample code: TaskQueues_Get.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskQueuesGet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskQueues().getWithResponse("rgopenapi", "testscheduler", "testtaskhub", "QueueA",
            com.azure.core.util.Context.NONE);
    }
}
