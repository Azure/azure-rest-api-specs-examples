
import com.azure.resourcemanager.mysqlflexibleserver.models.Configuration;
import com.azure.resourcemanager.mysqlflexibleserver.models.ConfigurationSource;

/**
 * Samples for Configurations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Configurations/preview/2023-06-01-preview/examples/
     * ConfigurationUpdate.json
     */
    /**
     * Sample code: Update a user configuration.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void updateAUserConfiguration(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        Configuration resource = manager.configurations()
            .getWithResponse("testrg", "testserver", "event_scheduler", com.azure.core.util.Context.NONE).getValue();
        resource.update().withValue("on").withSource(ConfigurationSource.USER_OVERRIDE).apply();
    }
}
