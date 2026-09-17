
/**
 * Samples for ExpressRouteCircuitPeerings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ExpressRouteCircuitPeeringGet.json
     */
    /**
     * Sample code: Get ExpressRouteCircuit Peering.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getExpressRouteCircuitPeering(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitPeerings().getWithResponse("rg1", "circuitName",
            "MicrosoftPeering", com.azure.core.util.Context.NONE);
    }
}
