
import com.azure.resourcemanager.postgresqlflexibleserver.models.CheckNameAvailabilityRequest;

/**
 * Samples for CheckNameAvailability Execute.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * CheckNameAvailability.json
     */
    /**
     * Sample code: NameAvailability.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void nameAvailability(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.checkNameAvailabilities().executeWithResponse(
            new CheckNameAvailabilityRequest().withName("name1").withType("Microsoft.DBforPostgreSQL/flexibleServers"),
            com.azure.core.util.Context.NONE);
    }
}
