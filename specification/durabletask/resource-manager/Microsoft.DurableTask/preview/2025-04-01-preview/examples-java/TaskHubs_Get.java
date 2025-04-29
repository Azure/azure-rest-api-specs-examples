
/**
 * Samples for TaskHubs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01-preview/TaskHubs_Get.json
     */
    /**
     * Sample code: TaskHubs_Get.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void taskHubsGet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.taskHubs().getWithResponse("rgopenapi", "testscheduler", "testtaskhub",
            com.azure.core.util.Context.NONE);
    }
}
