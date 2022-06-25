import com.azure.core.util.Context;

/** Samples for ExpressRouteCircuitPeerings Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCircuitPeeringDelete.json
     */
    /**
     * Sample code: Delete ExpressRouteCircuit Peerings.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteExpressRouteCircuitPeerings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCircuitPeerings()
            .delete("rg1", "circuitName", "peeringName", Context.NONE);
    }
}
