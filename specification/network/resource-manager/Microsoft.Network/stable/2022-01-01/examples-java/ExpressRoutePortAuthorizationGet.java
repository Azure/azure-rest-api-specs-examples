import com.azure.core.util.Context;

/** Samples for ExpressRoutePortAuthorizations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ExpressRoutePortAuthorizationGet.json
     */
    /**
     * Sample code: Get ExpressRoutePort Authorization.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRoutePortAuthorization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRoutePortAuthorizations()
            .getWithResponse("rg1", "expressRoutePortName", "authorizationName", Context.NONE);
    }
}
