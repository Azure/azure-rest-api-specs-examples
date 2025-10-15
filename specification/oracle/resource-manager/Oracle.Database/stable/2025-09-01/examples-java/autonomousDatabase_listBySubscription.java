
/**
 * Samples for AutonomousDatabases List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabase_listBySubscription.json
     */
    /**
     * Sample code: AutonomousDatabases_listBySubscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabasesListBySubscription(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().list(com.azure.core.util.Context.NONE);
    }
}
