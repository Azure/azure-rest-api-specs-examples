
/**
 * Samples for ExpressRouteProviderPortsLocation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/expressRouteProviderPortList.json
     */
    /**
     * Sample code: ExpressRouteProviderPortList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteProviderPortList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteProviderPortsLocations().listWithResponse(null,
            com.azure.core.util.Context.NONE);
    }
}
