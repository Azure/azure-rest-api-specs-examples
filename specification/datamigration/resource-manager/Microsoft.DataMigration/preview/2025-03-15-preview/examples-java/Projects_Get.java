
/**
 * Samples for Projects Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * Projects_Get.json
     */
    /**
     * Sample code: Projects_Get.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void projectsGet(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.projects().getWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkProject",
            com.azure.core.util.Context.NONE);
    }
}
