import com.azure.core.util.Context;

/** Samples for RouteTables GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/RouteTableGet.json
     */
    /**
     * Sample code: Get route table.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRouteTable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getRouteTables()
            .getByResourceGroupWithResponse("rg1", "testrt", null, Context.NONE);
    }
}
