import com.azure.core.util.Context;
import com.azure.resourcemanager.mysqlflexibleserver.fluent.models.ConfigurationInner;
import com.azure.resourcemanager.mysqlflexibleserver.models.ConfigurationSource;

/** Samples for Configurations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationUpdate.json
     */
    /**
     * Sample code: Update a user configuration.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void updateAUserConfiguration(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager
            .configurations()
            .update(
                "testrg",
                "testserver",
                "event_scheduler",
                new ConfigurationInner().withValue("on").withSource(ConfigurationSource.USER_OVERRIDE),
                Context.NONE);
    }
}
