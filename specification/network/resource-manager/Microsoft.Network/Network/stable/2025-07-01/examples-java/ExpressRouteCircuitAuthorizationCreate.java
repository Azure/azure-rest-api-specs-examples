
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCircuitAuthorizationInner;

/**
 * Samples for ExpressRouteCircuitAuthorizations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitAuthorizationCreate.json
     */
    /**
     * Sample code: Create ExpressRouteCircuit Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createExpressRouteCircuitAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitAuthorizations().createOrUpdate("rg1", "circuitName",
            "authorizatinName", new ExpressRouteCircuitAuthorizationInner(), com.azure.core.util.Context.NONE);
    }
}
