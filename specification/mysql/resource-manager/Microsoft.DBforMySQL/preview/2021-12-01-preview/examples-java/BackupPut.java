/** Samples for Backups Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/preview/2021-12-01-preview/examples/BackupPut.json
     */
    /**
     * Sample code: Create backup for a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void createBackupForAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.backups().putWithResponse("TestGroup", "mysqltestserver", "mybackup", com.azure.core.util.Context.NONE);
    }
}
