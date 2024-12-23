
/**
 * Samples for RouteTables ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/RouteTableList.json
     */
    /**
     * Sample code: List route tables in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRouteTablesInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getRouteTables().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
