
/**
 * Samples for HubRouteTables List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/HubRouteTableList.json
     */
    /**
     * Sample code: RouteTableList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeTableList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getHubRouteTables().list("rg1", "virtualHub1",
            com.azure.core.util.Context.NONE);
    }
}
