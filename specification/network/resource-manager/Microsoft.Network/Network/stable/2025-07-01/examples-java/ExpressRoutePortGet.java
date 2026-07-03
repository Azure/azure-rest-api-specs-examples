
/**
 * Samples for ExpressRoutePorts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRoutePortGet.json
     */
    /**
     * Sample code: ExpressRoutePortGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRoutePortGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePorts().getByResourceGroupWithResponse("rg1", "portName",
            com.azure.core.util.Context.NONE);
    }
}
