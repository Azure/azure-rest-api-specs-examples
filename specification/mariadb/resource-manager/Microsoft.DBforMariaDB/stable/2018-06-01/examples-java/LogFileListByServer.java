/** Samples for LogFiles ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/LogFileListByServer.json
     */
    /**
     * Sample code: LogFileList.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void logFileList(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.logFiles().listByServer("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
