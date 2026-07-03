
/**
 * Samples for RouteFilterRules ListByRouteFilter.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteFilterRuleListByRouteFilter.json
     */
    /**
     * Sample code: RouteFilterRuleListByRouteFilter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeFilterRuleListByRouteFilter(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteFilterRules().listByRouteFilter("rg1", "filterName",
            com.azure.core.util.Context.NONE);
    }
}
