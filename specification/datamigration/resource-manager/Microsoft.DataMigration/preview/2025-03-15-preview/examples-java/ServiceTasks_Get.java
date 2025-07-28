
/**
 * Samples for ServiceTasks Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * ServiceTasks_Get.json
     */
    /**
     * Sample code: Tasks_Get.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void tasksGet(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.serviceTasks().getWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkTask", null,
            com.azure.core.util.Context.NONE);
    }
}
