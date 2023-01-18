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
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/CreateDirectPeering.json
     */
    /**
     * Sample code: Create a direct peering.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void createADirectPeering(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager
            .peerings()
            .define("peeringName")
            .withRegion("eastus")
            .withExistingResourceGroup("rgName")
            .withSku(new PeeringSku().withName("Basic_Direct_Free"))
            .withKind(Kind.DIRECT)
            .withDirect(
                new PeeringPropertiesDirect()
                    .withConnections(
                        Arrays
                            .asList(
                                new DirectConnection()
                                    .withBandwidthInMbps(10000)
                                    .withSessionAddressProvider(SessionAddressProvider.PEER)
                                    .withUseForPeeringService(false)
                                    .withPeeringDBFacilityId(99999)
                                    .withBgpSession(
                                        new BgpSession()
                                            .withSessionPrefixV4("192.168.0.0/31")
                                            .withSessionPrefixV6("fd00::0/127")
                                            .withMaxPrefixesAdvertisedV4(1000)
                                            .withMaxPrefixesAdvertisedV6(100)
                                            .withMd5AuthenticationKey("fakeTokenPlaceholder"))
                                    .withConnectionIdentifier("5F4CB5C7-6B43-4444-9338-9ABC72606C16"),
                                new DirectConnection()
                                    .withBandwidthInMbps(10000)
                                    .withSessionAddressProvider(SessionAddressProvider.MICROSOFT)
                                    .withUseForPeeringService(true)
                                    .withPeeringDBFacilityId(99999)
                                    .withConnectionIdentifier("8AB00818-D533-4504-A25A-03A17F61201C")))
                    .withPeerAsn(
                        new SubResource().withId("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"))
                    .withDirectPeeringType(DirectPeeringType.EDGE))
            .withPeeringLocation("peeringLocation0")
            .create();
    }
}
