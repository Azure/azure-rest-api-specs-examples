import com.azure.core.util.Context;

/** Samples for ExpressRouteProviderPortsLocation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/expressRouteProviderPortList.json
     */
    /**
     * Sample code: ExpressRouteProviderPortList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRouteProviderPortList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteProviderPortsLocations()
            .listWithResponse(null, Context.NONE);
    }
}
