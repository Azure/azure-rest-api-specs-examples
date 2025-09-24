
/**
 * Samples for ServiceTasks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * ServiceTasks_Delete.json
     */
    /**
     * Sample code: Tasks_Delete.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void tasksDelete(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.serviceTasks().deleteWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkTask", null,
            com.azure.core.util.Context.NONE);
    }
}
