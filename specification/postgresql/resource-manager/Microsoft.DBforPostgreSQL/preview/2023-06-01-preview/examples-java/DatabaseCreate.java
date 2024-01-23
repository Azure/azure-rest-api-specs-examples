
/**
 * Samples for Databases Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/
     * DatabaseCreate.json
     */
    /**
     * Sample code: Create a database.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createADatabase(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.databases().define("db1").withExistingFlexibleServer("TestGroup", "testserver").withCharset("utf8")
            .withCollation("en_US.utf8").create();
    }
}
