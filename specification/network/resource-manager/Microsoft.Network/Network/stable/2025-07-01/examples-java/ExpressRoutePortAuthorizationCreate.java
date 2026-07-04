
import com.azure.resourcemanager.network.fluent.models.ExpressRoutePortAuthorizationInner;

/**
 * Samples for ExpressRoutePortAuthorizations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRoutePortAuthorizationCreate.json
     */
    /**
     * Sample code: Create ExpressRoutePort Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createExpressRoutePortAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRoutePortAuthorizations().createOrUpdate("rg1", "expressRoutePortName",
            "authorizatinName", new ExpressRoutePortAuthorizationInner(), com.azure.core.util.Context.NONE);
    }
}
