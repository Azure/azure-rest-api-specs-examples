
/**
 * Samples for ApplicationGateways Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ApplicationGatewayStop.json
     */
    /**
     * Sample code: Stop Application Gateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void stopApplicationGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGateways().stop("rg1", "appgw",
            com.azure.core.util.Context.NONE);
    }
}
