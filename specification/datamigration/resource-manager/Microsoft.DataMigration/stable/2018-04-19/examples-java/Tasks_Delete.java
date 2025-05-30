
/**
 * Samples for Tasks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Tasks_Delete.json
     */
    /**
     * Sample code: Tasks_Delete.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void tasksDelete(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.tasks().deleteWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkProject", "DmsSdkTask", null,
            com.azure.core.util.Context.NONE);
    }
}
