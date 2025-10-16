
/**
 * Samples for DbNodes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/dbNodes_get.json
     */
    /**
     * Sample code: DbNodes_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void dbNodesGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbNodes().getWithResponse("rg000", "cluster1", "ocid1....aaaaaa", com.azure.core.util.Context.NONE);
    }
}
