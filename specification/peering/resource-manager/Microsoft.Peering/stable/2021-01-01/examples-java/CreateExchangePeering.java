import com.azure.core.management.SubResource;
import com.azure.resourcemanager.peering.models.BgpSession;
import com.azure.resourcemanager.peering.models.ExchangeConnection;
import com.azure.resourcemanager.peering.models.Kind;
import com.azure.resourcemanager.peering.models.PeeringPropertiesExchange;
import com.azure.resourcemanager.peering.models.PeeringSku;
import java.util.Arrays;

/** Samples for Peerings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/CreateExchangePeering.json
     */
    /**
     * Sample code: Create an exchange peering.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void createAnExchangePeering(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager
            .peerings()
            .define("peeringName")
            .withRegion("eastus")
            .withExistingResourceGroup("rgName")
            .withSku(new PeeringSku().withName("Basic_Exchange_Free"))
            .withKind(Kind.EXCHANGE)
            .withExchange(
                new PeeringPropertiesExchange()
                    .withConnections(
                        Arrays
                            .asList(
                                new ExchangeConnection()
                                    .withPeeringDBFacilityId(99999)
                                    .withBgpSession(
                                        new BgpSession()
                                            .withPeerSessionIPv4Address("192.168.2.1")
                                            .withPeerSessionIPv6Address("fd00::1")
                                            .withMaxPrefixesAdvertisedV4(1000)
                                            .withMaxPrefixesAdvertisedV6(100)
                                            .withMd5AuthenticationKey("fakeTokenPlaceholder"))
                                    .withConnectionIdentifier("CE495334-0E94-4E51-8164-8116D6CD284D"),
                                new ExchangeConnection()
                                    .withPeeringDBFacilityId(99999)
                                    .withBgpSession(
                                        new BgpSession()
                                            .withPeerSessionIPv4Address("192.168.2.2")
                                            .withPeerSessionIPv6Address("fd00::2")
                                            .withMaxPrefixesAdvertisedV4(1000)
                                            .withMaxPrefixesAdvertisedV6(100)
                                            .withMd5AuthenticationKey("fakeTokenPlaceholder"))
                                    .withConnectionIdentifier("CDD8E673-CB07-47E6-84DE-3739F778762B")))
                    .withPeerAsn(
                        new SubResource().withId("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1")))
            .withPeeringLocation("peeringLocation0")
            .create();
    }
}
