
/**
 * Samples for Files Read.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/Files_Read.json
     */
    /**
     * Sample code: Files_List.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void filesList(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.files().readWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8",
            com.azure.core.util.Context.NONE);
    }
}
