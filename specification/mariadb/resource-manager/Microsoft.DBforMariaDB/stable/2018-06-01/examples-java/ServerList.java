/** Samples for Servers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerList.json
     */
    /**
     * Sample code: ServerList.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void serverList(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.servers().list(com.azure.core.util.Context.NONE);
    }
}
