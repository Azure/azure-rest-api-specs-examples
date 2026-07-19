
/**
 * Samples for LongTermRetentionBackups ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/LongTermRetentionBackupListByServer.json
     */
    /**
     * Sample code: Get all long term retention backups under the server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAllLongTermRetentionBackupsUnderTheServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups().listByServer("japaneast", "testserver", null, null,
            com.azure.core.util.Context.NONE);
    }
}
