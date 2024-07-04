
/**
 * Samples for DbSystemShapes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dbSystemShapes_get.json
     */
    /**
     * Sample code: Get a DbSystemShape by name.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void getADbSystemShapeByName(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbSystemShapes().getWithResponse("eastus", "EXADATA.X9M", com.azure.core.util.Context.NONE);
    }
}
