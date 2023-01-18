/** Samples for Databases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/DatabaseGet.json
     */
    /**
     * Sample code: DatabaseGet.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void databaseGet(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.databases().getWithResponse("TestGroup", "testserver", "db1", com.azure.core.util.Context.NONE);
    }
}
