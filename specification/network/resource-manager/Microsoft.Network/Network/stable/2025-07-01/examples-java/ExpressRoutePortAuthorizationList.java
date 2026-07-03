
/**
 * Samples for ExpressRoutePortAuthorizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRoutePortAuthorizationList.json
     */
    /**
     * Sample code: List ExpressRoutePort Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRoutePortAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePortAuthorizations().list("rg1", "expressRoutePortName",
            com.azure.core.util.Context.NONE);
    }
}
