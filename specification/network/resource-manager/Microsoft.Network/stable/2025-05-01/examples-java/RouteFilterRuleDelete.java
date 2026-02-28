
/**
 * Samples for RouteFilterRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/RouteFilterRuleDelete.json
     */
    /**
     * Sample code: RouteFilterRuleDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeFilterRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteFilterRules().delete("rg1", "filterName", "ruleName",
            com.azure.core.util.Context.NONE);
    }
}
