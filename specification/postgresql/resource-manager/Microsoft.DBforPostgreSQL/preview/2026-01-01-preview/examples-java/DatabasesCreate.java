
/**
 * Samples for Databases Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/DatabasesCreate.json
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
