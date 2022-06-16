import com.azure.core.util.Context;

/** Samples for Backups ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/BackupsListByServer.json
     */
    /**
     * Sample code: List backups for a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listBackupsForAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.backups().listByServer("TestGroup", "mysqltestserver", Context.NONE);
    }
}
