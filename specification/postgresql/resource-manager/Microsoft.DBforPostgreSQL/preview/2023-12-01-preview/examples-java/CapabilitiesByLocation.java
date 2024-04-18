
/**
 * Samples for LocationBasedCapabilities Execute.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/
     * CapabilitiesByLocation.json
     */
    /**
     * Sample code: CapabilitiesList.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void capabilitiesList(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.locationBasedCapabilities().execute("westus", com.azure.core.util.Context.NONE);
    }
}
