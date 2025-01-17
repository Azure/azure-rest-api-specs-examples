
import com.azure.resourcemanager.postgresqlflexibleserver.models.CheckNameAvailabilityRequest;

/**
 * Samples for CheckNameAvailabilityWithLocation Execute.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * CheckNameAvailabilityLocationBased.json
     */
    /**
     * Sample code: NameAvailability.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void nameAvailability(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.checkNameAvailabilityWithLocations().executeWithResponse("westus",
            new CheckNameAvailabilityRequest().withName("name1").withType("Microsoft.DBforPostgreSQL/flexibleServers"),
            com.azure.core.util.Context.NONE);
    }
}
