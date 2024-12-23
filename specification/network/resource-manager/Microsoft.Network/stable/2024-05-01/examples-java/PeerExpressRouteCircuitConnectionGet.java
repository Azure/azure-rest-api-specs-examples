
/**
 * Samples for PeerExpressRouteCircuitConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/
     * PeerExpressRouteCircuitConnectionGet.json
     */
    /**
     * Sample code: PeerExpressRouteCircuitConnectionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void peerExpressRouteCircuitConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPeerExpressRouteCircuitConnections().getWithResponse("rg1",
            "ExpressRouteARMCircuitA", "AzurePrivatePeering", "60aee347-e889-4a42-8c1b-0aae8b1e4013",
            com.azure.core.util.Context.NONE);
    }
}
