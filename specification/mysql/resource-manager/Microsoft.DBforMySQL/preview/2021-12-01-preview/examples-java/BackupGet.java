/** Samples for Backups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/preview/2021-12-01-preview/examples/BackupGet.json
     */
    /**
     * Sample code: Get a backup for a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getABackupForAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager
            .backups()
            .getWithResponse("TestGroup", "mysqltestserver", "daily_20210615T160516", com.azure.core.util.Context.NONE);
    }
}
