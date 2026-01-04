
/**
 * Samples for ExpressRoutePortAuthorizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ExpressRoutePortAuthorizationDelete.json
     */
    /**
     * Sample code: Delete ExpressRoutePort Authorization.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteExpressRoutePortAuthorization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRoutePortAuthorizations().delete("rg1",
            "expressRoutePortName", "authorizationName", com.azure.core.util.Context.NONE);
    }
}
