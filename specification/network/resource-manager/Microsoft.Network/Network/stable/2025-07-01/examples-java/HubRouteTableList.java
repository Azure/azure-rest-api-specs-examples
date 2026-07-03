
/**
 * Samples for HubRouteTables List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/HubRouteTableList.json
     */
    /**
     * Sample code: RouteTableList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeTableList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getHubRouteTables().list("rg1", "virtualHub1", com.azure.core.util.Context.NONE);
    }
}
