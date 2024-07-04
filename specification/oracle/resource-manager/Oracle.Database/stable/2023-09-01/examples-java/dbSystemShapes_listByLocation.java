
/**
 * Samples for DbSystemShapes ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dbSystemShapes_listByLocation.
     * json
     */
    /**
     * Sample code: List DbSystemShapes by location.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listDbSystemShapesByLocation(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbSystemShapes().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
