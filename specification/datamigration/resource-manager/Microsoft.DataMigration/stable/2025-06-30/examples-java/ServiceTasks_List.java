
/**
 * Samples for ServiceTasks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/ServiceTasks_List
     * .json
     */
    /**
     * Sample code: ServiceTasks_List.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void serviceTasksList(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.serviceTasks().list("DmsSdkRg", "DmsSdkService", null, com.azure.core.util.Context.NONE);
    }
}
