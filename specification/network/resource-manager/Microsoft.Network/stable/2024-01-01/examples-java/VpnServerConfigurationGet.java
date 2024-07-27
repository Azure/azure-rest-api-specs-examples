
/**
 * Samples for VpnServerConfigurations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/VpnServerConfigurationGet.
     * json
     */
    /**
     * Sample code: VpnServerConfigurationGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnServerConfigurationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnServerConfigurations().getByResourceGroupWithResponse("rg1",
            "vpnServerConfiguration1", com.azure.core.util.Context.NONE);
    }
}
