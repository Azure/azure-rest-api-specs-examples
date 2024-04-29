
import com.azure.resourcemanager.mysqlflexibleserver.models.ConfigurationSource;

/**
 * Samples for Configurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Configurations/preview/2023-06-01-preview/examples/
     * ConfigurationCreateOrUpdate.json
     */
    /**
     * Sample code: ConfigurationCreateOrUpdate.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void configurationCreateOrUpdate(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.configurations().define("event_scheduler").withExistingFlexibleServer("TestGroup", "testserver")
            .withValue("off").withSource(ConfigurationSource.USER_OVERRIDE).create();
    }
}
