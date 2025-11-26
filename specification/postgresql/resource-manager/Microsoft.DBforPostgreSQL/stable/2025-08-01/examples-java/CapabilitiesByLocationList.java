
/**
 * Samples for CapabilitiesByLocation List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * CapabilitiesByLocationList.json
     */
    /**
     * Sample code: List the capabilities available in a given location for a specific subscription.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listTheCapabilitiesAvailableInAGivenLocationForASpecificSubscription(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.capabilitiesByLocations().list("eastus", com.azure.core.util.Context.NONE);
    }
}
