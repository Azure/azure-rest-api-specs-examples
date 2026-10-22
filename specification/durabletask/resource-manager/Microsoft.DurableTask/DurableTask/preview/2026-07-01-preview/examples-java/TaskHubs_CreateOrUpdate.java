
import com.azure.resourcemanager.durabletask.models.TaskHubProperties;
import java.util.Arrays;

/**
 * Samples for TaskHubs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TaskHubs_CreateOrUpdate.json
     */
    /**
     * Sample code: TaskHubs_CreateOrUpdate.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskHubsCreateOrUpdate(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskHubs().define("testtaskhub").withExistingScheduler("rgopenapi", "testscheduler")
            .withProperties(new TaskHubProperties().withCapabilities(Arrays.asList("ExampleFutureCapability")))
            .create();
    }
}
