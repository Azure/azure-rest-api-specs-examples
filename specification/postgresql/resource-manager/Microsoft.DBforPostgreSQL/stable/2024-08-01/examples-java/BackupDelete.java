
/**
 * Samples for Backups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/BackupDelete.json
     */
    /**
     * Sample code: Delete a specific backup.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        deleteASpecificBackup(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backups().delete("TestGroup", "testserver", "backup_20210615T160516", com.azure.core.util.Context.NONE);
    }
}
