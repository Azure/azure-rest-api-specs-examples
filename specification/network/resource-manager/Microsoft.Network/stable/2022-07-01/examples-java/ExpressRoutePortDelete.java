import com.azure.core.util.Context;

/** Samples for ExpressRoutePorts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ExpressRoutePortDelete.json
     */
    /**
     * Sample code: ExpressRoutePortDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRoutePortDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRoutePorts().delete("rg1", "portName", Context.NONE);
    }
}
