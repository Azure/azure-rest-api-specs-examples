
/**
 * Samples for ExpressRouteCrossConnections GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ExpressRouteCrossConnectionGet.json
     */
    /**
     * Sample code: GetExpressRouteCrossConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRouteCrossConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCrossConnections().getByResourceGroupWithResponse(
            "CrossConnection-SiliconValley", "<circuitServiceKey>", com.azure.core.util.Context.NONE);
    }
}
