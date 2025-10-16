
/**
 * Samples for DbSystemShapes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/dbSystemShapes_get.json
     */
    /**
     * Sample code: DbSystemShapes_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void dbSystemShapesGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbSystemShapes().getWithResponse("eastus", "EXADATA.X9M", com.azure.core.util.Context.NONE);
    }
}
