
/**
 * Samples for ExpressRoutePorts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRoutePortListByResourceGroup.json
     */
    /**
     * Sample code: ExpressRoutePortListByResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRoutePortListByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePorts().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
