
/**
 * Samples for ExpressRouteCrossConnectionPeerings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCrossConnectionBgpPeeringDelete.json
     */
    /**
     * Sample code: DeleteExpressRouteCrossConnectionBgpPeering.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteExpressRouteCrossConnectionBgpPeering(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnectionPeerings().delete("CrossConnection-SiliconValley",
            "<circuitServiceKey>", "AzurePrivatePeering", com.azure.core.util.Context.NONE);
    }
}
