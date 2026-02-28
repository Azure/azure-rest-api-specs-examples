
/**
 * Samples for ExpressRouteLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ExpressRouteLinkGet.json
     */
    /**
     * Sample code: ExpressRouteLinkGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteLinkGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteLinks().getWithResponse("rg1", "portName", "linkName",
            com.azure.core.util.Context.NONE);
    }
}
