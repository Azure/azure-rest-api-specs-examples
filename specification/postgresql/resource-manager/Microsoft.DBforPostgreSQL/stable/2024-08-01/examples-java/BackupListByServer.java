
/**
 * Samples for Backups ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/BackupListByServer
     * .json
     */
    /**
     * Sample code: List backups for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        listBackupsForAServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backups().listByServer("TestGroup", "postgresqltestserver", com.azure.core.util.Context.NONE);
    }
}
