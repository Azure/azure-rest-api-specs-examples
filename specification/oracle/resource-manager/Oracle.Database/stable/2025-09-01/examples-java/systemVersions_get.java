
/**
 * Samples for SystemVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/systemVersions_get.json
     */
    /**
     * Sample code: systemVersions_listSystemVersions.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        systemVersionsListSystemVersions(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.systemVersions().getWithResponse("eastus", "22.x", com.azure.core.util.Context.NONE);
    }
}
