
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkPeeringInner;

/**
 * Samples for VirtualNetworkPeerings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkPeeringCreateWithRemoteVirtualNetworkEncryption.json
     */
    /**
     * Sample code: Create peering with remote virtual network encryption.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createPeeringWithRemoteVirtualNetworkEncryption(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkPeerings().createOrUpdate("peerTest", "vnet1",
            "peer",
            new VirtualNetworkPeeringInner().withAllowVirtualNetworkAccess(true).withAllowForwardedTraffic(true)
                .withAllowGatewayTransit(false).withUseRemoteGateways(false)
                .withRemoteVirtualNetwork(new SubResource().withId(
                    "/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2")),
            null, com.azure.core.util.Context.NONE);
    }
}
