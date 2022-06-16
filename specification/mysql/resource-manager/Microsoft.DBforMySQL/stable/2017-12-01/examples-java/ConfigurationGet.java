import com.azure.core.util.Context;

/** Samples for Configurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ConfigurationGet.json
     */
    /**
     * Sample code: ConfigurationGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void configurationGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.configurations().getWithResponse("TestGroup", "testserver", "event_scheduler", Context.NONE);
    }
}
