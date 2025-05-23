
/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/DatabaseCreate.json
     */
    /**
     * Sample code: DatabaseCreate.
     * 
     * @param manager Entry point to MariaDBManager.
     */
    public static void databaseCreate(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager.databases().define("db1").withExistingServer("TestGroup", "testserver").withCharset("utf8")
            .withCollation("utf8_general_ci").create();
    }
}
