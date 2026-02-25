
/**
 * Samples for CapabilitiesByLocation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/CapabilitiesByLocationList.json
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
