
import com.azure.resourcemanager.mysqlflexibleserver.models.NameAvailabilityRequest;

/**
 * Samples for CheckNameAvailabilityWithoutLocation Execute.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2021-05-01/examples/CheckNameAvailability
     * .json
     */
    /**
     * Sample code: Check name availability.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.checkNameAvailabilityWithoutLocations().executeWithResponse(
            new NameAvailabilityRequest().withName("name1").withType("Microsoft.DBforMySQL/flexibleServers"),
            com.azure.core.util.Context.NONE);
    }
}
