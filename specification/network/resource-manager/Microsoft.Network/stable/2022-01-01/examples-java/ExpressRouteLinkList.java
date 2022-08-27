import com.azure.core.util.Context;

/** Samples for ExpressRouteLinks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ExpressRouteLinkList.json
     */
    /**
     * Sample code: ExpressRouteLinkGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteLinkGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteLinks().list("rg1", "portName", Context.NONE);
    }
}
