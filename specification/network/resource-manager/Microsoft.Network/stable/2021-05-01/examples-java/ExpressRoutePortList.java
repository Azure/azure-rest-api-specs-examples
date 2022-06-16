import com.azure.core.util.Context;

/** Samples for ExpressRoutePorts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRoutePortList.json
     */
    /**
     * Sample code: ExpressRoutePortList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRoutePortList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRoutePorts().list(Context.NONE);
    }
}
