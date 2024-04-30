
/**
 * Samples for LongRunningBackupsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/preview/2023-10-01-preview/examples/
     * LongRunningBackupGet.json
     */
    /**
     * Sample code: Get a backup for a server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void getABackupForAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.longRunningBackupsOperations().getWithResponse("TestGroup", "mysqltestserver", "daily_20210615T160516",
            com.azure.core.util.Context.NONE);
    }
}
