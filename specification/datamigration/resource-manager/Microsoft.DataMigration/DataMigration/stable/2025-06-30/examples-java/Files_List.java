
/**
 * Samples for Files List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * Files_List.json
     */
    /**
     * Sample code: Files_List.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void filesList(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.files().list("DmsSdkRg", "DmsSdkService", "DmsSdkProject", com.azure.core.util.Context.NONE);
    }
}
