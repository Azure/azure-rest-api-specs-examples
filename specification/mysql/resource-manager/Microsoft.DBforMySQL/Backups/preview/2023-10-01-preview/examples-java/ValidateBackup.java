
/**
 * Samples for BackupAndExport ValidateBackup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/preview/2023-10-01-preview/examples/
     * ValidateBackup.json
     */
    /**
     * Sample code: Validate Backup.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void validateBackup(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.backupAndExports().validateBackupWithResponse("TestGroup", "mysqltestserver",
            com.azure.core.util.Context.NONE);
    }
}
