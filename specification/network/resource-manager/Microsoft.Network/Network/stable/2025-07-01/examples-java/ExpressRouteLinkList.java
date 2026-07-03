
/**
 * Samples for ExpressRouteLinks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteLinkList.json
     */
    /**
     * Sample code: ExpressRouteLinkGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteLinkGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLinks().list("rg1", "portName", com.azure.core.util.Context.NONE);
    }
}
