
/**
 * Samples for ExpressRouteCircuitConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/
     * ExpressRouteCircuitConnectionGet.json
     */
    /**
     * Sample code: ExpressRouteCircuitConnectionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteCircuitConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuitConnections().getWithResponse("rg1",
            "ExpressRouteARMCircuitA", "AzurePrivatePeering", "circuitConnectionUSAUS",
            com.azure.core.util.Context.NONE);
    }
}
