
/**
 * Samples for AutonomousDatabaseBackups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseBackup_delete.
     * json
     */
    /**
     * Sample code: AutonomousDatabaseBackups_Delete.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabaseBackupsDelete(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseBackups().delete("rg000", "databasedb1", "1711644130",
            com.azure.core.util.Context.NONE);
    }
}
