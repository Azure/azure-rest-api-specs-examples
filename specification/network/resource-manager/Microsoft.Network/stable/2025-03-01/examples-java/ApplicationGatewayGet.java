
/**
 * Samples for ApplicationGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ApplicationGatewayGet.json
     */
    /**
     * Sample code: Get ApplicationGateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getApplicationGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGateways().getByResourceGroupWithResponse("rg1",
            "appgw", com.azure.core.util.Context.NONE);
    }
}
