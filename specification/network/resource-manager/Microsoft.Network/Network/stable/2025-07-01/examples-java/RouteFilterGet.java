
/**
 * Samples for RouteFilters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteFilterGet.json
     */
    /**
     * Sample code: RouteFilterGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeFilterGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteFilters().getByResourceGroupWithResponse("rg1", "filterName", null,
            com.azure.core.util.Context.NONE);
    }
}
