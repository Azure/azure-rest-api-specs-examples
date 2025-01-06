
/**
 * Samples for Configurations Put.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * ConfigurationUpdate.json
     */
    /**
     * Sample code: Update a user configuration.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        updateAUserConfiguration(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.configurations().define("event_scheduler").withExistingFlexibleServer("testrg", "testserver")
            .withValue("on").withSource("user-override").create();
    }
}
