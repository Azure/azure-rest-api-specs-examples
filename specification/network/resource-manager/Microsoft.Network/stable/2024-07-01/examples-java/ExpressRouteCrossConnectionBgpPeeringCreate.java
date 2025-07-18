
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCrossConnectionPeeringInner;
import com.azure.resourcemanager.network.models.Ipv6ExpressRouteCircuitPeeringConfig;

/**
 * Samples for ExpressRouteCrossConnectionPeerings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * ExpressRouteCrossConnectionBgpPeeringCreate.json
     */
    /**
     * Sample code: ExpressRouteCrossConnectionBgpPeeringCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        expressRouteCrossConnectionBgpPeeringCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCrossConnectionPeerings()
            .createOrUpdate("CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering",
                new ExpressRouteCrossConnectionPeeringInner().withPeerAsn(200L)
                    .withPrimaryPeerAddressPrefix("192.168.16.252/30")
                    .withSecondaryPeerAddressPrefix("192.168.18.252/30").withVlanId(200)
                    .withIpv6PeeringConfig(new Ipv6ExpressRouteCircuitPeeringConfig()
                        .withPrimaryPeerAddressPrefix("3FFE:FFFF:0:CD30::/126")
                        .withSecondaryPeerAddressPrefix("3FFE:FFFF:0:CD30::4/126")),
                com.azure.core.util.Context.NONE);
    }
}
