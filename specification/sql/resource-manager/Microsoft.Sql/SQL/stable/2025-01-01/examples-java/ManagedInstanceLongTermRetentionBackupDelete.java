
/**
 * Samples for LongTermRetentionManagedInstanceBackups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceLongTermRetentionBackupDelete.json
     */
    /**
     * Sample code: Delete the long term retention backup.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteTheLongTermRetentionBackup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionManagedInstanceBackups().delete("japaneast", "testInstance",
            "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000;Hot",
            com.azure.core.util.Context.NONE);
    }
}
