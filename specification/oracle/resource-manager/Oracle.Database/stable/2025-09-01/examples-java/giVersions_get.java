
/**
 * Samples for GiVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/giVersions_get.json
     */
    /**
     * Sample code: GiVersions_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void giVersionsGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.giVersions().getWithResponse("eastus", "19.0.0.0", com.azure.core.util.Context.NONE);
    }
}
