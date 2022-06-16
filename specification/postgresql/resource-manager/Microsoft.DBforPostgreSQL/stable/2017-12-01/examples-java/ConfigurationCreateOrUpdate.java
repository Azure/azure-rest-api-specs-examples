/** Samples for Configurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ConfigurationCreateOrUpdate.json
     */
    /**
     * Sample code: ConfigurationCreateOrUpdate.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void configurationCreateOrUpdate(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .configurations()
            .define("array_nulls")
            .withExistingServer("TestGroup", "testserver")
            .withValue("off")
            .withSource("user-override")
            .create();
    }
}
