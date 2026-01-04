
/**
 * Samples for RouteFilterRules ListByRouteFilter.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * RouteFilterRuleListByRouteFilter.json
     */
    /**
     * Sample code: RouteFilterRuleListByRouteFilter.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeFilterRuleListByRouteFilter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteFilterRules().listByRouteFilter("rg1", "filterName",
            com.azure.core.util.Context.NONE);
    }
}
