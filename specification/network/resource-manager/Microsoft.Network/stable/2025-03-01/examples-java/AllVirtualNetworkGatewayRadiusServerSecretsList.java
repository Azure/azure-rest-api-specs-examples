
/**
 * Samples for VirtualNetworkGateways ListRadiusSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * AllVirtualNetworkGatewayRadiusServerSecretsList.json
     */
    /**
     * Sample code: ListAllVirtualNetworkGatewayRadiusServerSecrets.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllVirtualNetworkGatewayRadiusServerSecrets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().listRadiusSecretsWithResponse("rg1",
            "vpngw", com.azure.core.util.Context.NONE);
    }
}
