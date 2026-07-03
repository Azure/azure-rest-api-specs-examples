
/**
 * Samples for HubRouteTables Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/HubRouteTableDelete.json
     */
    /**
     * Sample code: RouteTableDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void routeTableDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getHubRouteTables().delete("rg1", "virtualHub1", "hubRouteTable1",
            com.azure.core.util.Context.NONE);
    }
}
