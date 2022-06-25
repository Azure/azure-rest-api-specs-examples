import com.azure.core.util.Context;

/** Samples for RouteFilterRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/RouteFilterRuleGet.json
     */
    /**
     * Sample code: RouteFilterRuleGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeFilterRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getRouteFilterRules()
            .getWithResponse("rg1", "filterName", "filterName", Context.NONE);
    }
}
