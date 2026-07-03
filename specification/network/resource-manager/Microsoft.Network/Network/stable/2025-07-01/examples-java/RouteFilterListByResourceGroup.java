
/**
 * Samples for RouteFilters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteFilterListByResourceGroup.json
     */
    /**
     * Sample code: RouteFilterListByResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeFilterListByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteFilters().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
