
/**
 * Samples for Capabilities ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/LocationCapabilityListByLocation.json
     */
    /**
     * Sample code: List subscription capabilities in the given location.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listSubscriptionCapabilitiesInTheGivenLocation(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getCapabilities().listByLocationWithResponse("eastus", null,
            com.azure.core.util.Context.NONE);
    }
}
