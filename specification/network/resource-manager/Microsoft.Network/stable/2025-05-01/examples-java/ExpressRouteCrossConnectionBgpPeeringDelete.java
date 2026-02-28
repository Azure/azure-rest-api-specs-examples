
/**
 * Samples for ExpressRouteCrossConnectionPeerings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ExpressRouteCrossConnectionBgpPeeringDelete.json
     */
    /**
     * Sample code: DeleteExpressRouteCrossConnectionBgpPeering.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteExpressRouteCrossConnectionBgpPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCrossConnectionPeerings().delete(
            "CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering",
            com.azure.core.util.Context.NONE);
    }
}
