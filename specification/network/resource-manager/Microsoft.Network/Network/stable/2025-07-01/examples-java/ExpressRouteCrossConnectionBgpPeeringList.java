
/**
 * Samples for ExpressRouteCrossConnectionPeerings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCrossConnectionBgpPeeringList.json
     */
    /**
     * Sample code: ExpressRouteCrossConnectionBgpPeeringList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        expressRouteCrossConnectionBgpPeeringList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnectionPeerings().list("CrossConnection-SiliconValley",
            "<circuitServiceKey>", com.azure.core.util.Context.NONE);
    }
}
