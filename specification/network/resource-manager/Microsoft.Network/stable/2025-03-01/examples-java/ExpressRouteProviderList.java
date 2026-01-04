
/**
 * Samples for ExpressRouteServiceProviders List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteProviderList.json
     */
    /**
     * Sample code: List ExpressRoute providers.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExpressRouteProviders(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteServiceProviders()
            .list(com.azure.core.util.Context.NONE);
    }
}
