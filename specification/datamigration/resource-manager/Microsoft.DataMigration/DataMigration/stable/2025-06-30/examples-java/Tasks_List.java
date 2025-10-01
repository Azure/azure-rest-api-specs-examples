
/**
 * Samples for Tasks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * Tasks_List.json
     */
    /**
     * Sample code: Tasks_List.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void tasksList(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.tasks().list("DmsSdkRg", "DmsSdkService", "DmsSdkProject", null, com.azure.core.util.Context.NONE);
    }
}
