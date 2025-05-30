
import com.azure.resourcemanager.orbital.models.AutoTrackingConfiguration;
import com.azure.resourcemanager.orbital.models.ContactProfileLink;
import com.azure.resourcemanager.orbital.models.ContactProfileLinkChannel;
import com.azure.resourcemanager.orbital.models.ContactProfileThirdPartyConfiguration;
import com.azure.resourcemanager.orbital.models.ContactProfilesPropertiesNetworkConfiguration;
import com.azure.resourcemanager.orbital.models.Direction;
import com.azure.resourcemanager.orbital.models.EndPoint;
import com.azure.resourcemanager.orbital.models.Polarization;
import com.azure.resourcemanager.orbital.models.Protocol;
import java.util.Arrays;

/**
 * Samples for ContactProfiles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/ContactProfileCreate.json
     */
    /**
     * Sample code: Create a contact profile.
     * 
     * @param manager Entry point to OrbitalManager.
     */
    public static void createAContactProfile(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.contactProfiles().define("CONTOSO-CP").withRegion("eastus2").withExistingResourceGroup("contoso-Rgp")
            .withNetworkConfiguration(new ContactProfilesPropertiesNetworkConfiguration().withSubnetId(
                "/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Network/virtualNetworks/contoso-vnet/subnets/orbital-delegated-subnet"))
            .withLinks(Arrays.asList(
                new ContactProfileLink().withName("contoso-uplink").withPolarization(Polarization.LHCP)
                    .withDirection(Direction.UPLINK).withGainOverTemperature(0.0F).withEirpdBW(45.0F)
                    .withChannels(Arrays.asList(
                        new ContactProfileLinkChannel().withName("contoso-uplink-channel").withCenterFrequencyMHz(2250f)
                            .withBandwidthMHz(2f).withEndPoint(new EndPoint().withIpAddress("10.1.0.4")
                                .withEndPointName("ContosoTest_Uplink").withPort("50000").withProtocol(Protocol.TCP)))),
                new ContactProfileLink().withName("contoso-downlink").withPolarization(Polarization.RHCP)
                    .withDirection(Direction.DOWNLINK).withGainOverTemperature(25.0F).withEirpdBW(0.0F)
                    .withChannels(Arrays.asList(new ContactProfileLinkChannel().withName("contoso-downlink-channel")
                        .withCenterFrequencyMHz(8160f).withBandwidthMHz(15f)
                        .withEndPoint(new EndPoint().withIpAddress("10.1.0.5").withEndPointName("ContosoTest_Downlink")
                            .withPort("50001").withProtocol(Protocol.UDP))))))
            .withMinimumViableContactDuration("PT1M").withMinimumElevationDegrees(5.0F)
            .withAutoTrackingConfiguration(AutoTrackingConfiguration.DISABLED)
            .withEventHubUri(
                "/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.EventHub/namespaces/contosoHub/eventhubs/contosoHub")
            .withThirdPartyConfigurations(Arrays.asList(
                new ContactProfileThirdPartyConfiguration().withProviderName("KSAT")
                    .withMissionConfiguration("Ksat_MissionConfiguration"),
                new ContactProfileThirdPartyConfiguration().withProviderName("VIASAT")
                    .withMissionConfiguration("Viasat_Configuration")))
            .create();
    }
}
