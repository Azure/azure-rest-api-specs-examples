
/**
 * Samples for Backups Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/BackupCreate.json
     */
    /**
     * Sample code: Create a new Backup for a flexible server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createANewBackupForAFlexibleServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backups().create("TestGroup", "postgresqltestserver", "backup_20210615T160516",
            com.azure.core.util.Context.NONE);
    }
}
