
/**
 * Samples for ExpressRouteLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteLinkGet.json
     */
    /**
     * Sample code: ExpressRouteLinkGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteLinkGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLinks().getWithResponse("rg1", "portName", "linkName",
            com.azure.core.util.Context.NONE);
    }
}
