
/**
 * Samples for ExpressRouteCircuitPeerings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitPeeringDelete.json
     */
    /**
     * Sample code: Delete ExpressRouteCircuit Peerings.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteExpressRouteCircuitPeerings(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitPeerings().delete("rg1", "circuitName", "peeringName",
            com.azure.core.util.Context.NONE);
    }
}
