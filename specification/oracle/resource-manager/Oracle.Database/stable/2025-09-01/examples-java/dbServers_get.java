
/**
 * Samples for DbServers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/dbServers_get.json
     */
    /**
     * Sample code: DbServers_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void dbServersGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbServers().getWithResponse("rg000", "infra1", "ocid1....aaaaaa", com.azure.core.util.Context.NONE);
    }
}
