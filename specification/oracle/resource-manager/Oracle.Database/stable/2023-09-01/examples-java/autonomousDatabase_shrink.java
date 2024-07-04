
/**
 * Samples for AutonomousDatabases Shrink.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabase_shrink.json
     */
    /**
     * Sample code: Perform shrink action on Autonomous Database.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void performShrinkActionOnAutonomousDatabase(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().shrink("rg000", "databasedb1", com.azure.core.util.Context.NONE);
    }
}
