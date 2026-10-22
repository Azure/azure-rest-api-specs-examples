
import com.azure.resourcemanager.durabletask.models.TaskQueueProperties;

/**
 * Samples for TaskQueues CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TaskQueues_CreateOrUpdate_MinimumSet.json
     */
    /**
     * Sample code: TaskQueues_CreateOrUpdate_MinimumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        taskQueuesCreateOrUpdateMinimumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskQueues().define("QueueA").withExistingTaskHub("rgopenapi", "testscheduler", "testtaskhub")
            .withProperties(new TaskQueueProperties()).create();
    }
}
