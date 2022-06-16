import com.azure.core.util.Context;

/** Samples for Configurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationGet.json
     */
    /**
     * Sample code: Get a configuration.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getAConfiguration(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.configurations().getWithResponse("TestGroup", "testserver", "event_scheduler", Context.NONE);
    }
}
