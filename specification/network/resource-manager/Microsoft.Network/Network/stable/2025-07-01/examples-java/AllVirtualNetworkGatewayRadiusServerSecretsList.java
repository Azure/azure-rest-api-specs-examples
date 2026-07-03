
/**
 * Samples for VirtualNetworkGateways ListRadiusSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AllVirtualNetworkGatewayRadiusServerSecretsList.json
     */
    /**
     * Sample code: ListAllVirtualNetworkGatewayRadiusServerSecrets.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllVirtualNetworkGatewayRadiusServerSecrets(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().listRadiusSecretsWithResponse("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
