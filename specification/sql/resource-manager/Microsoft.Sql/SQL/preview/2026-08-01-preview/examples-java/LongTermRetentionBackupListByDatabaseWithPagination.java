
/**
 * Samples for LongTermRetentionBackups ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/LongTermRetentionBackupListByDatabaseWithPagination.json
     */
    /**
     * Sample code: Get long term retention backups under the database with pagination.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getLongTermRetentionBackupsUnderTheDatabaseWithPagination(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups().listByDatabase("japaneast", "testserver", "testDatabase",
            null, null, com.azure.core.util.Context.NONE);
    }
}
