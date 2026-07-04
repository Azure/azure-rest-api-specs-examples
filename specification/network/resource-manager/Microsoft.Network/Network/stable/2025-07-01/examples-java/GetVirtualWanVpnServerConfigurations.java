
/**
 * Samples for VpnServerConfigurationsAssociatedWithVirtualWan List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/GetVirtualWanVpnServerConfigurations.json
     */
    /**
     * Sample code: GetVirtualWanVpnServerConfigurations.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualWanVpnServerConfigurations(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnServerConfigurationsAssociatedWithVirtualWans().list("rg1", "wan1",
            com.azure.core.util.Context.NONE);
    }
}
