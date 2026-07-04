
import com.azure.resourcemanager.network.fluent.models.ExpressRouteCircuitPeeringInner;

/**
 * Samples for ExpressRouteCircuitPeerings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitPeeringCreate.json
     */
    /**
     * Sample code: Create ExpressRouteCircuit Peerings.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createExpressRouteCircuitPeerings(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitPeerings().createOrUpdate("rg1", "circuitName",
            "AzurePrivatePeering",
            new ExpressRouteCircuitPeeringInner().withPeerAsn(200L).withPrimaryPeerAddressPrefix("192.168.16.252/30")
                .withSecondaryPeerAddressPrefix("192.168.18.252/30").withVlanId(200),
            com.azure.core.util.Context.NONE);
    }
}
