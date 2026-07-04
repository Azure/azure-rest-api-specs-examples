
/**
 * Samples for ExpressRoutePorts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRoutePortList.json
     */
    /**
     * Sample code: ExpressRoutePortList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRoutePortList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePorts().list(com.azure.core.util.Context.NONE);
    }
}
