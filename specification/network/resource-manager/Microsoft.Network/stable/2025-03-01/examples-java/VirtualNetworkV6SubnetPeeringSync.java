
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkPeeringInner;
import com.azure.resourcemanager.network.models.SyncRemoteAddressSpace;

/**
 * Samples for VirtualNetworkPeerings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkV6SubnetPeeringSync.json
     */
    /**
     * Sample code: Sync V6 Subnet Peering.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void syncV6SubnetPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkPeerings().createOrUpdate("peerTest", "vnet1",
            "peer",
            new VirtualNetworkPeeringInner().withAllowVirtualNetworkAccess(true).withAllowForwardedTraffic(true)
                .withAllowGatewayTransit(false).withUseRemoteGateways(false)
                .withRemoteVirtualNetwork(new SubResource().withId(
                    "/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"))
                .withPeerCompleteVnets(false).withEnableOnlyIPv6Peering(true),
            SyncRemoteAddressSpace.TRUE, com.azure.core.util.Context.NONE);
    }
}
