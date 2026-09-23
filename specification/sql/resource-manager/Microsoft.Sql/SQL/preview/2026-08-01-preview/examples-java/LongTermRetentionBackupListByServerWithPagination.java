
/**
 * Samples for LongTermRetentionBackups ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/LongTermRetentionBackupListByServerWithPagination.json
     */
    /**
     * Sample code: Get long term retention backups under the server with pagination.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getLongTermRetentionBackupsUnderTheServerWithPagination(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups().listByServer("japaneast", "testserver", null, null, null,
            null, com.azure.core.util.Context.NONE);
    }
}
