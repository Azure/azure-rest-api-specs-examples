
/**
 * Samples for GiVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/giVersions_get.json
     */
    /**
     * Sample code: Get a GiVersion by name.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void getAGiVersionByName(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.giVersions().getWithResponse("eastus", "19.0.0.0", com.azure.core.util.Context.NONE);
    }
}
