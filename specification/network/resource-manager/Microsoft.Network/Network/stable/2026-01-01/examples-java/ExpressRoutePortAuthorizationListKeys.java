
/**
 * Samples for ExpressRoutePortAuthorizations ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ExpressRoutePortAuthorizationListKeys.json
     */
    /**
     * Sample code: List ExpressRoutePort Authorization Keys.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRoutePortAuthorizationKeys(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePortAuthorizations().listKeysWithResponse("rg1", "expressRoutePortName",
            "authorizationName", com.azure.core.util.Context.NONE);
    }
}
