
/**
 * Samples for RouteFilters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/RouteFilterList.json
     */
    /**
     * Sample code: RouteFilterList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeFilterList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteFilters().list(com.azure.core.util.Context.NONE);
    }
}
