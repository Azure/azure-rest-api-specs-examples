import com.azure.core.util.Context;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Configuration;

/** Samples for Configurations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/ConfigurationUpdate.json
     */
    /**
     * Sample code: Update a user configuration.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void updateAUserConfiguration(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Configuration resource =
            manager
                .configurations()
                .getWithResponse("testrg", "testserver", "event_scheduler", Context.NONE)
                .getValue();
        resource.update().withValue("on").withSource("user-override").apply();
    }
}
