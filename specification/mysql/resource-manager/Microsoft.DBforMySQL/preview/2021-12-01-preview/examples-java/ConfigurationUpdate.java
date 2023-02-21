import com.azure.resourcemanager.mysqlflexibleserver.fluent.models.ConfigurationInner;
import com.azure.resourcemanager.mysqlflexibleserver.models.ConfigurationSource;

/** Samples for Configurations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/preview/2021-12-01-preview/examples/ConfigurationUpdate.json
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
                com.azure.core.util.Context.NONE);
    }
}
