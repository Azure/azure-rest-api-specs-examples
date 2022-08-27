import com.azure.core.util.Context;

/** Samples for ApplicationGateways ListAvailableResponseHeaders. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ApplicationGatewayAvailableResponseHeadersGet.json
     */
    /**
     * Sample code: Get Available Response Headers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableResponseHeaders(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getApplicationGateways()
            .listAvailableResponseHeadersWithResponse(Context.NONE);
    }
}
