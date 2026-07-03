
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkPeeringInner;
import java.util.Arrays;

/**
 * Samples for VirtualNetworkPeerings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkSubnetPeeringCreate.json
     */
    /**
     * Sample code: Create subnet peering.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createSubnetPeering(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkPeerings().createOrUpdate("peerTest", "vnet1", "peer",
            new VirtualNetworkPeeringInner().withAllowVirtualNetworkAccess(true).withAllowForwardedTraffic(true)
                .withAllowGatewayTransit(false).withUseRemoteGateways(false)
                .withRemoteVirtualNetwork(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"))
                .withPeerCompleteVnets(false).withEnableOnlyIPv6Peering(false)
                .withLocalSubnetNames(Arrays.asList("Subnet1", "Subnet4"))
                .withRemoteSubnetNames(Arrays.asList("Subnet2")),
            null, com.azure.core.util.Context.NONE);
    }
}
