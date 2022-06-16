import com.azure.core.util.Context;

/** Samples for Backups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/BackupGet.json
     */
    /**
     * Sample code: Get a backup for a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getABackupForAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.backups().getWithResponse("TestGroup", "mysqltestserver", "daily_20210615T160516", Context.NONE);
    }
}
