/** Samples for Backups ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/preview/2022-09-30-preview/examples/BackupsListByServer.json
     */
    /**
     * Sample code: List backups for a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listBackupsForAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.backups().listByServer("TestGroup", "mysqltestserver", com.azure.core.util.Context.NONE);
    }
}
