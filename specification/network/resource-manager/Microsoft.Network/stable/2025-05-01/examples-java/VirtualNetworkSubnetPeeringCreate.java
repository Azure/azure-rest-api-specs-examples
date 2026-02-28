
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkPeeringInner;
import java.util.Arrays;

/**
 * Samples for VirtualNetworkPeerings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkSubnetPeeringCreate.json
     */
    /**
     * Sample code: Create subnet peering.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createSubnetPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkPeerings().createOrUpdate("peerTest", "vnet1",
            "peer",
            new VirtualNetworkPeeringInner().withAllowVirtualNetworkAccess(true).withAllowForwardedTraffic(true)
                .withAllowGatewayTransit(false).withUseRemoteGateways(false)
                .withRemoteVirtualNetwork(new SubResource().withId(
                    "/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"))
                .withPeerCompleteVnets(false).withEnableOnlyIPv6Peering(false)
                .withLocalSubnetNames(Arrays.asList("Subnet1", "Subnet4"))
                .withRemoteSubnetNames(Arrays.asList("Subnet2")),
            null, com.azure.core.util.Context.NONE);
    }
}
