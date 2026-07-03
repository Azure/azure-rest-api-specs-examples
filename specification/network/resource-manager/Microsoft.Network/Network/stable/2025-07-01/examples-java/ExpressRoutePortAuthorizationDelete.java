
/**
 * Samples for ExpressRoutePortAuthorizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRoutePortAuthorizationDelete.json
     */
    /**
     * Sample code: Delete ExpressRoutePort Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteExpressRoutePortAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePortAuthorizations().delete("rg1", "expressRoutePortName",
            "authorizationName", com.azure.core.util.Context.NONE);
    }
}
