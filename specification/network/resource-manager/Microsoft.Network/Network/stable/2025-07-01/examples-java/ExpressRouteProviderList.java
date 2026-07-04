
/**
 * Samples for ExpressRouteServiceProviders List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteProviderList.json
     */
    /**
     * Sample code: List ExpressRoute providers.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRouteProviders(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteServiceProviders().list(com.azure.core.util.Context.NONE);
    }
}
