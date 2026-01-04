
/**
 * Samples for RouteFilters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/RouteFilterList.json
     */
    /**
     * Sample code: RouteFilterList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeFilterList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteFilters().list(com.azure.core.util.Context.NONE);
    }
}
