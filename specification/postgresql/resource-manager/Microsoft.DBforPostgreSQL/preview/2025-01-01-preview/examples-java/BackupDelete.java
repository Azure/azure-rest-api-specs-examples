
/**
 * Samples for Backups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * BackupDelete.json
     */
    /**
     * Sample code: Delete a specific backup.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        deleteASpecificBackup(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backups().delete("TestGroup", "testserver", "backup_20250303T160516", com.azure.core.util.Context.NONE);
    }
}
