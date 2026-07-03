
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkPeeringInner;
import com.azure.resourcemanager.network.models.SyncRemoteAddressSpace;

/**
 * Samples for VirtualNetworkPeerings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkV6SubnetPeeringSync.json
     */
    /**
     * Sample code: Sync V6 Subnet Peering.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void syncV6SubnetPeering(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkPeerings().createOrUpdate("peerTest", "vnet1", "peer",
            new VirtualNetworkPeeringInner().withAllowVirtualNetworkAccess(true).withAllowForwardedTraffic(true)
                .withAllowGatewayTransit(false).withUseRemoteGateways(false)
                .withRemoteVirtualNetwork(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"))
                .withPeerCompleteVnets(false).withEnableOnlyIPv6Peering(true),
            SyncRemoteAddressSpace.TRUE, com.azure.core.util.Context.NONE);
    }
}
