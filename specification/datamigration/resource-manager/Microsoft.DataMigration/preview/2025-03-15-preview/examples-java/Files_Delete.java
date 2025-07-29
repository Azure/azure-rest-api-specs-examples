
/**
 * Samples for Files Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * Files_Delete.json
     */
    /**
     * Sample code: Files_Delete.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void filesDelete(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.files().deleteWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8",
            com.azure.core.util.Context.NONE);
    }
}
