
/**
 * Samples for RouteFilters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteFilterDelete.json
     */
    /**
     * Sample code: RouteFilterDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeFilterDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteFilters().delete("rg1", "filterName", com.azure.core.util.Context.NONE);
    }
}
