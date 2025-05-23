
import com.azure.resourcemanager.postgresqlflexibleserver.models.Configuration;

/**
 * Samples for Configurations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/
     * ConfigurationUpdate.json
     */
    /**
     * Sample code: Update a user configuration.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        updateAUserConfiguration(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Configuration resource = manager.configurations()
            .getWithResponse("testrg", "testserver", "event_scheduler", com.azure.core.util.Context.NONE).getValue();
        resource.update().withValue("on").withSource("user-override").apply();
    }
}
