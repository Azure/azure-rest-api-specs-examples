
/**
 * Samples for ExpressRouteCrossConnectionPeerings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCrossConnectionBgpPeeringGet.json
     */
    /**
     * Sample code: GetExpressRouteCrossConnectionBgpPeering.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getExpressRouteCrossConnectionBgpPeering(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnectionPeerings().getWithResponse(
            "CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering",
            com.azure.core.util.Context.NONE);
    }
}
