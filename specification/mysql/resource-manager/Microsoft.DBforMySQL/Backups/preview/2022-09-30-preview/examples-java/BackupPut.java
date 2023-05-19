/** Samples for Backups Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/preview/2022-09-30-preview/examples/BackupPut.json
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
