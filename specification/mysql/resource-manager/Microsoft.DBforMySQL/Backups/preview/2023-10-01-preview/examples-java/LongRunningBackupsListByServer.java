
/**
 * Samples for LongRunningBackupsOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/preview/2023-10-01-preview/examples/
     * LongRunningBackupsListByServer.json
     */
    /**
     * Sample code: List backups for a server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void listBackupsForAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.longRunningBackupsOperations().list("TestGroup", "mysqltestserver", com.azure.core.util.Context.NONE);
    }
}
