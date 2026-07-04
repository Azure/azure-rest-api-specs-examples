
/**
 * Samples for ExpressRouteCircuitConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitConnectionGet.json
     */
    /**
     * Sample code: ExpressRouteCircuitConnectionGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteCircuitConnectionGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitConnections().getWithResponse("rg1", "ExpressRouteARMCircuitA",
            "AzurePrivatePeering", "circuitConnectionUSAUS", com.azure.core.util.Context.NONE);
    }
}
