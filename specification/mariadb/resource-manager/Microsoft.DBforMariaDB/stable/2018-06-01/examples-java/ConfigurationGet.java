/** Samples for Configurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ConfigurationGet.json
     */
    /**
     * Sample code: ConfigurationGet.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void configurationGet(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .configurations()
            .getWithResponse("TestGroup", "testserver", "event_scheduler", com.azure.core.util.Context.NONE);
    }
}
