
/**
 * Samples for ExpressRouteCircuitPeerings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitPeeringList.json
     */
    /**
     * Sample code: List ExpressRouteCircuit Peerings.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRouteCircuitPeerings(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitPeerings().list("rg1", "circuitName",
            com.azure.core.util.Context.NONE);
    }
}
