
/**
 * Samples for GiVersions ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/giVersions_listByLocation.json
     */
    /**
     * Sample code: List GiVersions by location.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listGiVersionsByLocation(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.giVersions().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
