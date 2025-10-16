
/**
 * Samples for AutonomousDatabases Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabase_delete.json
     */
    /**
     * Sample code: AutonomousDatabases_delete.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabasesDelete(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().delete("rg000", "databasedb1", com.azure.core.util.Context.NONE);
    }
}
