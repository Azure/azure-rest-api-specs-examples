import com.azure.core.util.Context;

/** Samples for ExpressRouteCircuits GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteCircuitGet.json
     */
    /**
     * Sample code: Get ExpressRouteCircuit.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRouteCircuit(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCircuits()
            .getByResourceGroupWithResponse("rg1", "circuitName", Context.NONE);
    }
}
