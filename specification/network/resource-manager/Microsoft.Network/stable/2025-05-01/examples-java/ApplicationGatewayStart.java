
/**
 * Samples for ApplicationGateways Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ApplicationGatewayStart.json
     */
    /**
     * Sample code: Start Application Gateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startApplicationGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGateways().start("rg1", "appgw",
            com.azure.core.util.Context.NONE);
    }
}
