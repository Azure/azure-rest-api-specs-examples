
/**
 * Samples for Databases Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/DatabasesCreate.
     * json
     */
    /**
     * Sample code: Create a database.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createADatabase(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.databases().define("exampledatabase")
            .withExistingFlexibleServer("exampleresourcegroup", "exampleserver").withCharset("utf8")
            .withCollation("en_US.utf8").create();
    }
}
