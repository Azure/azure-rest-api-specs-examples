
import com.azure.resourcemanager.durabletask.models.TaskQueueProperties;
import com.azure.resourcemanager.durabletask.models.TaskQueueThroughput;

/**
 * Samples for TaskQueues CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TaskQueues_CreateOrUpdate.json
     */
    /**
     * Sample code: TaskQueues_CreateOrUpdate.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskQueuesCreateOrUpdate(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskQueues().define("QueueA").withExistingTaskHub("rgopenapi", "testscheduler", "testtaskhub")
            .withProperties(
                new TaskQueueProperties().withThroughput(new TaskQueueThroughput().withActionsPerSecond(40)))
            .create();
    }
}
