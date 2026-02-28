
/**
 * Samples for ExpressRoutePortAuthorizations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ExpressRoutePortAuthorizationGet.json
     */
    /**
     * Sample code: Get ExpressRoutePort Authorization.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRoutePortAuthorization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRoutePortAuthorizations().getWithResponse("rg1",
            "expressRoutePortName", "authorizationName", com.azure.core.util.Context.NONE);
    }
}
