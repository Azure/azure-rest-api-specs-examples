import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.VpnConnectionInner;
import com.azure.resourcemanager.network.fluent.models.VpnSiteLinkConnectionInner;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewayConnectionProtocol;
import com.azure.resourcemanager.network.models.VpnLinkConnectionMode;
import java.util.Arrays;

/** Samples for VpnConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VpnConnectionPut.json
     */
    /**
     * Sample code: VpnConnectionPut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnConnectionPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnConnections()
            .createOrUpdate(
                "rg1",
                "gateway1",
                "vpnConnection1",
                new VpnConnectionInner()
                    .withRemoteVpnSite(
                        new SubResource()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"))
                    .withTrafficSelectorPolicies(Arrays.asList())
                    .withVpnLinkConnections(
                        Arrays
                            .asList(
                                new VpnSiteLinkConnectionInner()
                                    .withName("Connection-Link1")
                                    .withVpnSiteLink(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink1"))
                                    .withVpnLinkConnectionMode(VpnLinkConnectionMode.DEFAULT)
                                    .withVpnConnectionProtocolType(VirtualNetworkGatewayConnectionProtocol.IKEV2)
                                    .withConnectionBandwidth(200)
                                    .withSharedKey("key")
                                    .withUsePolicyBasedTrafficSelectors(false))),
                Context.NONE);
    }
}
