
/**
 * Samples for AutonomousDatabaseVersions ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/
     * autonomousDatabaseVersion_listByLocation.json
     */
    /**
     * Sample code: List an autonomous versions by location.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listAnAutonomousVersionsByLocation(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseVersions().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
