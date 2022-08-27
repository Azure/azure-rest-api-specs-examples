import com.azure.core.util.Context;

/** Samples for RouteFilters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/RouteFilterGet.json
     */
    /**
     * Sample code: RouteFilterGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeFilterGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getRouteFilters()
            .getByResourceGroupWithResponse("rg1", "filterName", null, Context.NONE);
    }
}
