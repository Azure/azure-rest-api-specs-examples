import com.azure.core.util.Context;

/** Samples for LogFiles ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/LogFileListByServer.json
     */
    /**
     * Sample code: LogFileList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void logFileList(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.logFiles().listByServer("TestGroup", "testserver", Context.NONE);
    }
}
