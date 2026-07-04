
/**
 * Samples for RouteFilterRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteFilterRuleDelete.json
     */
    /**
     * Sample code: RouteFilterRuleDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeFilterRuleDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteFilterRules().delete("rg1", "filterName", "ruleName",
            com.azure.core.util.Context.NONE);
    }
}
