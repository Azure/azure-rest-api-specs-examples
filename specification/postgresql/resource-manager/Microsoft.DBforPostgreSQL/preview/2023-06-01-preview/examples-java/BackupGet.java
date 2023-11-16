/** Samples for Backups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/BackupGet.json
     */
    /**
     * Sample code: Get a backup for a server.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getABackupForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .backups()
            .getWithResponse(
                "TestGroup", "postgresqltestserver", "daily_20210615T160516", com.azure.core.util.Context.NONE);
    }
}
