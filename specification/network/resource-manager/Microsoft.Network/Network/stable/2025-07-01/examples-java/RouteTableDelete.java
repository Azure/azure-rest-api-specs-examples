
/**
 * Samples for RouteTables Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableDelete.json
     */
    /**
     * Sample code: Delete route table.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteRouteTable(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteTables().delete("rg1", "testrt", com.azure.core.util.Context.NONE);
    }
}
