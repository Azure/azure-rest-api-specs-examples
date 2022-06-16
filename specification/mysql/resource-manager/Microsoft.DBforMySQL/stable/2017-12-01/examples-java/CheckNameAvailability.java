import com.azure.core.util.Context;
import com.azure.resourcemanager.mysql.models.NameAvailabilityRequest;

/** Samples for CheckNameAvailability Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/CheckNameAvailability.json
     */
    /**
     * Sample code: NameAvailability.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void nameAvailability(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .checkNameAvailabilities()
            .executeWithResponse(
                new NameAvailabilityRequest().withName("name1").withType("Microsoft.DBforMySQL"), Context.NONE);
    }
}
