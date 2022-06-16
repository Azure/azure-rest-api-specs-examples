import com.azure.core.util.Context;

/** Samples for LocationBasedCapabilities Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/CapabilitiesByLocation.json
     */
    /**
     * Sample code: CapabilitiesList.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void capabilitiesList(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.locationBasedCapabilities().execute("westus", Context.NONE);
    }
}
