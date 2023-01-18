/** Samples for Servers Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2020-01-01/examples/ServerStop.json
     */
    /**
     * Sample code: ServerStop.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void serverStop(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.servers().stop("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
