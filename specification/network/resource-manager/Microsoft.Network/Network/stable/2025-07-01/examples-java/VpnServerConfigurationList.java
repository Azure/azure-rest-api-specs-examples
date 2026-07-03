
/**
 * Samples for VpnServerConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnServerConfigurationList.json
     */
    /**
     * Sample code: VpnServerConfigurationList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnServerConfigurationList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnServerConfigurations().list(com.azure.core.util.Context.NONE);
    }
}
