
import com.azure.resourcemanager.durabletask.models.TaskHubProperties;

/**
 * Samples for TaskHubs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01-preview/TaskHubs_CreateOrUpdate.json
     */
    /**
     * Sample code: TaskHubs_CreateOrUpdate.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskHubsCreateOrUpdate(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskHubs().define("testtaskhub").withExistingScheduler("rgopenapi", "testscheduler")
            .withProperties(new TaskHubProperties()).create();
    }
}
