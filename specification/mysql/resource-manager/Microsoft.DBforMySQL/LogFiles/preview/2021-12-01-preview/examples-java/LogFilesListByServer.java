/** Samples for LogFiles ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/LogFiles/preview/2021-12-01-preview/examples/LogFilesListByServer.json
     */
    /**
     * Sample code: List all server log files for a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listAllServerLogFilesForAServer(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.logFiles().listByServer("testrg", "mysqltestsvc1", com.azure.core.util.Context.NONE);
    }
}
