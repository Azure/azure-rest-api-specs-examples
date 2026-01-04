
/**
 * Samples for ExpressRouteCrossConnections ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ExpressRouteCrossConnectionListByResourceGroup.json
     */
    /**
     * Sample code: ExpressRouteCrossConnectionListByResourceGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        expressRouteCrossConnectionListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCrossConnections()
            .listByResourceGroup("CrossConnection-SiliconValley", com.azure.core.util.Context.NONE);
    }
}
