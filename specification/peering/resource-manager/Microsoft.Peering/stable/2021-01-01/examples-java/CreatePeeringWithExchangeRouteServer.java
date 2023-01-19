import com.azure.core.management.SubResource;
import com.azure.resourcemanager.peering.models.BgpSession;
import com.azure.resourcemanager.peering.models.DirectConnection;
import com.azure.resourcemanager.peering.models.DirectPeeringType;
import com.azure.resourcemanager.peering.models.Kind;
import com.azure.resourcemanager.peering.models.PeeringPropertiesDirect;
import com.azure.resourcemanager.peering.models.PeeringSku;
import com.azure.resourcemanager.peering.models.SessionAddressProvider;
import java.util.Arrays;

/** Samples for Peerings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/CreatePeeringWithExchangeRouteServer.json
     */
    /**
     * Sample code: Create a peering with exchange route server.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void createAPeeringWithExchangeRouteServer(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager
            .peerings()
            .define("peeringName")
            .withRegion("eastus")
            .withExistingResourceGroup("rgName")
            .withSku(new PeeringSku().withName("Premium_Direct_Free"))
            .withKind(Kind.DIRECT)
            .withDirect(
                new PeeringPropertiesDirect()
                    .withConnections(
                        Arrays
                            .asList(
                                new DirectConnection()
                                    .withBandwidthInMbps(10000)
                                    .withSessionAddressProvider(SessionAddressProvider.PEER)
                                    .withUseForPeeringService(true)
                                    .withPeeringDBFacilityId(99999)
                                    .withBgpSession(
                                        new BgpSession()
                                            .withSessionPrefixV4("192.168.0.0/24")
                                            .withMicrosoftSessionIPv4Address("192.168.0.123")
                                            .withPeerSessionIPv4Address("192.168.0.234")
                                            .withMaxPrefixesAdvertisedV4(1000)
                                            .withMaxPrefixesAdvertisedV6(100))
                                    .withConnectionIdentifier("5F4CB5C7-6B43-4444-9338-9ABC72606C16")))
                    .withPeerAsn(
                        new SubResource().withId("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"))
                    .withDirectPeeringType(DirectPeeringType.IX_RS))
            .withPeeringLocation("peeringLocation0")
            .create();
    }
}
