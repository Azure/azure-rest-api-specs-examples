/** Samples for Servers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerListByResourceGroup.json
     */
    /**
     * Sample code: ServerListByResourceGroup.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void serverListByResourceGroup(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.servers().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
