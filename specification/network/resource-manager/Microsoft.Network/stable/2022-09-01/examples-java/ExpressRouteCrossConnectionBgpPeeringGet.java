/** Samples for ExpressRouteCrossConnectionPeerings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/ExpressRouteCrossConnectionBgpPeeringGet.json
     */
    /**
     * Sample code: GetExpressRouteCrossConnectionBgpPeering.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRouteCrossConnectionBgpPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCrossConnectionPeerings()
            .getWithResponse(
                "CrossConnection-SiliconValley",
                "<circuitServiceKey>",
                "AzurePrivatePeering",
                com.azure.core.util.Context.NONE);
    }
}
