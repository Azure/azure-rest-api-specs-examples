
/**
 * Samples for SystemVersions ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/systemVersions_listByLocation.
     * json
     */
    /**
     * Sample code: systemVersions_listByLocation.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        systemVersionsListByLocation(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.systemVersions().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
