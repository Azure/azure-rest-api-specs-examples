
/**
 * Samples for AutonomousDatabaseVersions ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabaseVersion_listByLocation.json
     */
    /**
     * Sample code: AutonomousDatabaseVersions_listByLocation.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void autonomousDatabaseVersionsListByLocation(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseVersions().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
