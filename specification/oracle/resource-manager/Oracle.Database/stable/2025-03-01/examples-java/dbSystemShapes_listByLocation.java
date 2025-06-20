
/**
 * Samples for DbSystemShapes ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/dbSystemShapes_listByLocation.json
     */
    /**
     * Sample code: DbSystemShapes_listByLocation.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        dbSystemShapesListByLocation(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbSystemShapes().listByLocation("eastus", null, com.azure.core.util.Context.NONE);
    }
}
