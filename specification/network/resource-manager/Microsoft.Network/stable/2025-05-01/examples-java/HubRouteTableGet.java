
/**
 * Samples for HubRouteTables Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/HubRouteTableGet.json
     */
    /**
     * Sample code: RouteTableGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeTableGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getHubRouteTables().getWithResponse("rg1", "virtualHub1",
            "hubRouteTable1", com.azure.core.util.Context.NONE);
    }
}
