
import com.azure.resourcemanager.network.fluent.models.ExpressRoutePortAuthorizationInner;

/**
 * Samples for ExpressRoutePortAuthorizations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ExpressRoutePortAuthorizationCreate.json
     */
    /**
     * Sample code: Create ExpressRoutePort Authorization.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createExpressRoutePortAuthorization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRoutePortAuthorizations().createOrUpdate("rg1",
            "expressRoutePortName", "authorizatinName", new ExpressRoutePortAuthorizationInner(),
            com.azure.core.util.Context.NONE);
    }
}
