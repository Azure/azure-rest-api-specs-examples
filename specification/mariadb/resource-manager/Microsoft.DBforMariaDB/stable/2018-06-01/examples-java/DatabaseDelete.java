/** Samples for Databases Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/DatabaseDelete.json
     */
    /**
     * Sample code: DatabaseDelete.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void databaseDelete(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.databases().delete("TestGroup", "testserver", "db1", com.azure.core.util.Context.NONE);
    }
}
