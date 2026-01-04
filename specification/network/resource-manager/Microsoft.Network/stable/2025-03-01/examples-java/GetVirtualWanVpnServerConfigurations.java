
/**
 * Samples for VpnServerConfigurationsAssociatedWithVirtualWan List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * GetVirtualWanVpnServerConfigurations.json
     */
    /**
     * Sample code: GetVirtualWanVpnServerConfigurations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualWanVpnServerConfigurations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnServerConfigurationsAssociatedWithVirtualWans().list("rg1",
            "wan1", com.azure.core.util.Context.NONE);
    }
}
