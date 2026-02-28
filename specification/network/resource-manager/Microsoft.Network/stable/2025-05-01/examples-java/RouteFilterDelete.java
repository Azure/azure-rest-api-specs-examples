
/**
 * Samples for RouteFilters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/RouteFilterDelete.json
     */
    /**
     * Sample code: RouteFilterDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeFilterDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteFilters().delete("rg1", "filterName",
            com.azure.core.util.Context.NONE);
    }
}
