
/**
 * Samples for AutonomousDatabaseBackups ListByAutonomousDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/
     * autonomousDatabaseBackup_listByParent.json
     */
    /**
     * Sample code: AutonomousDatabaseBackups_ListByAutonomousDatabase.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void autonomousDatabaseBackupsListByAutonomousDatabase(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseBackups().listByAutonomousDatabase("rg000", "databasedb1",
            com.azure.core.util.Context.NONE);
    }
}
