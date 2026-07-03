
/**
 * Samples for ExpressRoutePortsLocations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRoutePortsLocationList.json
     */
    /**
     * Sample code: ExpressRoutePortsLocationList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRoutePortsLocationList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePortsLocations().list(com.azure.core.util.Context.NONE);
    }
}
