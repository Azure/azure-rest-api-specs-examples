
/**
 * Samples for LongRunningBackup Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/preview/2023-10-01-preview/examples/
     * LongRunningBackup.json
     */
    /**
     * Sample code: Create backup for a server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void createBackupForAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.longRunningBackups().define("testback").withExistingFlexibleServer("TestGroup", "mysqltestserver")
            .create();
    }
}
