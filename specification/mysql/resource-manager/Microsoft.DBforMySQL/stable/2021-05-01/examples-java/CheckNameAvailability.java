import com.azure.core.util.Context;
import com.azure.resourcemanager.mysqlflexibleserver.models.NameAvailabilityRequest;

/** Samples for CheckNameAvailability Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/CheckNameAvailability.json
     */
    /**
     * Sample code: Check name availability.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager
            .checkNameAvailabilities()
            .executeWithResponse(
                "SouthEastAsia",
                new NameAvailabilityRequest().withName("name1").withType("Microsoft.DBforMySQL/flexibleServers"),
                Context.NONE);
    }
}
