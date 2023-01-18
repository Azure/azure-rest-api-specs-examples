import com.azure.resourcemanager.mariadb.models.NameAvailabilityRequest;

/** Samples for CheckNameAvailability Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/CheckNameAvailability.json
     */
    /**
     * Sample code: NameAvailability.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void nameAvailability(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        manager
            .checkNameAvailabilities()
            .executeWithResponse(
                new NameAvailabilityRequest().withName("name1").withType("Microsoft.DBforMariaDB"),
                com.azure.core.util.Context.NONE);
    }
}
