
/**
 * Samples for DbNodes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/dbNodes_get.json
     */
    /**
     * Sample code: Get DbNode.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void getDbNode(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbNodes().getWithResponse("rg000", "cluster1", "ocid1....aaaaaa", com.azure.core.util.Context.NONE);
    }
}
