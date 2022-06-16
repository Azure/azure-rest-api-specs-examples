import com.azure.core.util.Context;

/** Samples for LogFiles ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/LogFileListByServer.json
     */
    /**
     * Sample code: LogFileList.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void logFileList(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.logFiles().listByServer("TestGroup", "testserver", Context.NONE);
    }
}
