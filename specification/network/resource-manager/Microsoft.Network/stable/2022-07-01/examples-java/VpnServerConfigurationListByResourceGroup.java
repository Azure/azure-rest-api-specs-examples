import com.azure.core.util.Context;

/** Samples for VpnServerConfigurations ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VpnServerConfigurationListByResourceGroup.json
     */
    /**
     * Sample code: VpnServerConfigurationListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnServerConfigurationListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnServerConfigurations()
            .listByResourceGroup("rg1", Context.NONE);
    }
}
