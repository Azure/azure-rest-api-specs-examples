
/**
 * Samples for RouteFilterRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteFilterRuleGet.json
     */
    /**
     * Sample code: RouteFilterRuleGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeFilterRuleGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteFilterRules().getWithResponse("rg1", "filterName", "filterName",
            com.azure.core.util.Context.NONE);
    }
}
