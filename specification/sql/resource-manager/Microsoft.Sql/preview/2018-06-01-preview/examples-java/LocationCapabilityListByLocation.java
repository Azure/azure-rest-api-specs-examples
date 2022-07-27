import com.azure.core.util.Context;

/** Samples for Capabilities ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/LocationCapabilityListByLocation.json
     */
    /**
     * Sample code: List subscription capabilities in the given location.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSubscriptionCapabilitiesInTheGivenLocation(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getCapabilities()
            .listByLocationWithResponse("eastus2euap", null, Context.NONE);
    }
}
