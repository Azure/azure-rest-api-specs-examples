
/**
 * Samples for ExpressRouteCrossConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/
     * ExpressRouteCrossConnectionList.json
     */
    /**
     * Sample code: ExpressRouteCrossConnectionList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteCrossConnectionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCrossConnections()
            .list(com.azure.core.util.Context.NONE);
    }
}
