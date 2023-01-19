/** Samples for Databases ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/DatabaseListByServer.json
     */
    /**
     * Sample code: DatabaseList.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void databaseList(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.databases().listByServer("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
