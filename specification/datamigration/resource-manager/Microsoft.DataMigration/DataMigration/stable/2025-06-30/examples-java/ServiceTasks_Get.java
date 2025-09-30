
/**
 * Samples for ServiceTasks Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
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
