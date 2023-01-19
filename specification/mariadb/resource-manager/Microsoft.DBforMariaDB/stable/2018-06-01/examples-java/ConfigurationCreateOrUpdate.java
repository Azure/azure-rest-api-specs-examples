/** Samples for Configurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ConfigurationCreateOrUpdate.json
     */
    /**
     * Sample code: ConfigurationCreateOrUpdate.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void configurationCreateOrUpdate(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .configurations()
            .define("event_scheduler")
            .withExistingServer("TestGroup", "testserver")
            .withValue("off")
            .withSource("user-override")
            .create();
    }
}
