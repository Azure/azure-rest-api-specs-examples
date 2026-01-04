
/**
 * Samples for ExpressRouteCircuitPeerings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteCircuitPeeringGet
     * .json
     */
    /**
     * Sample code: Get ExpressRouteCircuit Peering.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRouteCircuitPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuitPeerings().getWithResponse("rg1",
            "circuitName", "MicrosoftPeering", com.azure.core.util.Context.NONE);
    }
}
