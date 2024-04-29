
/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Databases/preview/2023-06-01-preview/examples/
     * DatabaseCreate.json
     */
    /**
     * Sample code: Create a database.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void createADatabase(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.databases().define("db1").withExistingFlexibleServer("TestGroup", "testserver").withCharset("utf8")
            .withCollation("utf8_general_ci").create();
    }
}
