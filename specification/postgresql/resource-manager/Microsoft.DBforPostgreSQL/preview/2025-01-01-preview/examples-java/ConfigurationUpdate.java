
/**
 * Samples for Configurations Put.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * ConfigurationUpdate.json
     */
    /**
     * Sample code: Update a user configuration.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        updateAUserConfiguration(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.configurations().define("constraint_exclusion").withExistingFlexibleServer("testrg", "testserver")
            .withValue("on").withSource("user-override").create();
    }
}
