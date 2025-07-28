
/**
 * Samples for Files Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * Files_Get.json
     */
    /**
     * Sample code: Files_List.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void filesList(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.files().getWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8",
            com.azure.core.util.Context.NONE);
    }
}
