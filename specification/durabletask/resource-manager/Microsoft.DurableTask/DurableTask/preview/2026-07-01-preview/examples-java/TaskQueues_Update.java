
import com.azure.resourcemanager.durabletask.models.TaskQueue;
import com.azure.resourcemanager.durabletask.models.TaskQueueProperties;
import com.azure.resourcemanager.durabletask.models.TaskQueueThroughput;

/**
 * Samples for TaskQueues Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TaskQueues_Update.json
     */
    /**
     * Sample code: TaskQueues_Update.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskQueuesUpdate(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        TaskQueue resource = manager.taskQueues()
            .getWithResponse("rgopenapi", "testscheduler", "testtaskhub", "QueueA", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(
            new TaskQueueProperties().withThroughput(new TaskQueueThroughput().withActionsPerSecond(80))).apply();
    }
}
