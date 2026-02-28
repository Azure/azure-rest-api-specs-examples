
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCircuitAuthorizationInner;

/**
 * Samples for ExpressRouteCircuitAuthorizations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ExpressRouteCircuitAuthorizationCreate.json
     */
    /**
     * Sample code: Create ExpressRouteCircuit Authorization.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createExpressRouteCircuitAuthorization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuitAuthorizations().createOrUpdate("rg1",
            "circuitName", "authorizatinName", new ExpressRouteCircuitAuthorizationInner(),
            com.azure.core.util.Context.NONE);
    }
}
