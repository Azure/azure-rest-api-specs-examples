
/**
 * Samples for RouteTables GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/RouteTableGet.json
     */
    /**
     * Sample code: Get route table.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getRouteTable(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getRouteTables().getByResourceGroupWithResponse("rg1", "testrt", null,
            com.azure.core.util.Context.NONE);
    }
}
