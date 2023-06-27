/** Samples for LogFiles ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/LogFilesListByServer.json
     */
    /**
     * Sample code: List all server log files for a server.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAllServerLogFilesForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.logFiles().listByServer("testrg", "postgresqltestsvc1", com.azure.core.util.Context.NONE);
    }
}
