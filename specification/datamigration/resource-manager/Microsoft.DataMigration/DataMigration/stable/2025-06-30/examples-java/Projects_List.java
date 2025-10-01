
/**
 * Samples for Projects List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * Projects_List.json
     */
    /**
     * Sample code: Projects_List.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void projectsList(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.projects().list("DmsSdkRg", "DmsSdkService", com.azure.core.util.Context.NONE);
    }
}
