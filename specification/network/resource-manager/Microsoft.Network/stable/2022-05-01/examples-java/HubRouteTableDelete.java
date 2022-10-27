import com.azure.core.util.Context;

/** Samples for HubRouteTables Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/HubRouteTableDelete.json
     */
    /**
     * Sample code: RouteTableDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeTableDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getHubRouteTables()
            .delete("rg1", "virtualHub1", "hubRouteTable1", Context.NONE);
    }
}
