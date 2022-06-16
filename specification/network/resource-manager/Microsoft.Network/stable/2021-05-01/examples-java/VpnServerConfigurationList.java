import com.azure.core.util.Context;

/** Samples for VpnServerConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnServerConfigurationList.json
     */
    /**
     * Sample code: VpnServerConfigurationList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnServerConfigurationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnServerConfigurations().list(Context.NONE);
    }
}
