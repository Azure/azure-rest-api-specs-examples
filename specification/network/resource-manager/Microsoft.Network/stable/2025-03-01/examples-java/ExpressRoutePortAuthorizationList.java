
/**
 * Samples for ExpressRoutePortAuthorizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ExpressRoutePortAuthorizationList.json
     */
    /**
     * Sample code: List ExpressRoutePort Authorization.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExpressRoutePortAuthorization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRoutePortAuthorizations().list("rg1",
            "expressRoutePortName", com.azure.core.util.Context.NONE);
    }
}
