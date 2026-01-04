
/**
 * Samples for ExpressRoutePortsLocations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRoutePortsLocationGet.
     * json
     */
    /**
     * Sample code: ExpressRoutePortsLocationGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRoutePortsLocationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRoutePortsLocations().getWithResponse("locationName",
            com.azure.core.util.Context.NONE);
    }
}
