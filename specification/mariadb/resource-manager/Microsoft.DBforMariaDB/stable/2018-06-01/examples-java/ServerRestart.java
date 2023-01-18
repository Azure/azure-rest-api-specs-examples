/** Samples for Servers Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerRestart.json
     */
    /**
     * Sample code: ServerRestart.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void serverRestart(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.servers().restart("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
