/** Samples for Configurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ConfigurationCreateOrUpdate.json
     */
    /**
     * Sample code: ConfigurationCreateOrUpdate.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void configurationCreateOrUpdate(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .configurations()
            .define("event_scheduler")
            .withExistingServer("TestGroup", "testserver")
            .withValue("off")
            .withSource("user-override")
            .create();
    }
}
