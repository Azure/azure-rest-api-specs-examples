
/**
 * Samples for Projects List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
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
