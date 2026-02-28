
import com.azure.resourcemanager.network.fluent.models.ConnectionSharedKeyResultInner;
import com.azure.resourcemanager.network.models.SharedKeyProperties;

/**
 * Samples for VpnLinkConnections SetOrInitDefaultSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VpnSiteLinkConnectionDefaultSharedKeyPut.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionDefaultSharedKeyPut.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteLinkConnectionDefaultSharedKeyPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnLinkConnections().setOrInitDefaultSharedKey(
            "rg1", "gateway1", "vpnConnection1", "Connection-Link1", new ConnectionSharedKeyResultInner()
                .withProperties(new SharedKeyProperties().withSharedKey("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
    }
}
